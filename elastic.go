package main

import (
	"gopkg.in/olivere/elastic.v3"
	"github.com/pborman/uuid"
)

func initESClient() *elastic.Client {
	elasticClient, err := elastic.NewClient(elastic.SetURL("localhost:9200"), elastic.SetSniff(true))

	if err != nil {
		return nil
	}

	return elasticClient
}

func IndexDocJSONBytes(client *elastic.Client, indexName, docType string, body string) (*elastic.IndexResponse, error) {
	resp, err := client.Index().Index(indexName).Type(docType).Id(uuid.New()).BodyString(body).Do()
	return resp, err
}