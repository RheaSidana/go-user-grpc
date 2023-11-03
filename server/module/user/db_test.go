package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	InitDB()
	assert.Equal(t, "Steve", DB.Users[1].Fname)
	assert.Equal(t, "LA", DB.Users[1].City)
	assert.Equal(t, int64(1234567890), DB.Users[1].Phone)
	assert.Equal(t, 5.8, DB.Users[1].Height)
	assert.Equal(t, true, DB.Users[1].Married)

	assert.Equal(t, "Alice", DB.Users[2].Fname)
	assert.Equal(t, "New York", DB.Users[2].City)
	assert.Equal(t, int64(1234567890), DB.Users[2].Phone)
	assert.Equal(t, 5.6, DB.Users[2].Height)
	assert.Equal(t, true, DB.Users[2].Married)

	assert.Equal(t, "Bob", DB.Users[3].Fname)
	assert.Equal(t, "San Francisco", DB.Users[3].City)
	assert.Equal(t, int64(9876543210), DB.Users[3].Phone)
	assert.Equal(t, 6.0, DB.Users[3].Height)
	assert.Equal(t, false, DB.Users[3].Married)

	assert.Equal(t, "Charlie", DB.Users[4].Fname)
	assert.Equal(t, "Los Angeles", DB.Users[4].City)
	assert.Equal(t, int64(5555555555), DB.Users[4].Phone)
	assert.Equal(t, 5.10, DB.Users[4].Height)
	assert.Equal(t, true, DB.Users[4].Married)

	assert.Equal(t, "David", DB.Users[5].Fname)
	assert.Equal(t, "Chicago", DB.Users[5].City)
	assert.Equal(t, int64(1112223333), DB.Users[5].Phone)
	assert.Equal(t, 5.8, DB.Users[5].Height)
	assert.Equal(t, false, DB.Users[5].Married)
}
