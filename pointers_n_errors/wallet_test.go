package pointersnerrors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Money(10))
		assertBalance(t, wallet, Money(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Money(20)}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet, Money(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Money(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Money(100))

		assertError(t, err, ErrInsuficientFounds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Money) {
	t.Helper()
	result := wallet.Balance()

	if result != expected {
		t.Errorf("got: %s but expected: %s", result, expected)
	}
}

func assertError(t testing.TB, err, expected error) {
	t.Helper()
	if err == nil {
		t.Error("expected an error but didn't got one")
	}
	if err != expected {
		t.Errorf("got: %q, but expected: %q", err, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("expected an error but didn't got one")
	}

}
