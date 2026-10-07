package main

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/baldurstod/vdf"
)

type npcHeroes struct {
	heroesVDF *vdf.KeyValue
	heroes    map[string]*hero
	units     map[string]*unit
}

func (npcHeroes *npcHeroes) init(heroes []byte, units []byte, getFileContent vdf.GetFileContent) {
	npcHeroes.initHeroes(heroes, getFileContent)
	npcHeroes.initUnits(units)
}

func (npcHeroes *npcHeroes) initHeroes(heroes []byte, getFileContent vdf.GetFileContent) {
	vdf := vdf.VDF{}
	root := vdf.Parse(heroes, getFileContent)
	npcHeroes.heroes = make(map[string]*hero)
	npcHeroes.heroesVDF, _ = root.Get("DOTAHeroes")

	for _, hero := range npcHeroes.heroesVDF.GetChilds() {
		if strings.HasPrefix(hero.Key, "npc_") {
			npcHeroes.addHero(hero)
		}
	}
}

func (npcHeroes *npcHeroes) initUnits(units []byte) {
	vdf := vdf.VDF{}
	root := vdf.Parse(units, nil)
	npcHeroes.units = make(map[string]*unit)
	npcHeroes.heroesVDF, _ = root.Get("DOTAUnits")

	for _, unit := range npcHeroes.heroesVDF.GetChilds() {
		if strings.HasPrefix(unit.Key, "npc_dota_") {
			npcHeroes.addUnit(unit)
		}
	}
}

func (npcHeroes *npcHeroes) addHero(kv *vdf.KeyValue) {
	h := &hero{npc: kv.Key, attributes: kv.GetValue().(map[string][]*vdf.KeyValue)}
	if h.isHero() {
		npcHeroes.heroes[kv.Key] = h
	}
}

func (npcHeroes *npcHeroes) addUnit(kv *vdf.KeyValue) {
	h := &unit{npc: kv.Key, attributes: kv.GetValue().(map[string][]*vdf.KeyValue)}
	npcHeroes.units[kv.Key] = h
}

func (npcHeroes *npcHeroes) MarshalJSON() ([]byte, error) {
	heroes := []interface{}{}
	units := map[string]interface{}{}
	ret := map[string]interface{}{
		"heroes": &heroes,
		"units":  &units,
	}

	orderToNPC := make(map[int]string)

	// Sort keys
	keys := make([]int, 0, len(npcHeroes.heroes))
	for k, h := range npcHeroes.heroes {
		heroOrderId := h.getHeroOrderId()
		orderToNPC[heroOrderId] = k
		keys = append(keys, heroOrderId)
	}
	sort.Ints(keys)

	for _, k := range keys {
		heroes = append(heroes, npcHeroes.heroes[orderToNPC[k]])
	}

	npcHeroes.marshalUnits(&units)

	return json.Marshal(ret)
}

func (npcHeroes *npcHeroes) marshalUnits(units *map[string]interface{}) {
	for k, h := range npcHeroes.units {
		(*units)[k] = h
	}
}
