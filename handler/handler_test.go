package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/laracarvalho/goboards/router"
	"github.com/stretchr/testify/assert"
)

func TestCreateListingHandler(t *testing.T) {
	r := router.Start()

	w := httptest.NewRecorder()

	tests := []struct {
		payload []byte
	}{
		{[]byte("")},
		{[]byte(`{
			"company": "MyCompany",
			"location": "US",
			"remote": true,
			"link": "localhost",
			"salary": 80000,
			"description": "We're looking for a good Sr. Developer experienced with Go!",
		}`)},
		{[]byte("foobar")},
	}

	for _, tt := range tests {
		body := bytes.NewReader(tt.payload)
		req, _ := http.NewRequest("POST", "/api/v1/listing", body)
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	}
}

func TestUpdateListingHandler(t *testing.T) {
	r := router.Start()

	w := httptest.NewRecorder()

	tests := []struct {
		payload []byte
	}{
		{[]byte("")},
	}

	for _, tt := range tests {
		body := bytes.NewReader(tt.payload)
		req, _ := http.NewRequest("PUT", "/api/v1/listing", body)
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	}
}

func TestDeleteListingHandler(t *testing.T) {
	r := router.Start()

	w := httptest.NewRecorder()

	tests := []struct {
		payload []byte
	}{
		{[]byte("")},
	}

	for _, tt := range tests {
		body := bytes.NewReader(tt.payload)
		req, _ := http.NewRequest("DELETE", "/api/v1/listing", body)
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	}
}

func TestShowListingHandler(t *testing.T) {
	r := router.Start()

	w := httptest.NewRecorder()

	tests := []struct {
		payload []byte
	}{
		{[]byte("")},
	}

	for _, tt := range tests {
		body := bytes.NewReader(tt.payload)
		req, _ := http.NewRequest("GET", "/api/v1/listing?id=", body)
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	}
}
