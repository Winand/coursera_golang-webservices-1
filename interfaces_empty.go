// Пустой интерфейс

package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

// Реализация интерфейса Stringer (внутренний тип в Go)
// Вызывается, например, fmt.Printf("%s", obj)
func (w *Wallet) String() string {
	return "Кошелёк, в котором " + strconv.Itoa(w.Cash) + " денег"
}

type Payer interface {
	Pay(int) error
}

// Принимает пустой интерфейс, т. е. любой объект
func Buy(in interface{}) {
	var p Payer
	var ok bool
	// Приведение типа для проверки, реализует ли объект интерфейс Payer
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не является платёжным средством\n\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	// fmt.Printf принимает пустой интерфейс,
	// т. е. любой объект
	fmt.Printf("Raw payment: %#v\n", myWallet)
	fmt.Printf("Способ оплаты: %s\n", myWallet)

	// ----------------------------------------

	Buy(myWallet)
	Buy([]int{1, 2, 3})
	Buy(3.14)
}
