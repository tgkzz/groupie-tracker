package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// BenchmarkHealthCheck tests the performance of the HealthCheck handler.
func BenchmarkHealthCheck(b *testing.B) {
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		HealthCheck(response, request)
	}
}

// BenchmarkIndexHandler tests the performance of the IndexHandler handler.
func BenchmarkIndexHandler(b *testing.B) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		IndexHandler(response, request)
	}
}

// BenchmarkGroupHandler tests the performance of the GroupHandler handler.
func BenchmarkGroupHandler(b *testing.B) {
	request, _ := http.NewRequest("GET", "/groups/1", nil)
	response := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		GroupHandler(response, request)
	}
}

func BenchmarkFilterHandler(b *testing.B) {
	form := url.Values{}
	form.Add("creation_date_from", "1990")
	form.Add("creation_date_to", "2000")
	form.Add("firstAlbum_from", "1992")
	form.Add("firstAlbum_to", "1999")
	form.Add("members[]", "John")
	form.Add("members[]", "Paul")
	form.Add("location", "Liverpool")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		request, _ := http.NewRequest("POST", "/filter", strings.NewReader(form.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		FilterHandler(response, request)
	}
}

func BenchmarkSearchHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		request, _ := http.NewRequest("GET", "/search?searchText=Beatles", nil)
		response := httptest.NewRecorder()

		SearchHandler(response, request)
	}
}
