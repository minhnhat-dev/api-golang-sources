package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t) 

	amount := int64(10)
	require.NotEmpty(t, amount)
	require.NotEmpty(t, account1)
	require.NotEmpty(t, account2)

	result, err := store.TransferTx(context.Background(), TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	})
	
	require.NoError(t, err)
	require.NotEmpty(t, result)
}