package hello

import (
	"rsc.io/quote"
	v3 "rsc.io/quote/v3"
)

// Hello returns a string literal
func Hello() string {
	return quote.Hello()
}

// Proverb returns a string literal
func Proverb() string {
	return v3.Concurrency()

}
