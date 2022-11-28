package hello

import (
    "example.com/greeting"
    "example.com/hello/quotes"
)

func Hello() string {
	return quotes.Hello()
}

func SayHello(lang string) string {
    return greeting.Hello(lang)
}
