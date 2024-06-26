package main

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
)

var cli *elasticsearch.Client

func init() {
	var err error
	cli, err = elasticsearch.NewClient(elasticsearch.Config{
		//Addresses: []string{"https://e210ce1769e34271aa23971cc6ff433b.us-central1.gcp.cloud.es.io"},
		CloudID: "be82d87231024e399ed55350b992b7f6:dXMtY2VudHJhbDEuZ2NwLmNsb3VkLmVzLmlvJGUyMTBjZTE3NjllMzQyNzFhYTIzOTcxY2M2ZmY0MzNiJDlkZGQ3Y2Q2ODcyYTRmZGQ5NzZmNWFhYTdhOWE1YWUx",
		APIKey:  "TmZreEpaQUJmV3phNjQ5dEkyS046dWNwUEVjWUpTQXV0cmJEaUJFX0ktQQ==",
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	res, err := cli.Indices.Create("my_index")
	if err != nil {
		return err
	}
	println(jsonMarshal(res))

	return nil
}

func jsonMarshal(t any) string {
	bytes, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}
