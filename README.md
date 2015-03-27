# go-connpass
CONNPASS の API (http://connpass.com/about/api/) を Go でラップしたものです。

## 使い方

	query := connpass.Query{Start: 1, Order: connpass.CREATE}
	query.KeywordOr = []string{"go", "golang"}
	query.Time = []connpass.Time{connpass.Time{Year: 2015, Month: 3}, connpass.Time{Year: 2015, Month: 4}}
	result, e := query.Search()
	
	if e != nil {
	   log.Fatalf("Failed to fetch the result: %v\n", e)
	} else {
	   log.Printf("Available: %d, Returned: %d, Offset: %d\n", result.Available, result.Returned, result.Start)
	}
