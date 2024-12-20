package main

import (
	"os"
	"strings"

	"github.com/baldurstod/vdf"
)

type language struct {
	lang   string
	tokens map[string]string
}

func (lg *language) init(path string) {
	dat, _ := os.ReadFile(path)
	v := vdf.VDF{}
	languageVdf := v.Parse(dat)

	lang, ok := languageVdf.Get("lang")
	if !ok {
		panic("lang key not found")
	}
	language, ok := lang.GetString("Language")
	if !ok {
		panic("Language key not found")
	}

	tokens, ok := lang.Get("Tokens")
	if !ok {
		panic("Tokens key not found")
	}

	lg.lang = language
	lg.tokens = make(map[string]string)
	for _, val := range tokens.Value.([]*vdf.KeyValue) {
		lg.tokens[strings.ToLower(val.Key)] = val.Value.(string)
	}
}

func (lg *language) getToken(token string) (string, bool) {
	token = strings.TrimPrefix(token, "#")
	s, ok := lg.tokens[strings.ToLower(token)]
	return s, ok
}
