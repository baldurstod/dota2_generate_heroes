package main

import (
	"fmt"
	"encoding/json"
	"github.com/baldurstod/vdf"
)

type hero struct {
	npc string
	attributes []*vdf.KeyValue
}

func (this *hero) MarshalJSON() ([]byte, error) {
	ret :=  make(map[string]interface{})

	fmt.Println("MarshalJSON", this.npc)

	ret["id"] = this.npc
	this.setIfExists(&ret, "Model")
	this.setIfExists(&ret, "NameAliases")


	//ret["Model"] = this.getStringAttribute("Model")

	/*if s, ok := this.getStringAttribute("item_name"); ok {
		ret["name"] = getStringToken(s)
	}*/

	ret["Name"] = getStringToken(this.npc + ":n")

	//$hero->name = getStringToken($heroId.':n');

	//NameAliases


	/*if s, ok := this.getStringAttribute("item_name"); ok {
		ret["name"] = getStringToken(s)
	}*/

	/*for _, kv := range this.attributes {
		//ret = append(kv.Key, kv.Value)
		fmt.Println(kv.Key, kv.Value)
		ret[kv.Key] = kv.Value
	}*/

	return json.Marshal(ret)
}

func (this *hero) setIfExists(ret *map[string]interface{}, attribute string) {
	if s, ok := this.getStringAttribute(attribute); ok {
		(*ret)[attribute] = getStringToken(s)
	}
}

func (this *hero) getStringAttribute(attributeName string) (string, bool) {
	for _, kv := range this.attributes {
		if kv.Key == attributeName {
			return kv.Value.(string), true
		}
	}

/*
	if s, ok := this.attributes.GetString(attributeName); ok {
		return s, true
	}*/
	return "", false
}
