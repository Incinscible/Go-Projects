package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Technology struct {
	Tech      string   `json:"Tech"`
	Matchers  []string `json:"Matchers"`
	FullMatch bool     `json:"FullMatch"`
}

func loadTechnologiesMatchers(filename string) ([]Technology, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var technologies []Technology
	err = json.Unmarshal(data, &technologies)
	if err != nil {
		return nil, err
	}
	return technologies, nil
}

func main() {
	args := os.Args
	fmt.Println("yo le gang :", args)

	if len(args) < 2 {
		fmt.Println("trop d'arguments envoyés mon gâté")
	}

	url := args[1]
	fmt.Println("[+] recherche de techno sur :", args[1])

	technologies, err := loadTechnologiesMatchers("matchers.json")
	if err != nil {
		fmt.Println("error reading the file")
	}

	fmt.Println(technologies)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("bad url: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body: ", err)
	}
	// fmt.Println(string(body))
	bodyStr := string(body)

	var potentialTechs []string

	for _, tech := range technologies {
		cptMatch := 0
		for _, match := range tech.Matchers {
			if strings.Contains(bodyStr, match) {
				fmt.Println("- Found something with: ", tech.Tech)
				fmt.Println("- With this match: ", match)
				fmt.Println("")
				cptMatch++
			}
		}
		if cptMatch == len(tech.Matchers) {
			potentialTechs = append(potentialTechs, tech.Tech)
		}
	}
	fmt.Println("...")
	fmt.Println("[+] All potential techs found: ", potentialTechs)

}
