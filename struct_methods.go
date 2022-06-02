package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

//-----------------------------------

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}
func main() {
	pers := Person{1, "Vasily"}
	// Если pers является указателем:
	// pers := &Person{1, "Vasily"} или pers := new(Person)
	// вызов pers.SetName производится по указателю
	pers.SetName("Vasily Romanov")
	// Даже если pers не является указателем,
	// компилятор автоматически делает вызов по указателю
	// (&pers).SetName("Vasily Romanov")
	// т.к. в объявлении функции указано *Person
	fmt.Printf("updated person: %#v\n", pers)

	acc := Account{
		Id:   1,
		Name: "rvasily",
		Person: Person{
			Id:   2,
			Name: "Vasily Romanov",
		},
	}
	// Т.к. Account имеет метод SetName, будет вызван он
	// В противном случае был бы вызван Person.SetName
	acc.SetName("romanov.vasily")
	fmt.Printf("%#v\n", acc)
	// Вызов метода встроенной структуры можно произвести
	// через указание её имени
	acc.Person.SetName("V.R.")
	fmt.Printf("%#v\n", acc)

	//----------------------------------

	sl := MySlice([]int{1, 2})
	sl.Add(5)
	fmt.Println(sl.Count(), sl)
}
