package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/baldurstod/vdf"
)

type language struct {
	lang   string
	tokens map[string]string
}

func (lg *language) init(path string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read file %s: %w", path, err)
	}

	v := vdf.VDF{}
	languageVdf := v.Parse(dat, nil)

	lang, err := languageVdf.Get("lang")
	if err != nil {
		panic("lang key not found")
	}
	language, err := lang.GetString("Language")
	if err != nil {
		panic("Language key not found")
	}

	tokens, err := lang.Get("Tokens")
	if err != nil {
		panic("Tokens key not found")
	}

	lg.lang = language
	lg.tokens = make(map[string]string)
	for _, val := range tokens.GetValue().(map[string][]*vdf.KeyValue) {
		for _, val2 := range val {
			lg.tokens[strings.ToLower(val2.Key)] = val2.GetValue().(string)
		}
	}
	return nil
}

func (lg *language) getToken(token string) (string, bool) {
	token = strings.TrimPrefix(token, "#")
	s, ok := lg.tokens[strings.ToLower(token)]
	return s, ok
}
