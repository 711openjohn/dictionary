package domain

type Example struct {
	Text        string `json:"text"`
	Translation string `json:"translation"`
}

type Meaning struct {
	Text        string    `json:"text"`
	Translation string    `json:"translation"`
	Examples    []Example `json:"examples"`
}

type Definition struct {
	POS               string  `json:"pos"`
	PronunciationLink string  `json:"pronunciationLink"`
	Meaning           Meaning `json:"meaning"`
}

type DictionaryPage struct {
	Symbol      string       `json:"symbol"`
	Definitions []Definition `json:"definitions"`
}
