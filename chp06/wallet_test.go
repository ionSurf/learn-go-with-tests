package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10.0))
	})
	t.Run("Widthraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(10.0),
		}
		err := wallet.Widthraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(5.0))
		assertNotErr(t, err)
	})
	t.Run("Widthraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10.0)
		wallet := Wallet{
			balance: startingBalance,
		}
		err := wallet.Widthraw(Bitcoin(15.0))
		expectedErrString := ErrInsufficientFunds

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, expectedErrString)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	if wallet.Balance() != expected {
		t.Errorf("Expected %s, result %s", expected, wallet.balance)
	}
}
func assertNotErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Not execting an error")
	}
}
func assertError(t testing.TB, resultErr error, expectedErr error) {
	t.Helper()
	if resultErr == nil {
		t.Fatal("Expected an error to come through")
	}
	if resultErr != expectedErr {
		t.Errorf("Expected error %q received %q", expectedErr, resultErr)
	}
}
