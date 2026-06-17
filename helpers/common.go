package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateReference() string {
	now := time.Now()
	timeFormat := now.Format("20060102150405")
	randomNumber := rand.Intn(100)
	ref := fmt.Sprintf("%s%d", timeFormat, randomNumber)
	return ref
}
