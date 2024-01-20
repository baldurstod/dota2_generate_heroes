package main

import (
	"strings"
	"github.com/baldurstod/vdf"
	"encoding/json"
)

type npcHeroes struct {
	heroesVDF *vdf.KeyValue
	heroes map[string]*hero
}

func (this *npcHeroes) init(dat []byte) {
	vdf := vdf.VDF{}
	root := vdf.Parse(dat)
	this.heroes = make(map[string]*hero)
	this.heroesVDF, _ = root.Get("DOTAHeroes")

	for _, hero := range this.heroesVDF.GetChilds() {
		if strings.HasPrefix(hero.Key, "npc_") {
			this.addHero(hero)
		}
	}
}

func (this *npcHeroes) addHero(kv *vdf.KeyValue) {
	h := &hero{npc: kv.Key, attributes: kv.Value.([]*vdf.KeyValue)}
	if h.isHero() {
		this.heroes[kv.Key] = h
	}
}

func (this *npcHeroes) MarshalJSON() ([]byte, error) {
	ret := []interface{}{}

	for _, val := range this.heroes {
		ret = append(ret, val)
	}

	return json.Marshal(ret)
}
