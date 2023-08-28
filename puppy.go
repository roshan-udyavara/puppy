package puppy

import (
	"github.com/roshan-udyavara/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBark() string {
	return dog.WhenGrownUp(Barks())
}
