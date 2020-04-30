package main
/*
略模式
定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则。
 */

import "fmt"
type PaymentContext struct {
	Name, CardID string
	Money        int
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}
//-----------------------------------------------------------------
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}


func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

//-----------------------------
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

//-----------------------------
type Bank struct{}
func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)

}
func main() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash

	payment2 := NewPayment("Bob", "0002", 888, &Bank{})
	payment2.Pay()
}

