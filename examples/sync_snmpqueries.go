package main

import (
	"fmt"

	"github.com/eferro/gosnmpquerier"
)

const (
	CONTENTION = 4
)

func main() {

	querier := gosnmpquerier.NewSyncQuerier(CONTENTION)

	for id := 0; id < 10; id++ {
		q, _ := gosnmpquerier.FromJson(`{"command":"walk", "destination":"localhost", "community":"public", "oid":"1.3.6.1"}`)
		q.Id = id

		fmt.Println("Result:", querier.ExecuteQuery(*q))
	}
}
