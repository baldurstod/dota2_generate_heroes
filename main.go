package main

import (
	"encoding/json"
	"os"
	_ "log"
	"flag"
	"fmt"
	"path"
)

var lg language
var dota language
var languages []*language

func main() {
	var lang string
	var outputFolder string
	var inputFolder string
	var resourceFolder string

	flag.StringVar(&lang, "l", "english", "Language")
	flag.StringVar(&outputFolder, "o", "", "Output folder")
	flag.StringVar(&inputFolder, "i", "", "Input folder")
	flag.StringVar(&resourceFolder, "r", "", "Resource folder")
	flag.Parse()

	if inputFolder == "" {
		fmt.Println("No input folder provided. Use the flag -i")
		os.Exit(1)
	}
	if resourceFolder == "" {
		fmt.Println("No resource folder provided. Use the flag -r")
		os.Exit(1)
	}
	if outputFolder == "" {
		fmt.Println("No output folder provided. Use the flag -o")
		os.Exit(1)
	}

	lg = language{}
	lg.init(path.Join(resourceFolder, "abilities_" + lang + ".txt"))

	dota = language{}
	dota.init(path.Join(resourceFolder, "dota_" + lang + ".txt"))

	languages = []*language{&lg, &dota}

	heroes := npcHeroes{}
	npcHeroesDatas, _ := os.ReadFile(path.Join(inputFolder, "npc_heroes.txt"))
	heroes.init(npcHeroesDatas)

	j, _ := json.MarshalIndent(&heroes, "", "\t")
	os.WriteFile(path.Join(outputFolder, "heroes.json"), j, 0666)

/*
	heroes := ig.getItemsPerHero()

	for npc, hero := range heroes {
		j, _ := json.MarshalIndent(hero, "", "\t")
		os.WriteFile(path.Join(outputFolder, npc + ".json"), j, 0666)
	}

	j, _ := json.MarshalIndent(ig.getColors(), "", "\t")
	os.WriteFile(path.Join(outputFolder, "colors.json"), j, 0666)

	p, _ := json.MarshalIndent(ig.getParticles(), "", "\t")
	os.WriteFile(path.Join(outputFolder, "particles.json"), p, 0666)
	*/
}

func getStringToken(token string) string {
	for _, language := range languages {
		s, exist := language.getToken(token)

		if (exist) {
			return s
		}
	}
	return token
}
