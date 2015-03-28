package main

import (
	"fmt"
	connpass "github.com/hkurokawa/go-connpass"
)

func main() {
	query := connpass.Query{Start: 1, Order: connpass.CREATE}
	query.KeywordOr = []string{"go", "golang"}
	query.Time = []connpass.Time{connpass.Time{Year: 2015, Month: 4}}
	var allEvents []connpass.Event
	for {
		res, err := query.Search()
		if err != nil {
			fmt.Errorf("Failed to execute search: %v.", err)
			return
		}
		allEvents = append(allEvents, res.Events...)
		offset := res.Start + res.Returned
		if offset > res.Available {
			break
		}
		query.Start = offset
	}
	for _, e := range allEvents {
		fmt.Printf("\t%s\t%d\t%s\n", e.Start, e.Id, e.Title)
	}
}
