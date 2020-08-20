package dice

import (
	"crypto/rand"
	"math/big"
)

type Roller interface {
	Dots() int64
}

type defaultDice struct {
	dots int64
}

func (dd defaultDice) Dots() int64 {
	return dd.dots
}

func New(dots int64) (dice Roller) {
	if dots <= 0 {
		dots = 6
	}

	return defaultDice{
		dots: dots,
	}
}

func Roll(dice Roller) (roll uint64, err error) {
	max := big.NewInt(dice.Dots())
	random, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return random.Uint64() + 1, nil
}

func RollMany(dice Roller, count int) (rolls []uint64, err error) {
	rolls = make([]uint64, count)
	for r := 1; r <= count; r++ {
		roll, err := Roll(dice)
		if err != nil {
			return []uint64{}, err
		}
		rolls[r-1] = roll
	}
	return rolls, nil
}
