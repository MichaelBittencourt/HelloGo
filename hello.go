package main

import (
	"fmt"
	"github.com/MichaelBittencourt/greatings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	message := greatings.Hello("Michael")
	fmt.Println(message)
}
