package misc

import (
	"fmt"
	"log"
)

func ErrorUsingPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func ErrorUsingFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func PrintError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
