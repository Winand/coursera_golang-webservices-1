package main

import "fmt"

type Payer interface {
	Pay(int) error
}

// Для реализации интерфейса структура должна иметь реализацию всех элементов интерфейса (duck typing)
// или включать встроенную (embedded) структуру, которая, в свою очередь, реализует этот интерфейс
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

// Функция Buy принимает любой объект, который реализует интерфейс Payer
func Buy(p Payer) {
	switch p.(type) { // Проверка типа объекта, реализующего интерфейс Payer
	case *Wallet:
		fmt.Println("Оплата наличными?")
	case *Card:
		// Приведение к типу *Card
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("Не удалось преобразовать к типу *Card")
		}
		fmt.Println("Вставляйте карту,", plasticCard.CardHolder)
	default:
		fmt.Println("Что-то новое!")
	}
	err := p.Pay(10)
	if err != nil {
		// panic(err)
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	// Verb %T - тип объекта
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

// ------------------------------------------------

type Card struct {
	Balance    int
	ValidUntil string
	CardHolder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Не хватает денег на карте")
	}
	c.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Не хватает денег на счёте Apple Pay")
	}
	a.Money -= amount
	return nil
}

// ------------------------------------------------

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	// В переменную типа Payer можно поместить любой объект,
	// который реализует этот интерфейс
	var myMoney Payer
	myMoney = &Card{Balance: 100, CardHolder: "rvasily"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}
