package gr

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiURL = "https://www.goodreads.com/"

type Goodreads struct {
	Key    string
	Secret string
}

type Author struct {
	Id   int
	Name string
}

type GoodreadsResponse struct {
	Search Search
	Query  string
}

type Search struct {
	ResultsStart     int `xml:"results-start"`
	ResultsEnd       int `xml:"results-end"`
	TotalResults     int `xml:"total-results"`
	Source           string
	QueryTimeSeconds float64 `xml:"query-time-seconds"`
	Results          Results
}

type Results struct {
	Works []Work
}

type Work struct {
	BooksCount               int `xml:"books_count"`
	Id                       int
	OriginalPublicationDay   int      `xml:"original_publication_day",omitempty`
	OriginalPublicationMonth int      `xml:"original_publication_month",omitempty`
	OriginalPublicationYear  int      `xml:"original_publication_year",omitempty`
	RatingsCount             int      `xml:"ratings_count"`
	TextReviewsCount         int      `xml:"text_reviews_count"`
	AverageRating            float64  `xml:"average_rating"`
	BestBook                 BestBook `xml:"best_book"`
}

type BestBook struct {
	Id            int
	Title         string
	Author        Author
	ImageURL      string `xml:"image_url"`
	SmallImageURL string `xml:"small_image_url"`
}

func (g *Goodreads) getRequest(params map[string]string, endpoint string) ([]byte, error) {
	v := url.Values{}
	v.Set("key", g.Key)
	for key, val := range params {
		v.Set(key, val)
	}
	u := apiURL + endpoint + "?" + v.Encode()
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (g *Goodreads) SearchBooks(q string) (GoodreadsResponse, error) {
	p := map[string]string{"q": q}
	var gr GoodreadsResponse
	resp, err := g.getRequest(p, "search.xml")
	if err != nil {
		return GoodreadsResponse{}, err
	}
	err = xml.Unmarshal(resp, &gr)
	if err != nil {
		return GoodreadsResponse{}, err
	}
	return gr, nil
}
