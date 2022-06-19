package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	// Переменную нужно передавать в горутину как аргумент, чтобы закрепить её значение
	// https://www.coursera.org/learn/golang-webservices-1/discussions/forums/4-A-GM1uEee4iBJpbmnIvg/threads/0i1G0HswEemBSQpvxxG8fA/replies/m_pdt1kPQqS6XbdZD6Kkiw
	// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
	// https://go.dev/doc/faq#closures_and_goroutines
	var nextIn chan interface{}
	wg := &sync.WaitGroup{}

	for _, j := range jobs {
		j := j // https://go.dev/doc/effective_go#channels
		out := make(chan interface{})
		wg.Add(1)
		go func(in chan interface{}) {
			defer wg.Done()
			defer close(out)
			j(in, out)
		}(nextIn)
		nextIn = out
	}
	wg.Wait()
}

type Pair struct {
	Idx int
	Val string
}

func SingleHash(in chan interface{}, out chan interface{}) {
	var md5s []string
	ch1, ch2 := make(chan Pair), make(chan Pair)
	// hash (parallel)
	cursor := 0
	for v := range in {
		v := strconv.Itoa(v.(int))
		md5s = append(md5s, DataSignerMd5(v))
		go func(curs int) {
			ch1 <- Pair{curs, DataSignerCrc32(v)}
		}(cursor)
		go func(curs int) {
			ch2 <- Pair{curs, DataSignerCrc32(md5s[curs])}
		}(cursor)
		cursor++
	}
	// collect
	crcs := make(map[int]string)
	for i := 0; i < len(md5s); i++ {
		p := <-ch1
		crcs[100+p.Idx] = p.Val
		p = <-ch2
		crcs[200+p.Idx] = p.Val
	}
	// send
	for i := 0; i < len(md5s); i++ {
		out <- crcs[100+i] + "~" + crcs[200+i]
	}
}

func MultiHash(in chan interface{}, out chan interface{}) {
	chParts := make(chan Pair)
	// hash (parallel)
	idx := 0
	for v := range in {
		// fmt.Println("in", v)
		for th := 0; th < 6; th++ {
			th := th // Иначе в горутинах будет последнее значение
			v := v   //
			go func(idxIn int) {
				chParts <- Pair{idxIn*10 + th, DataSignerCrc32(fmt.Sprint(th, v))}
			}(idx)
		}
		idx++
	}
	// collect (1)
	crcs := make(map[int]map[int]string)
	count := 0
	for part := range chParts {
		// fmt.Println(part, part.Idx/10, part.Idx%10)
		if crcs[part.Idx/10] == nil {
			crcs[part.Idx/10] = make(map[int]string)
		}
		crcs[part.Idx/10][part.Idx%10] = part.Val
		if count++; count == idx*6 {
			break
		}
	}
	// sort (2)
	for i := 0; i < len(crcs); i++ {
		result := ""
		for el := 0; el < 6; el++ {
			result += crcs[i][el]
		}
		// fmt.Println("out", i, result)
		out <- result
	}
}

func CombineResults(in chan interface{}, out chan interface{}) {
	var results []string
	var result string
	for v := range in {
		results = append(results, v.(string))
	}
	sort.Strings(results)
	for i, s := range results {
		result += s
		if i < len(results)-1 {
			result += "_"
		}
	}
	out <- result
}
