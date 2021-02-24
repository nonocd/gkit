package identity

import (
	"fmt"
	"testing"
)

func TestID(t *testing.T) {
	id := ID()
	fmt.Println(id)
}
