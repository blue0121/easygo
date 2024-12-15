package convert

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestToJsonBytes(t *testing.T) {
	u := user{
		Name: "blue",
		Age:  18,
	}
	var data = ToJsonBytes(u)
	var v user
	var err = json.Unmarshal(data, &v)
	assert.NoError(t, err)
	assert.Equal(t, u.Name, v.Name)
	assert.Equal(t, u.Age, v.Age)
}

func TestToJsonString(t *testing.T) {
	u := user{
		Name: "blue",
		Age:  18,
	}
	var str = ToJsonString(u)
	var v user
	var err = json.Unmarshal([]byte(str), &v)
	assert.NoError(t, err)
	assert.Equal(t, u.Name, v.Name)
	assert.Equal(t, u.Age, v.Age)
}
