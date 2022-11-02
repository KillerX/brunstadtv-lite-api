package model

import "encoding/json"

type TranslatedString map[string]string

func (ts *TranslatedString) Scan(lang string) error {
	return json.Unmarshal([]byte(lang), ts)
}

func (ts TranslatedString) Get(lang string) string {
	if s, ok := ts[lang]; ok {
		return s
	}

	return ""
}
