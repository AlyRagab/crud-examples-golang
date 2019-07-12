package main

import (
	"context"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

var indexName = ""
var deleteIndx = ""

type information struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	JobTitle string `json:"jobtitle"`
}

var elastiClient *elastic.Client
var err error

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	// Connecting to Elasticsearch
	elastiClient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	check(err)
	fmt.Println("Connected Successfully !!")

	// Create Index
	creatingIndex()

	// Create or Insert Document into the Index
	insertDocument()

	// Delete Index
	// deleteIndex()

	// Search in Index
	search()
}

func creatingIndex() {
	mapping := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"properties":{
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				}
			
			}
		}
	}`
	ctx := context.Background()
	createIndex, err := elastiClient.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	check(err)

	if !createIndex.Acknowledged {
		fmt.Println("Not Created")
	}
	fmt.Println("Created Successfully !")
}
func insertDocument() {
	ctx := context.Background()
	doc := information{ID: 101, Name: "Aly Ragab", JobTitle: "DevOps Engineer"}

	ins, err := elastiClient.Index().Index(indexName).Id("1").BodyJson(doc).Do(ctx)
	check(err)
	fmt.Println("The Index :", indexName, "\n The ID is :", ins.Id)
	fmt.Println("Document is Inserted Successfully !!")

}

func search() {
	ctx := context.Background()
	query := elastic.NewTermQuery("Aly", "DevOps Engineer")
	result, err := elastiClient.Search().Index(indexName).Query(query).Pretty(true).Do(ctx)
	check(err)

	fmt.Println("The Query took :\n", result.TookInMillis, "The Number of hits :", result.Hits.TotalHits)

	// Getting the Data

}

func deleteIndex() {
	ctx := context.Background()
	deleteIndex, err := elastiClient.DeleteIndex(deleteIndx).Do(ctx)
	check(err)

	if deleteIndex.Acknowledged {
		fmt.Println("Index is Deleted Successfully !!")
	}
}

