package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	client, err := selectClient(db, clientID)
	require.NoError(t, err)

	assert.Equal(t, clientID, client.ID)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1

	// напиши тест здесь
	client, err := selectClient(db, clientID)

	require.ErrorIs(t, err, sql.ErrNoRows)
	assert.Empty(t, client.ID)
	assert.Empty(t, client.Birthday)
	assert.Empty(t, client.Email)
	assert.Empty(t, client.Login)
	assert.Empty(t, client.FIO)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	cl.ID = id
	require.NoError(t, err)
	require.NotEmpty(t, id)

	client, err := selectClient(db, id)
	require.NoError(t, err)
	assert.Equal(t, cl.ID, client.ID)
	assert.Equal(t, cl.Birthday, client.Birthday)
	assert.Equal(t, cl.Email, client.Email)
	assert.Equal(t, cl.FIO, client.FIO)
	assert.Equal(t, cl.Login, client.Login)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	require.NotEmpty(t, id)
	require.NoError(t, err)

	client, err := selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, client.ID)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.ErrorIs(t, err, sql.ErrNoRows)

}
