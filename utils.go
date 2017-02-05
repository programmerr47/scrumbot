package main

import (
	"math/rand"
	"strings"
	"fmt"
)

func randString(array []string) string {
	rndElem := rand.Intn(len(array))
	return array[rndElem]
}

func formatIfHas(string string, args ...string) string {
	if strings.Contains(string, "%s") {
		return fmt.Sprintf(string, args)
	} else {
		return string
	}
}
