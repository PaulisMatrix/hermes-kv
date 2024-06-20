package hermeskv

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetKV(t *testing.T) {
	store := GetNewKV(1)
	defer store.Close()

	key := "hello"
	value := "world"

	store.Set(key, value)

	val, err := store.Get(key)
	require.Nil(t, err)
	assert.EqualValues(t, "world", val)

}

func TestGetKV(t *testing.T) {
	store := GetNewKV(1)
	defer store.Close()

	key := "hello"
	value := "world"

	store.Set(key, value)
	invalidKey := "invalid"

	_, err := store.Get(invalidKey)
	require.NotNil(t, err)
	expectedError := ErrNoKey
	assert.EqualError(t, err, expectedError.Error())

}

func TestDeleteKV(t *testing.T) {
	store := GetNewKV(1)
	defer store.Close()

	key := "hello"
	value := "world"

	store.Set(key, value)

	err := store.Delete(key)
	require.Nil(t, err)

	_, err = store.Get(key)
	require.NotNil(t, err)
	expectedError := ErrNoKey
	assert.EqualError(t, err, expectedError.Error())

}

func TestKVCapBreach(t *testing.T) {
	capacity := 4
	store := GetNewKV(capacity)
	defer store.Close()

	for i := 0; i < capacity; i++ {
		key := fmt.Sprintf("key:%d", i)
		val := fmt.Sprintf("value:%d", i)

		store.Set(key, val)
	}

	// capacity is breached for this kv pair
	store.Set("key:4", "value4")

	// if we check for key:0 then it should throw key not found error
	val, err := store.Get("key:0")
	expectedError := ErrNoKey
	require.Nil(t, val)
	assert.EqualError(t, err, expectedError.Error())
}

func TestZeroCapKV(t *testing.T) {
	capacity := 0

	var f assert.PanicTestFunc

	f = func() {
		store := GetNewKV(capacity)
		defer store.Close()

		store.Set("hello", "world")
	}

	require.Panics(t, f)
}

func TestKVSetRacer(t *testing.T) {
	capacity := 5
	var wg sync.WaitGroup

	store := GetNewKV(capacity)
	defer store.Close()

	// set values
	for i := 0; i < capacity; i++ {
		wg.Add(1)
		go func(store *Store, id int) {
			defer wg.Done()

			key := fmt.Sprintf("key:%d", id)
			value := fmt.Sprintf("value:%d", id)
			err := store.Set(key, value)
			require.Nil(t, err)

		}(store, i)
	}

	// wait for keys to store
	time.Sleep(2 * time.Second)

	// get values
	for i := 0; i < capacity; i++ {
		wg.Add(1)
		go func(store *Store, id int) {
			defer wg.Done()

			key := fmt.Sprintf("key:%d", id)
			_, err := store.Get(key)
			require.Nil(t, err)

		}(store, i)
	}

	wg.Wait()
}

func TestKVDeleteRacer(t *testing.T) {
	capacity := 5
	var wg sync.WaitGroup

	store := GetNewKV(capacity)
	defer store.Close()

	// set values
	for i := 0; i < capacity; i++ {
		wg.Add(1)
		go func(store *Store, id int) {
			defer wg.Done()

			key := fmt.Sprintf("key:%d", id)
			value := fmt.Sprintf("value:%d", id)
			err := store.Set(key, value)
			require.Nil(t, err)

		}(store, i)
	}

	// wait for keys to store
	time.Sleep(2 * time.Second)

	// get values
	for i := 0; i < capacity; i++ {
		wg.Add(1)
		go func(store *Store, id int) {
			defer wg.Done()

			key := fmt.Sprintf("key:%d", id)
			err := store.Delete(key)
			require.Nil(t, err)

		}(store, i)
	}

	wg.Wait()
}
