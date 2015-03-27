package main

import (
	"fmt"
	connpass "github.com/hkurokawa/go-connpass"
)

func main() {
	query := connpass.Query{Start: 1, Order: connpass.CREATE}
	query.KeywordOr = []string{"go", "golang"}
	query.Time = []connpass.Time{connpass.Time{Year: 2015, Month: 3}, connpass.Time{Year: 2015, Month: 4}}
	res, err := query.Search()
	if err != nil {
		fmt.Errorf("Failed to execute search: %v.", err)
		return
	}
	fmt.Printf("Num returned: %d\n", res.Returned)
	fmt.Printf("Num available: %d\n", res.Available)
	fmt.Printf("Start position: %d\n", res.Start)
	for _, e := range res.Events {
		fmt.Printf("\t%s\t%d\t%s\n", e.Start, e.Id, e.Title)
	}
}
