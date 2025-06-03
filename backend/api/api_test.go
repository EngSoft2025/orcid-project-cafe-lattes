//go:build integration

package api

import (
	"io"
	"net/http"
	"testing"
)

func TestBiographySeiji(t *testing.T) {
	orcidId := "0000-0003-1574-0784"

	resp, err := http.Get("http://localhost:8080/api/searchBiography?orcid_id=" + orcidId)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	println("Response Status:", resp.Status)
	for _, values := range resp.Header {
		for _, value := range values {
			println("Header_data:", value)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	defer resp.Body.Close()
	println("Response Body:", string(body))
}

func TestWorkSeiji(t *testing.T) {
	orcidId := "0000-0003-1574-0784"
	resp, err := http.Get("http://localhost:8080/api/searchWork?orcid_id=" + orcidId)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	println("Response Status:", resp.Status)
	for _, values := range resp.Header {
		for _, value := range values {
			println("Header_data:", value)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	defer resp.Body.Close()
	println("Response Body:", string(body))
}

func TestRecord(t *testing.T) {
	orcidId := "0000-0003-3496-0256" // jorge
	resp, err := http.Get("http://localhost:8080/api/searchRecord?orcid_id=" + orcidId)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	println("Response Status:", resp.Status)
	for _, values := range resp.Header {
		for _, value := range values {
			println("Header_data:", value)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}
	defer resp.Body.Close()
	println("Response Body:", string(body))
}
