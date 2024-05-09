package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetKV(t *testing.T) {
	store := GetNewKV(1)

	key := "hello"
	value := "world"

	store.Set(key, value)

	val, err := store.Get(key)
	require.Nil(t, err)
	assert.EqualValues(t, "world", val)

}

func TestGetKV(t *testing.T) {
	store := GetNewKV(1)

	key := "hello"
	value := "world"

	store.Set(key, value)
	invalidKey := "invalid"

	_, err := store.Get(invalidKey)
	require.NotNil(t, err)
	expectedError := errors.New("Key doesn't exist")
	assert.EqualError(t, err, expectedError.Error())

}

func TestKVCapBreach(t *testing.T) {
	capacity := 4
	store := GetNewKV(capacity)

	for i := 0; i < capacity; i++ {
		key := fmt.Sprintf("key:%d", i)
		val := fmt.Sprintf("value:%d", i)

		store.Set(key, val)
	}

	// capacity is breached for this kv pair
	store.Set("key:4", "value4")

	// if we check for key:0 then it should throw key not found error
	val, err := store.Get("key:0")
	expectedError := errors.New("Key doesn't exist")
	require.Nil(t, val)
	assert.EqualError(t, err, expectedError.Error())
}

func TestZeroCapKV(t *testing.T) {
	capacity := 0

	var f assert.PanicTestFunc

	f = func() {
		s := GetNewKV(capacity)
		s.Set("hello", "world")
	}

	require.Panics(t, f)
}
