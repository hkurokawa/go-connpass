# go-connpass
[![GoDoc](https://godoc.org/github.com/hkurokawa/go-connpass?status.svg)](https://godoc.org/github.com/hkurokawa/go-connpass)

[![Circle CI](https://circleci.com/gh/hkurokawa/go-connpass.svg?style=shield)](https://circleci.com/gh/hkurokawa/go-connpass)

connpass のサーチ API (http://connpass.com/about/api/) を Go で実装したものです。

## インストール
    $ go get github.com/hkurokawa/go-connpass

## 使い方

	query := connpass.Query{Start: 1, Order: connpass.CREATE}
	query.KeywordOr = []string{"go", "golang"}
	query.Time = []connpass.Time{connpass.Time{Year: 2015, Month: 3}, connpass.Time{Year: 2015, Month: 4}}
	result, e := query.Search()
	
	if e != nil {
	   log.Errorf("Failed to fetch the result: %v\n", e)
	} else {
	   fmt.Printf("Num returned: %d\n", res.Returned)
	   fmt.Printf("Num available: %d\n", res.Available)
	   fmt.Printf("Start position: %d\n", res.Start)
	   for _, e := range res.Events {
	      fmt.Printf("\t%s\t%d\t%s\n", e.Start, e.Id, e.Title)
	   }
	}

## ライセンス
MIT

## 作者
Hiroshi Kurokawa
