package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if len(r.name) == 5 {
		return r.name, nil
	}
	if r.name == "" {
		r.Reset()
		return r.name, nil
	}
	return "", errors.New("Invalid name")
}

func (r *Robot) Reset() {
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	char1 := rune(rand.Intn(26) + 65)
	char2 := rune(rand.Intn(26) + 65)
	num := int64(ran.Intn(999))
	r.name = fmt.Sprintf("%c%c%03d", char1, char2, num)

}
