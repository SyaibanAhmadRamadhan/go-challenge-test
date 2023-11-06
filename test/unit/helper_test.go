package unit

import (
	"testing"

	"challenge-test-synapsis/helper"
	"challenge-test-synapsis/usecase"
)

func TestNewUlid(t *testing.T) {
	id, err := helper.NewUlid("01HEDXSY934PJ0G5XH2DDH21X")
	t.Log(err)
	t.Log(id)
}

func TestBcrypt(t *testing.T) {
	bcrypt, err := usecase.HashBcrypt("test123")
	if err != nil {
		return
	}
	t.Log(bcrypt)
}
