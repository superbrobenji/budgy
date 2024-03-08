package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
	services "github.com/superbrobenji/budgy/core/service/createItemTest"
)

type book struct {
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func getItems() (*aggregate.Item, error) {
	itemService, err := services.NewItemService()
	if err != nil {
		return nil, err
	}
	item, error := itemService.CreateItem(uuid.New())
	if error != nil {
		return nil, error
	}

	return item, nil
}

func main() {
	lambda.Start(getItems)
}
