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
	storeWithTx := getStoreWithTx(capacity)
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
	storeWithTx := getStoreWithTx(capacity)
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
