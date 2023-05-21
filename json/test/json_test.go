package test

import (
	"testing"

	"github.com/m4gshm/expressions/json"
	"github.com/stretchr/testify/assert"
)

func Test_json(t *testing.T) {

	type User struct {
		Name, Lastname string
	}

	user := User{"Bob", "Ryan"}
	bytes, err := json.Marshal(user)
	assert.NoError(t, err)

	user2, err2 := json.Unmarshal[User](bytes)
	assert.NoError(t, err2)

	assert.Equal(t, user, user2)

}
