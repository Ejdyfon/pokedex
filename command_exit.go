package main

import (
	"fmt"
	"os"
)

func commandExit(config *config, param string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
