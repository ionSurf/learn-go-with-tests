package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		result := wallet.Balance()
		expected := Bitcoin(10.0)

		if result != expected {
			t.Errorf("%v expected %s result %s", wallet, expected, result)
		}
	})
	t.Run("Widthraw", func(t *testing.T) {
		wallet := Wallet{
			balance: 10.0,
		}

		wallet.Widthraw(Bitcoin(5))
		result := wallet.Balance()
		expected := Bitcoin(5.0)

		if result != expected {
			t.Errorf("%v expected %s result %s", wallet, expected, result)
		}
	})
}
