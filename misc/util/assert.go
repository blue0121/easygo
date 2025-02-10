package util

import (
	"log"
)

func AssertIsTrue(exp bool, msg string) {
	if !exp {
		log.Panicln(msg)
	}
}

func AssertIsFalse(exp bool, msg string) {
	if exp {
		log.Panicln(msg)
	}
}
