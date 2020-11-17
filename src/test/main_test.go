package test

import (
	"fmt"
	"libs"
	"testing"
)

func TestAdd(t *testing.T){
	got := libs.Add(1,2)
	want := 3

	if got != want{
		t.Error("got", got, "want", want)
	}
}

func ExampleAdd(){
	sum := libs.Add(2,3)
	fmt.Println(sum)
	// Output: 5
}

func TestWallet(t *testing.T){
	wallet := libs.Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := libs.Bitcon(15)

	if got != want{
		t.Errorf("got is %s, want is %s", got, want)
	}
}
