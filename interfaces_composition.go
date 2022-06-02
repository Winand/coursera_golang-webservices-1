// Композиция интерфейсов

package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Ringer interface {
	Ring(string) error
}

type NFCPhone interface {
	// Встроенные интерфейсы
	Payer
	Ringer
}

// Принимает NFCPhone
func PayForMetroWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Ошибка при оплате %v\n\n", err)
		return
	}
	fmt.Printf("Тукникет открыт через %T\n\n", phone)
}

// Реализует NFCPhone
type Phone struct {
	Money int
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("Not enough money on account")
	}
	p.Money -= amount
	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("Please, enter phone number")
	}
	return nil
}

func main() {
	myPhone := &Phone{Money: 9}
	PayForMetroWithPhone(myPhone)
}
