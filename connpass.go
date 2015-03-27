package connpass

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"strconv"
)

const BASE_URL = "http://connpass.com/api/v1/event/"

type ResultSet struct {
	Returned  int     `json:"results_returned"`
	Available int     `json:"results_available"`
	Start     int     `json:"results_start"`
	Events    []Event `json:"events"`
}

type Event struct {
	Id            int    `json:"event_id"`
	Title         string `json:"title"`
	Catch         string `json:"catch"`
	Description   string `json:"description"`
	Url           string `json:"event_url"`
	Tag           string `json:"hash_tag"`
	Start         string `json:"started_at"`
	End           string `json:"ended_at"`
	Limit         int    `json:"limit"`
	Etype         string `json:"event_type"`
	Address       string `json:"address"`
	Place         string `json:"place"`
	Lat           string `json:"lat"`
	Lon           string `json:"lon"`
	OwnerID       int    `json:"owner_id"`
	OwnerNickname string `json:"owner_nickname"`
	OwnerName     string `json:"owner_display_name"`
	Accepted      int    `json:"accepted"`
	Waiting       int    `json:"waiting"`
	Updated       string `json:"updated_at"`
}

// Order of the returned events.
type Order int

const (
	UPDATE Order = 1 + iota // 1: descending in updated time
	START                   // 2: descending in event start time
	CREATE                  // 3: descending in created time
)

// Holding date of the event. If 0 is specified as Date, it is specified as the value of "ym", "ymd" otherwise.
type Time struct {
	Year  int
	Month int
	Date  int
}

// Format of the response.
type Format string

const (
	JSON Format = "json"
)

type Query struct {
	EventId     []int    // Event ID
	KeywordAnd  []string // Keywords combined with AND operator
	KeywordOr   []string // Keywords combined with OR operator
	Time        []Time   // Holding date of the event
	Participant []string //  Nickname of participants
	Owner       []string // Nickname of the owner
	SeriesId    []int    // Series ID
	Start       int      // Offset
	Order       Order    // Order of the result
	Count       int      // Max number of results
	Format               // Format of response
}

func (q Query) buildURL() string {
	u, err := url.Parse(BASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	v := url.Values{}
	setInts(v, "event_id", q.EventId)
	setStrings(v, "keyword", q.KeywordAnd)
	setStrings(v, "keyword_or", q.KeywordOr)
	if q.Time != nil && len(q.Time) > 0 {
		ymd, ym := printTimeArray(q.Time)
		if len(ymd) > 0 {
			v.Set("ymd", ymd)
		}
		if len(ym) > 0 {
			v.Set("ym", ym)
		}
	}
	setStrings(v, "nickname", q.Participant)
	setStrings(v, "owner_nickname", q.Owner)
	setInts(v, "series_id", q.SeriesId)
	if q.Start > 0 {
		v.Set("start", fmt.Sprint(q.Start))
	}
	if q.Order > 0 {
		v.Set("order", fmt.Sprint(q.Order))
	}
	if q.Count > 0 {
		v.Set("count", fmt.Sprint(q.Count))
	}

	u.RawQuery = v.Encode()
	return u.String()
}

func setInts(p url.Values, k string, v []int) {
	if v != nil {
		for _, n := range v {
			p.Add(k, strconv.Itoa(n))
		}
	}
}

func setStrings(p url.Values, k string, v []string) {
	if v != nil {
		for _, e := range v {
			p.Add(k, e)
		}
	}
}

func printTimeArray(arr []Time) (string, string) {
	ymd := make([]string, 0)
	ym := make([]string, 0)
	for _, v := range arr {
		if v.Year > 0 && v.Month > 0 {
			if v.Date > 0 {
				ymd = append(ymd, fmt.Sprintf("%04d%02d%02d", v.Year, v.Month, v.Date))
			} else {
				ym = append(ym, fmt.Sprintf("%04d%02d", v.Year, v.Month))
			}
		}
	}
	return strings.Join(ymd, ","), strings.Join(ym, ",")
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
	url := q.buildURL()
	log.Printf("Executing URL: %s\n", url)
	res, err := http.Get(url)
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
