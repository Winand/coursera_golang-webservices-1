/* Calculate total download size for NYC taxi data for 2021

For each month, we have two files: green and yellow. For example:

    https://d37ci6vzurychx.cloudfront.net/trip-data/green_tripdata_2021-03.parquet
	https://d37ci6vzurychx.cloudfront.net/trip-data/yellow_tripdata_2021-03.parquet

See also TLC Trip Record Data
https://www1.nyc.gov/site/tlc/about/tlc-trip-record-data.page
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	urlTemplate = "https://d37ci6vzurychx.cloudfront.net/trip-data/%s_tripdata_2021-%02d.parquet"
	colors      = []string{"green", "yellow"}
)

type Size struct {
	Size int
	Err  error
}

func downloadSize(url string, ch chan Size) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		ch <- Size{0, err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- Size{0, err}
		return
	}

	if resp.StatusCode != http.StatusOK {
		ch <- Size{0, fmt.Errorf(resp.Status)}
		return
	}

	v, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	fmt.Printf("Got %d length of %s at %s\n", v, url, time.Now().Format("15:04:05.999"))
	ch <- Size{v, err}
}

func main() {
	// os.Setenv("https_proxy", "http://host.docker.internal:8888/")
	start := time.Now()
	size := 0
	ch := make(chan Size, 1)
	for month := 1; month <= 12; month++ {
		for _, color := range colors {
			url := fmt.Sprintf(urlTemplate, color, month)
			go downloadSize(url, ch)
		}
	}
	for i := 0; i < 12*len(colors); i++ {
		sz := <-ch
		n, err := sz.Size, sz.Err
		if err != nil {
			log.Fatal(err)
		}
		size += n
	}

	duration := time.Since(start)
	fmt.Println(size, duration)
}
