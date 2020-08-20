package dice

import (
	"fmt"
	"testing"
)

func TestRoll(t *testing.T) {
	dice := New(6)
	for r := 1; r <= 100; r++ {
		roll, err := Roll(dice)
		if err != nil {
			t.Fail()
		}
		fmt.Println(roll)
	}
}

func TestRollMany(t *testing.T) {
	dice := New(45)
	rolls, err := RollMany(dice, 5)
	if err != nil {
		t.Fail()
	}
	fmt.Println(rolls)
}
