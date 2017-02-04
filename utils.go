package main

import "math/rand"

func randString(array []string) string {
	rndElem := rand.Intn(len(array))
	return array[rndElem]
}
