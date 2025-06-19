package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
	"main/model"
)


func (r*Repository) GetCitationCount(doi string) (int, error) {

	if doi == "" { return -1, fmt.Errorf("empty doi")}
	// getting the citation count for each doi 
	url := fmt.Sprintf("https://api.semanticscholar.org/graph/v1/paper/DOI:%s?fields=citationCount", doi)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, fmt.Errorf("error fetching DOI %s: status %d", doi, resp.StatusCode)
	}

	var paper model.Paper
	err = json.NewDecoder(resp.Body).Decode(&paper)
	if err != nil {
		return -1, err
	}

	return paper.CitationCount, nil
}

// calculates the h index manually 
func (r*Repository) CalculateHIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	h := 0
	for i, c := range citations {
		if c >= i+1 {
			h = i + 1

		} else {
			break
		}
	}
	return h
}

///------
