package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HelloHandler))
	defer ts.Close()
	res, err := http.Get(ts.URL)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
