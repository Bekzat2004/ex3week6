package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("could not create HTTP GET request: %v", err)
	}

	rec := httptest.NewRecorder()
	index().ServeHTTP(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("received status code %v, expected: %v", res.StatusCode, http.StatusOK)
		body, _ := ioutil.ReadAll(res.Body)
		t.Fatalf("response body: %s", string(body))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %v", err)
	}

	if string(body) == "" {
		t.Errorf("response body should not be empty")
	}
}

func TestServerSetup(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", index())

	port := "8077"
	addr := ":" + port
	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			t.Fatalf("server failed: %v", err)
		}
	}()

	req, err := http.NewRequest(http.MethodGet, "http://localhost"+addr+"/", nil)
	if err != nil {
		t.Fatalf("could not create HTTP GET request: %v", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("could not send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("received status code %v, expected: %v", res.StatusCode, http.StatusOK)
		body, _ := ioutil.ReadAll(res.Body)
		t.Fatalf("response body: %s", string(body))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %v", err)
	}

	if string(body) == "" {
		t.Errorf("response body should not be empty")
	}

	err = server.Close()
	if err != nil {
		t.Fatalf("could not close server: %v", err)
	}
}

func TestExample1(t *testing.T) {
	if 1+0 != 1 {
		t.Error("expected 1+0 to equal 1")
	}
}

func TestExample2(t *testing.T) {
	if 2*1 != 2 {
		t.Error("expected 2*1 to equal 2")
	}
}

func TestExample3(t *testing.T) {
	if 3-0 != 3 {
		t.Error("expected 3-0 to equal 3")
	}
}

func TestExample4(t *testing.T) {
	if 4/1 != 4 {
		t.Error("expected 4/1 to equal 4")
	}
}

func TestExample5(t *testing.T) {
	if 5%5 != 0 {
		t.Error("expected 5%5 to equal 0")
	}
}

func TestExample6(t *testing.T) {
	if 6-0 != 6 {
		t.Error("expected 6-0 to equal 6")
	}
}

func TestExample7(t *testing.T) {
	if 7*1 != 7 {
		t.Error("expected 7*1 to equal 7")
	}
}

func TestExample8(t *testing.T) {
	if 8/1 != 8 {
		t.Error("expected 8/1 to equal 8")
	}
}

func TestExample9(t *testing.T) {
	if 9%3 != 0 {
		t.Error("expected 9%3 to equal 0")
	}
}

func TestExample10(t *testing.T) {
	if 10-0 != 10 {
		t.Error("expected 10-0 to equal 10")
	}
}

func TestFailFast(t *testing.T) {
	t.Fatalf("failing deliberately to demonstrate failfast behavior")
}
