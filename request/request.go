package request

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Request struct{}

var ErrParseURL = errors.New("error to parse url string!")

func (r *Request) Get(rawUrl string) ([]byte, error) {
	url, err := validateUrl(rawUrl)
	if err != nil {
		return nil, ErrParseURL
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func validateUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	const defaultScheme = "https"
	if u.Scheme == "" {
		u.Scheme = defaultScheme
	}
	if u.Host == "" && u.Path == "" {
		return "", ErrParseURL
	}
	return u.String(), nil
}

func New() *Request {
	return &Request{}
}
