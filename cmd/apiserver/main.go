package main

import (
	"github.com/bumblebizoom/delivery_by_bumblebizoom/internal/app/myapiserver"
	"log"
)

func main() {
	s := myapiserver.APIserver{}
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
