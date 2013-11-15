package gr

type Goodreads struct {
	Key    string
	Secret string
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
	return resp, nil
}
