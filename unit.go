package main

import (
	"encoding/json"
	"github.com/baldurstod/vdf"
)

type unit struct {
	npc        string
	attributes []*vdf.KeyValue
}

func (u *unit) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})

	u.setIfExists(&u.attributes, &ret, "Model")
	u.setIfExists(&u.attributes, &ret, "include_keys_from")
	ret["name"] = getStringToken(u.npc)

	return json.Marshal(ret)
}

func (u *unit) setIfExists(attributes *[]*vdf.KeyValue, ret *map[string]interface{}, attribute string) {
	if s, ok := getStringAttribute(attributes, attribute); ok {
		(*ret)[attribute] = s
	}
}
