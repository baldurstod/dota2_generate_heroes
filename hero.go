package main

import (
	_ "fmt"
	"strconv"
	"encoding/json"
	"github.com/baldurstod/vdf"
)

type hero struct {
	npc string
	attributes []*vdf.KeyValue
}

func (this *hero) MarshalJSON() ([]byte, error) {
	ret :=  make(map[string]interface{})

	ret["ID"] = this.npc
	ret["Name"] = getStringToken(this.npc + ":n")
	this.setIfExists(&this.attributes, &ret, "Model")
	this.setIfExists(&this.attributes, &ret, "NameAliases")
	this.setIfExists(&this.attributes, &ret, "HeroID")
	this.setIfExists(&this.attributes, &ret, "HeroOrderID")
	this.setIfExists(&this.attributes, &ret, "ModelScale")
	this.setIfExists(&this.attributes, &ret, "LoadoutScale")

	this.marshalSlots(&ret)

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

	if itemslots, ok := this.getAttribute("ItemSlots"); ok {
		for _, kv := range itemslots {
			slotAttributes := kv.Value.([]*vdf.KeyValue)
			slot := make(map[string]interface{})

			this.setIfExists(&slotAttributes, &slot, "SlotIndex")
			this.setIfExists(&slotAttributes, &slot, "SlotName")
			this.setIfExists(&slotAttributes, &slot, "SlotText")
			this.setIfExists(&slotAttributes, &slot, "LoadoutPreviewMode")
			this.setIfExists(&slotAttributes, &slot, "DisplayInLoadout")

			slots[slot["SlotName"].(string)] = slot
		}
	}

	if len(slots) > 0 {
		(*ret)["ItemSlots"] = slots
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

func (this *hero) getAttribute(attributeName string) ([]*vdf.KeyValue, bool) {
	for _, kv := range this.attributes {
		if kv.Key == attributeName {
			return kv.Value.([]*vdf.KeyValue), true
		}
	}
	return nil, false
}
