package libs

import "fmt"

func Func1(str string) string{
	return "hello, " + str;
}

func Add(i1 int, i2 int) int{
	return i1 + i2;
}

type Bitcon int
type Wallet struct{
	balance Bitcon
}


func (w *Wallet) Deposit(input Bitcon) {
	w.balance += input
}

func (w *Wallet) Balance() Bitcon {
	return w.balance
}

func (b Bitcon) String() string{
	return fmt.Sprintf("BC %d", b)
}
