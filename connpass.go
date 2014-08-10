package connpass

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "http://connpass.com/api/v1/event/"

type ResultSet struct {
	Returned  int     `json:"results_returned"`
	Available int     `json:"results_available"`
	Start     int     `json:"results_start"`
	Events    []Event `json:"events"`
}

type Event struct {
	Id            int     `json:"event_id"`
	Title         string  `json:"title"`
	Catch         string  `json:"catch"`
	Description   string  `json:"description"`
	Url           string  `json:"event_url"`
	Tag           string  `json:"hash_tag"`
	Start         string  `json:"started_at"`
	End           string  `json:"ended_at"`
	Limit         int     `json:"limit"`
	Etype         string  `json:"event_type"`
	Address       string  `json:"address"`
	Place         string  `json:"place"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	OwnerID       int     `json:"owner_id"`
	OwnerNickname string  `json:"owner_nickname"`
	OwnerName     string  `json:"owner_display_name"`
	Accepted      int     `json:"accepted"`
	Waiting       int     `json:"waiting"`
	Updated       string  `json:"updated_at"`
}

type Order int

const (
	_      Order = iota
	UPDATE       // 1: descending in updated time
	START        // 2: descending in event start time
	CREATE       // 3: descending in created time
)

type Query struct {
	Start int
	Order Order
	Count int
}

func NewQuery() Query {
	q := Query{0, CREATE, 100}
	return q
}

func (q Query) buildURL() string {
	return fmt.Sprint(baseURL, "?start=", q.Start, "&order=", q.Order, "&count=", q.Count)
}

func parse(jsonBlob []byte) (*ResultSet, error) {
	res := new(ResultSet)
	err := json.Unmarshal(jsonBlob, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (q Query) Search() (*ResultSet, error) {
	res, err := http.Get(q.buildURL())
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return parse(body)
}
