package puppy

import (
	"github.com/andyrestart9/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BidBark() string {
	return dog.WhenGrownUp(Bark())
}

func BidBarks() string {
	return dog.WhenGrownUp(Barks())
}
