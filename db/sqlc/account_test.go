package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/mokshchadha/go_banks/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	return account
}

func TestCreatAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	fetched, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetched)

	require.Equal(t, fetched.Owner, account.Owner)
	require.Equal(t, fetched.Balance, account.Balance)
	require.Equal(t, fetched.Currency, account.Currency)
	require.Equal(t, fetched.ID, account.ID)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	updated, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updated)

	require.Equal(t, updated.Owner, account1.Owner)
	require.Equal(t, updated.Balance, arg.Balance)
	require.Equal(t, updated.ID, arg.ID)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	testQueries.DeleteAccount(context.Background(), account.ID)

	fetched, err := testQueries.GetAccount(context.Background(), account.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, fetched)

}
