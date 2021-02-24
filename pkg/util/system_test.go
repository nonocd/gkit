package util

import (
	"fmt"
	"testing"
)

func TestGetIPv4(t *testing.T) {
	if ip, ok := GetIPv4(); ok {
		fmt.Println(ip)
	}
}
