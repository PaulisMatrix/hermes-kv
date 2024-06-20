package hermeskv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSet1KVWithTx(t *testing.T) {
	// * Begin transaction
	// * Set key2 to val2
	// * Get key2
	//   * Expect val2
	// * Rollback
	// * Get key2
	//   * Expect an error case as key2 is not set
	capacity := 10
	storeWithTx := GetStoreWithTx(capacity)
	defer storeWithTx.Close()

	storeWithTx.Begin()

	storeWithTx.Set("key2", "val2")
	val, err := storeWithTx.Get("key2")
	require.Nil(t, err)
	assert.EqualValues(t, "val2", val)

	storeWithTx.Rollback()
	storeWithTx.End()

	_, err = storeWithTx.Get("key2")
	require.NotNil(t, err)
	expectedError := ErrNoKey
	assert.EqualError(t, err, expectedError.Error())

}

func TestSet2KVWithTx(t *testing.T) {
	// * Set key0 to val0
	// * Get key0
	//   * Expect val0
	// * Begin transaction
	// * Within transaction: Get key0
	//   * Expect val0
	// * Within transaction: Set key1 to val1
	// * Within transaction: Get key1
	//   * Expect val1
	// * Commit transaction
	// * From the global state: Get key1
	//   * Expect val1

	capacity := 10
	storeWithTx := GetStoreWithTx(capacity)
	defer storeWithTx.Close()

	storeWithTx.Set("key0", "val0")
	val, err := storeWithTx.Get("key0")
	require.Nil(t, err)
	assert.EqualValues(t, "val0", val)

	storeWithTx.Begin()

	val, err = storeWithTx.Get("key0")
	require.Nil(t, err)
	assert.EqualValues(t, "val0", val)

	storeWithTx.Set("key1", "val1")
	val, err = storeWithTx.Get("key1")
	require.Nil(t, err)
	assert.EqualValues(t, "val1", val)

	storeWithTx.Commit()
	storeWithTx.End()

	val, err = storeWithTx.Get("key1")
	require.Nil(t, err)
	assert.EqualValues(t, "val1", val)

}

func TestDelete1KVWithTx(t *testing.T) {
	// * Set key0 to val0
	// * Get key0
	//   * Expect val0
	// * Begin transaction
	// * Within transaction: Get key0
	//   * Expect val0
	// * Within transaction: Delete key0
	// * Within transaction: Get key0
	//   * Expect an error case as key0 is not set
	// * Commit transaction
	// * From the global state: Get key0
	//   * Expect an error case as key0 is not set

	capacity := 10
	expectedError := ErrNoKey
	storeWithTx := GetStoreWithTx(capacity)
	defer storeWithTx.Close()

	storeWithTx.Set("key0", "val0")
	val, err := storeWithTx.Get("key0")
	require.Nil(t, err)
	assert.EqualValues(t, "val0", val)

	storeWithTx.Begin()

	val, err = storeWithTx.Get("key0")
	require.Nil(t, err)
	assert.EqualValues(t, "val0", val)

	storeWithTx.Delete("key0")
	_, err = storeWithTx.Get("key0")
	require.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	storeWithTx.Commit()
	storeWithTx.End()

	_, err = storeWithTx.Get("key0")
	require.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

	storeWithTx.Set("key1", "val1")
	val, err = storeWithTx.Get("key1")
	require.Nil(t, err)
	assert.EqualValues(t, "val1", val)

	storeWithTx.Delete("key1")
	_, err = storeWithTx.Get("key1")
	require.NotNil(t, err)
	assert.EqualError(t, err, expectedError.Error())

}
