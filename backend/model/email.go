package model

type OrcidEmail struct {
	Emails []struct {
		Email      string `json:"email"`
		Verified   bool   `json:"verified"`
		Primary    bool   `json:"primary"`
		PutCode    uint64 `json:"put-code"`
		Visibility string `json:"visibility"`
	} `json:"email"`
}

// made the redundacy of the struct above to maintain the stardant over the models
type EmailData struct {
	Emails []struct {
		Email      string `json:"email"`
		Verified   bool   `json:"verified"`
		Primary    bool   `json:"primary"`
		PutCode    uint64 `json:"put-code"`
		Visibility string `json:"visibility"`
	} `json:"email"`
}
