package main

import (
	"encoding/json"
	"github.com/baldurstod/vdf"
	"strconv"
)

type hero struct {
	npc        string
	attributes []*vdf.KeyValue
}

func (this *hero) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})

	ret["ID"] = this.npc
	ret["Name"] = getStringToken(this.npc + ":n")
	this.setIfExists(&this.attributes, &ret, "Model")
	this.setIfExists(&this.attributes, &ret, "Model1")
	this.setIfExists(&this.attributes, &ret, "Model2")
	this.setIfExists(&this.attributes, &ret, "Model3")
	this.setIfExists(&this.attributes, &ret, "NameAliases")
	this.setIfExists(&this.attributes, &ret, "HeroID")
	this.setIfExists(&this.attributes, &ret, "HeroOrderID")
	this.setIfExists(&this.attributes, &ret, "ModelScale")
	this.setIfExists(&this.attributes, &ret, "LoadoutScale")
	this.setIfExists(&this.attributes, &ret, "AttributePrimary")

	this.marshalSlots(&ret)
	this.marshalAdjectives(&ret)

	return json.Marshal(ret)
}

func (this *hero) getHeroOrderId() int {
	if s, ok := getStringAttribute(&this.attributes, "HeroOrderID"); ok {
		i, _ := strconv.Atoi(s)
		return i
	}
	return -1
}

func (this *hero) isHero() bool {
	_, ok := getStringAttribute(&this.attributes, "HeroOrderID")
	return ok
}

func (this *hero) marshalSlots(ret *map[string]interface{}) {
	slots := make(map[string]interface{})

	if itemslots, ok := getAttribute(&this.attributes, "ItemSlots"); ok {
		for _, kv := range itemslots {
			slotAttributes := kv.Value.([]*vdf.KeyValue)
			slot := make(map[string]interface{})

			this.setIfExists(&slotAttributes, &slot, "SlotIndex")
			this.setIfExists(&slotAttributes, &slot, "SlotName")
			this.setIfExists(&slotAttributes, &slot, "SlotText")
			this.setIfExists(&slotAttributes, &slot, "LoadoutPreviewMode")
			this.setIfExists(&slotAttributes, &slot, "DisplayInLoadout")
			if generatesUnits, ok := getAttribute(&slotAttributes, "GeneratesUnits"); ok {
				units := make(map[string]interface{})
				for _, kv := range generatesUnits {
					units[kv.Key] = kv.Value
				}
				slot["GeneratesUnits"] = units
			}

			slots[slot["SlotName"].(string)] = slot
		}
	}

	if len(slots) > 0 {
		(*ret)["ItemSlots"] = slots
	}
}
func (this *hero) marshalAdjectives(ret *map[string]interface{}) {
	adjectives := make(map[string]interface{})

	if adj, ok := getAttribute(&this.attributes, "Adjectives"); ok {
		for _, kv := range adj {
			adjectives[kv.Key] = kv.Value
		}
	}

	if len(adjectives) > 0 {
		(*ret)["Adjectives"] = adjectives
	}
}

func (this *hero) setIfExists(attributes *[]*vdf.KeyValue, ret *map[string]interface{}, attribute string) {
	if s, ok := getStringAttribute(attributes, attribute); ok {
		(*ret)[attribute] = getStringToken(s)
	}
}

func getStringAttribute(attributes *[]*vdf.KeyValue, attributeName string) (string, bool) {
	for _, kv := range *attributes {
		if kv.Key == attributeName {
			return kv.Value.(string), true
		}
	}
	return "", false
}

func getAttribute(attributes *[]*vdf.KeyValue, attributeName string) ([]*vdf.KeyValue, bool) {
	for _, kv := range *attributes {
		if kv.Key == attributeName {
			return kv.Value.([]*vdf.KeyValue), true
		}
	}
	return nil, false
}
