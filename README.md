# go-connpass
CONNPASS の API (http://connpass.com/about/api/) を Go でラップしたものです。

## 使い方

      query := connpass.NewQuery()
	  query.Order = connpass.CREATE
	  query.Start = 0
	  result, e := query.Search()
	  if e != nil {
	    log.Fatalf("Failed to fetch the result: %v\n", e)
	  } else {
	    log.Printf("Available: %d, Returned: %d, Offset: %d\n", result.Available, result.Returned, result.Start)
	  }
