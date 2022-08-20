package db

import (
	"context"
	"testing"

	util "github.com/OktarianTB/stock-trading-simulator-golang/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUsername(),
		HashedPassword: util.RandomString(6),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.NotZero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
}

func TestAddUserBalance(t *testing.T) {
	user1 := createRandomUser(t)
	amount := util.RandomFloat(1, 100)
	user2, err := testQueries.AddUserBalance(context.Background(), AddUserBalanceParams{
		Username: user1.Username,
		Amount: amount,
	})

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Balance + amount, user2.Balance)
}

func TestRemoveUserBalance(t *testing.T) {
	user1 := createRandomUser(t)

	amount1 := util.RandomFloat(50, 100)
	user2, err := testQueries.AddUserBalance(context.Background(), AddUserBalanceParams{
		Username: user1.Username,
		Amount: amount1,
	})
	require.NoError(t, err)

	amount2 := util.RandomFloat(1, 50)
	user3, err := testQueries.RemoveUserBalance(context.Background(), RemoveUserBalanceParams{
		Username: user1.Username,
		Amount: amount2,
	})

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.Balance - amount2, user3.Balance)
}