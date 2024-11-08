package main

import (
	"encoding/json"
	"strconv"

	"github.com/baldurstod/vdf"
)

type hero struct {
	npc        string
	attributes []*vdf.KeyValue
}

func (h *hero) MarshalJSON() ([]byte, error) {
	ret := make(map[string]interface{})

	ret["ID"] = h.npc
	ret["Name"] = getStringToken(h.npc + ":n")
	h.setIfExists(&h.attributes, &ret, "Model")
	h.setIfExists(&h.attributes, &ret, "Model1")
	h.setIfExists(&h.attributes, &ret, "Model2")
	h.setIfExists(&h.attributes, &ret, "Model3")
	h.setIfExists(&h.attributes, &ret, "NameAliases")
	h.setIfExists(&h.attributes, &ret, "HeroID")
	h.setIfExists(&h.attributes, &ret, "HeroOrderID")
	h.setIfExists(&h.attributes, &ret, "ModelScale")
	h.setIfExists(&h.attributes, &ret, "LoadoutScale")
	h.setIfExists(&h.attributes, &ret, "AttributePrimary")

	h.marshalSlots(&ret)
	h.marshalAdjectives(&ret)

	return json.Marshal(ret)
}

func (h *hero) getHeroOrderId() int {
	if s, ok := getStringAttribute(&h.attributes, "HeroOrderID"); ok {
		i, _ := strconv.Atoi(s)
		return i
	}
	if s, ok := getStringAttribute(&h.attributes, "HeroID"); ok {
		i, _ := strconv.Atoi(s)
		return i
	}
	return 1000
}

func (h *hero) isHero() bool {
	_, ok := getStringAttribute(&h.attributes, "HeroID")
	return ok
}

func (h *hero) marshalSlots(ret *map[string]interface{}) {
	slots := make(map[string]interface{})

	if itemslots, ok := getAttribute(&h.attributes, "ItemSlots"); ok {
		for _, kv := range itemslots {
			slotAttributes := kv.Value.([]*vdf.KeyValue)
			slot := make(map[string]interface{})

			h.setIfExists(&slotAttributes, &slot, "SlotIndex")
			h.setIfExists(&slotAttributes, &slot, "SlotName")
			h.setIfExists(&slotAttributes, &slot, "SlotText")
			h.setIfExists(&slotAttributes, &slot, "LoadoutPreviewMode")
			h.setIfExists(&slotAttributes, &slot, "DisplayInLoadout")
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
func (h *hero) marshalAdjectives(ret *map[string]interface{}) {
	adjectives := make(map[string]interface{})

	if adj, ok := getAttribute(&h.attributes, "Adjectives"); ok {
		for _, kv := range adj {
			adjectives[kv.Key] = kv.Value
		}
	}

	if len(adjectives) > 0 {
		(*ret)["Adjectives"] = adjectives
	}
}

func (h *hero) setIfExists(attributes *[]*vdf.KeyValue, ret *map[string]interface{}, attribute string) {
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
