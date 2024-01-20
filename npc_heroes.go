package main

import (
	"fmt"
	"strings"
	"github.com/baldurstod/vdf"
	"encoding/json"
)

/*type itemMap map[string]*item
type colorMap map[string]*color
type particleMap map[string]*particle*/

type npcHeroes struct {
	heroesVDF *vdf.KeyValue
	heroes map[string]*hero
	/*medals bool `default:false`
	Prefabs itemMap
	Items itemMap
	Colors colorMap
	Particles particleMap*/
}
/*
func (this *itemsGame) getItems() (*itemMap) {
	items := make(itemMap)

	for itemId, item := range this.Items {
		items[itemId] = item
	}
	return &items
}*/

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

/*	this.Prefabs = make(itemMap)
	this.Items = make(itemMap)
	this.Colors = make(colorMap)
	this.Particles = make(particleMap)

	if prefabs, ok := this.itemsVDF.Get("prefabs"); ok {
		for _, val := range prefabs.GetChilds() {
			var it = item{}
			if it.init(this, val) {
				this.Prefabs[it.Id] = &it
			}
		}
	}

	if items, ok := this.itemsVDF.Get("items"); ok {
		for _, val := range items.GetChilds() {
			var it = item{}
			if it.init(this, val) {
				this.Items[it.Id] = &it
			}
		}
	}

	if colors, ok := this.itemsVDF.Get("colors"); ok {
		for _, val := range colors.GetChilds() {
			var c = color{}
			if c.init(this, val) {
				this.Colors[c.Id] = &c
			}
		}
	}

	if particles, ok := this.itemsVDF.Get("attribute_controlled_attached_particles"); ok {
		for _, val := range particles.GetChilds() {
			var p = particle{}
			if p.init(this, val) {
				this.Particles[p.Id] = &p
			}
		}
	}

	for _, it := range this.Items {
		it.initPrefabs()
	}*/
}

func (this *npcHeroes) addHero(kv *vdf.KeyValue) {
	fmt.Println("add hero", kv.Key, kv.Value)
	this.heroes[kv.Key] = &hero{npc: kv.Key, attributes: kv.Value.([]*vdf.KeyValue)}
}

func (this *npcHeroes) MarshalJSON() ([]byte, error) {
	ret := []interface{}{}

	for _, val := range this.heroes {
		ret = append(ret, val)
	}

	return json.Marshal(ret)
}


func (this *npcHeroes) getHeroes() map[string]*hero {
	heroes := make(map[string]*hero)
/*
	for _, item := range this.Items {
		for _, npc := range item.getUsedByHeroes() {
			h, ok := heroes[npc]
			if !ok {
				h = &hero{npc: npc}
				heroes[npc] = h
			}

			h.addItem(item)
		}
	}*/
	return heroes
}

/*
func (this *itemsGame) getPrefab(prefabName string) *item {
	return this.Prefabs[prefabName]
}

func (this *itemsGame) getItemsPerHero() map[string]*hero {
	heroes := make(map[string]*hero)

	for _, item := range this.Items {
		for _, npc := range item.getUsedByHeroes() {
			h, ok := heroes[npc]
			if !ok {
				h = &hero{npc: npc}
				heroes[npc] = h
			}

			h.addItem(item)
		}
	}
	return heroes
}

func (this *itemsGame) getColors() []*color {
	ret := []*color{}

	for _, color := range this.Colors {
		ret = append(ret, color)
	}

	return ret
}

func (this *itemsGame) getParticles() []*particle {
	ret := []*particle{}

	for _, particle := range this.Particles {
		ret = append(ret, particle)
	}

	return ret
}
*/
