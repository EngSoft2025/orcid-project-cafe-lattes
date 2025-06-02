package model

type OrcidKeyword struct {
	Keyword []struct {
		Content    string `json:"content"`
		Visibility string `json:"visibility"`
		PutCode    uint64 `json:"put-code"`
	} `json:"keyword"`
}

// redundancy of the above struct to mantain a standart over the models
type KeywordData struct {
	Keywords []struct {
		Content    string `json:"content"`
		Visibility string `json:"visibility"`
		PutCode    uint64 `json:"put-code"`
	} `json:"keyword"`
}
