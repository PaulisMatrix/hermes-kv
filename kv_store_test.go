package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetKV(t *testing.T) {
	store := GetNewKV(0)

	key := "hello"
	value := "world"

	store.Set(key, value)

	val, err := store.Get(key)
	require.Nil(t, err)
	assert.EqualValues(t, "world", val)

}

func TestGetKV(t *testing.T) {
	store := GetNewKV(0)

	key := "hello"
	value := "world"

	store.Set(key, value)
	invalidKey := "invalid"

	_, err := store.Get(invalidKey)
	require.NotNil(t, err)
	expectedError := errors.New("Key doesn't exist")
	assert.EqualError(t, err, expectedError.Error())

}
