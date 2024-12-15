package fetch

import (
	"io"
	"net/http"
	"time"
)

type FetchArgs struct {
	Method        string
	Url           string
	Params        map[string]string
	Headers       map[string]string
	Body          io.Reader
	Jar           http.CookieJar
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Timeout       time.Duration
	Transport     http.RoundTripper
}

func Fetch(args FetchArgs) (*http.Response, error) {
	client := &http.Client{
		Timeout:       args.Timeout,
		Transport:     args.Transport,
		CheckRedirect: args.CheckRedirect,
		Jar:           args.Jar,
	}
	if len(args.Method) == 0 {
		args.Method = "GET"
	}
	req, err := http.NewRequest(args.Method, args.Url, args.Body)
	if err != nil {
		return nil, err
	}
	for k, v := range args.Headers {
		req.Header.Add(k, v)
	}
	if len(args.Params) != 0 {
		q := req.URL.Query()
		for k, v := range args.Params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
