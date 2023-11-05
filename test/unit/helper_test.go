package unit

import (
	"testing"

	"challenge-test-synapsis/helper"
)

func TestNewUlid(t *testing.T) {
	id, err := helper.NewUlid("01HEDXSY934PJ0G5XH2DDH21X")
	t.Log(err)
	t.Log(id)
}
