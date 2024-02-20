package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Item struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

const tempJson = `{"5447a9cd4bdc2dbd208b4567": {"id": "5447a9cd4bdc2dbd208b4567","name": "Colt M4A1 5.56x45 assault rifle","shortName": "M4A1"}, "544fb45d4bdc2dee738b4568": {"id": "544fb45d4bdc2dee738b4568","name": "Salewa first aid kit","shortName": "Salewa"}}`

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	data, err := os.ReadFile("items.en.json")
	if err != nil {
		log.Fatal(err)
	}

	var items map[string]Item
	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(items)

	// ctx := context.Background()
	// client := graphql.NewClient("https://api.tarkov.dev/graphql", http.DefaultClient)
	// resp, err := getItems(ctx, client, "GEN M3")
	// fmt.Println(resp.Items[0].Name, err)
	// fmt.Println(resp.Items[0].ShortName, err)
	// fmt.Println(resp.Items[0].Avg24hPrice, err)
}
