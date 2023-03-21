package main

import (
	"log"
	"os"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func main() {
	_, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}

	err = New("emit macho dwarf: elf header corrupted")

	if err != nil {
		log.Fatal(err)
	}

}
