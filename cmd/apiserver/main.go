package main

import (
	"log"

	"github.com/dreddsa5dies/httprestapi/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
