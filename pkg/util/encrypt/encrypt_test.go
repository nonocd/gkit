package encrypt

import (
	"fmt"
	"testing"
)

func TestPasswordEncode(t *testing.T) {
	pass := "admin"
	hp := PasswordEncode(pass)
	mp := MD5Encode(hp)
	fmt.Printf("hashed: %s, len: %d\n", hp, len(hp))
	fmt.Printf("hashed: %s, len: %d\n", mp, len(mp))
}
