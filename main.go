package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var passthruRequestHeaderKeys = [...]string{
	"Accept",
	"Accept-Encoding",
	"Accept-Language",
	"Cache-Control",
	"Cookie",
	"Referer",
	"User-Agent",
}

var passthruResponseHeaderKeys = [...]string{
	"Content-Encoding",
	"Content-Language",
	"Content-Type",
	"Cache-Control",
	"Date",
	"Etag",
	"Expires",
	"Last-Modified",
	"Location",
	"Server",
	"Vary",
}

func main() {
	handler := http.DefaultServeMux

	handler.HandleFunc("/", handleFunc)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		return
	}
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("--> %v %v\n", r.Method, r.URL)

	hh := http.Header{}
	for _, hk := range passthruRequestHeaderKeys {
		if hv, ok := r.Header[hk]; ok {
			hh[hk] = hv
		}
	}

	// Construct request to send to origin server
	rr := http.Request{
		Method:        r.Method,
		URL:           r.URL,
		Header:        hh,
		Body:          r.Body,
		ContentLength: r.ContentLength,
		Close:         r.Close,
	}

	resp, err := http.DefaultTransport.RoundTrip(&rr)
	if err != nil {
		http.Error(w, "Could not reach origin server", 500)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if false {
		fmt.Printf("<-- %v %+v\n", resp.Status, resp.Header)
	} else {
		fmt.Printf("<-- %v\n", resp.Status)
	}

	respH := w.Header()
	for _, hk := range passthruResponseHeaderKeys {
		if hv, ok := resp.Header[hk]; ok {
			respH[hk] = hv
		}
	}
	w.WriteHeader(resp.StatusCode)

	if resp.ContentLength > 0 {
		_, err := io.CopyN(w, resp.Body, resp.ContentLength)
		if err != nil {
			return
		}
	} else if resp.Close {
		for {
			if _, err := io.Copy(w, resp.Body); err != nil {
				break
			}
		}
	}
}
