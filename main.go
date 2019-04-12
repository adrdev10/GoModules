package hello

import (
	"rsc.io/quote/v3"
)

// Hello returns a string literal
func Hello() string {
	return quote.HelloV3()
}

// Proverb returns a string literal
func Proverb() string {
	return quote.Concurrency()

}
