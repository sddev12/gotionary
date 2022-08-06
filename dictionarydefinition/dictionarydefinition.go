package dictionarydefinition

type phonetics struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

type definitions struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type meanings struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []definitions `json:"definitions"`
}

type Word struct {
	Word      string      `json:"word"`
	Phonetics []phonetics `json:"phonetics"`
	Meanings  []meanings  `json:"meanings"`
}
