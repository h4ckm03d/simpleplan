package main

import (
	"io"
	"log"
)

// dclose closer with err check
func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
