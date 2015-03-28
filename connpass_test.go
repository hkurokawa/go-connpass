package connpass

import (
	"io/ioutil"
	"testing"
	"net/url"
)

func TestParse(t *testing.T) {
	file := "sample.json"
	jsonBlob, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to open %v. Reason:%s.", file, err)
	}
	result, err := parse(jsonBlob)
	if err != nil {
		t.Fatalf("Failed to parse the JSON. Reason:%s.", err)
	}

	if result.Returned != 10 {
		t.Errorf("Unexpected number of returned events: %v.", result.Returned)
	}
}

func TestDefaultQuery(t *testing.T) {
	query := Query{}
	s := query.buildURL()
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("Invalid URL: %s.", s, err)
	}
	if u.RawQuery != "" {
		t.Errorf("Query should be empty: %s", s)
	}
}

func TestBasicQuery(t *testing.T) {
	query := Query{Start: 10, Count: 100, Order: CREATE}
	s := query.buildURL()
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("Invalid URL: %s.", s, err)
	}
	assertParam(t, u, "start", "10")
	assertParam(t, u, "count", "100")
	assertParam(t, u, "order", "3")
}

func TestAdvancedQuery(t *testing.T) {
	query := Query{}
	query.KeywordAnd = []string{"test", "http"}
	query.KeywordOr = []string{"go", "golang"}
	query.Time = []Time{Time{Year: 2015, Month: 4}, Time{Year: 2015, Month: 3}, Time{Year: 2015, Month: 6, Date: 26}}
	query.Participant = []string{"hoge", "fuga"}
	query.Owner = []string{"foo"}
	query.SeriesId = []int{999}
	query.EventId = []int{42, 88}

	s := query.buildURL()
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("Invalid URL: %s.", s, err)
	}
	assertParam(t, u, "keyword", "test", "http")
	assertParam(t, u, "keyword_or", "go", "golang")
	assertParam(t, u, "ym", "201504,201503")
	assertParam(t, u, "ymd", "20150626")
	assertParam(t, u, "nickname", "hoge", "fuga")
	assertParam(t, u, "owner_nickname", "foo")
	assertParam(t, u, "series_id", "999")
	assertParam(t, u, "event_id", "42", "88")
}

func assertParam(t *testing.T, u *url.URL, k string, e ...string) {
	v := u.Query()[k]
	if !equalStrings(v, e) {
		t.Errorf("Value for the Query param %s is invalid. '%s' ≠ '%s'.", u, v, e)
	}
}

func equalStrings(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}