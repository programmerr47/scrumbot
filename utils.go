package main

import (
	"math/rand"
	"strings"
	"fmt"
	"time"
)

func randString(array []string) string {
	rand.Seed(time.Now().UnixNano())
	rndElem := rand.Intn(len(array))
	return array[rndElem]
}

func formatIfHas(string string, args ...interface{}) string {
	if strings.Contains(string, "%s") {
		return fmt.Sprintf(string, args...)
	} else {
		return string
	}
}
