package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	ctx := context.Background()
	client := graphql.NewClient("https://api.tarkov.dev/graphql", http.DefaultClient)
	resp, err := getItems(ctx, client, "GEN M3")
	fmt.Println(resp.Items[0].Name, err)
	fmt.Println(resp.Items[0].ShortName, err)
	fmt.Println(resp.Items[0].Avg24hPrice, err)
}
