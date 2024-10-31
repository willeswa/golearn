package main

import "testing"

func TestPointers(t *testing.T) {
	t.Run("deposit money", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("withdraw money", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		wallet.Withdraw(Bitcoin(5))

		got := wallet.Balance()
		want := Bitcoin(15)

		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("withdraw more than you have", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(Bitcoin(20))

		if err == nil {
			t.Errorf("expected error but got nil")
		}

	})
}
