package common

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomHexString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
