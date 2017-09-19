package main

import (
	"github.com/slarti5191/splendid"
)

// main initializes the configuration
// and kicks off the collectors
func main() {
	s := new(splendid.Splendid)
	s.Run()
}
