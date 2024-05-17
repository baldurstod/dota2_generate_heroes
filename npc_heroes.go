package main

import (
	"encoding/json"
	"github.com/baldurstod/vdf"
	"sort"
	"strings"
)

type npcHeroes struct {
	heroesVDF *vdf.KeyValue
	heroes    map[string]*hero
	units     map[string]*unit
}

func (this *npcHeroes) init(heroes []byte, units []byte) {
	this.initHeroes(heroes)
	this.initUnits(units)
}

func (this *npcHeroes) initHeroes(heroes []byte) {
	vdf := vdf.VDF{}
	root := vdf.Parse(heroes)
	this.heroes = make(map[string]*hero)
	this.heroesVDF, _ = root.Get("DOTAHeroes")

	for _, hero := range this.heroesVDF.GetChilds() {
		if strings.HasPrefix(hero.Key, "npc_") {
			this.addHero(hero)
		}
	}
}

func (this *npcHeroes) initUnits(units []byte) {
	vdf := vdf.VDF{}
	root := vdf.Parse(units)
	this.units = make(map[string]*unit)
	this.heroesVDF, _ = root.Get("DOTAUnits")

	for _, unit := range this.heroesVDF.GetChilds() {
		if strings.HasPrefix(unit.Key, "npc_dota_") {
			this.addUnit(unit)
		}
	}
}

func (this *npcHeroes) addHero(kv *vdf.KeyValue) {
	h := &hero{npc: kv.Key, attributes: kv.Value.([]*vdf.KeyValue)}
	if h.isHero() {
		this.heroes[kv.Key] = h
	}
}

func (this *npcHeroes) addUnit(kv *vdf.KeyValue) {
	h := &unit{npc: kv.Key, attributes: kv.Value.([]*vdf.KeyValue)}
	this.units[kv.Key] = h
}

func (this *npcHeroes) MarshalJSON() ([]byte, error) {
	heroes := []interface{}{}
	units := map[string]interface{}{}
	ret := map[string]interface{}{
		"heroes": &heroes,
		"units":  &units,
	}

	orderToNPC := make(map[int]string)

	// Sort keys
	keys := make([]int, 0, len(this.heroes))
	for k, h := range this.heroes {
		heroOrderId := h.getHeroOrderId()
		orderToNPC[heroOrderId] = k
		keys = append(keys, heroOrderId)
	}
	sort.Ints(keys)

	for _, k := range keys {
		heroes = append(heroes, this.heroes[orderToNPC[k]])
	}

	this.marshalUnits(&units)

	return json.Marshal(ret)
}

func (this *npcHeroes) marshalUnits(units *map[string]interface{}) {
	for k, h := range this.units {
		(*units)[k] = h
	}
}
