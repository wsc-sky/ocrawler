package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
	"ocrawler/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (chan engine.Item, error) {
	// TODO: try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("ItemSaver: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, item, index)
			if err != nil {
				log.Printf("ItemSaver: error saving item %v %v", item, err)
			}

		}
	}()
	return out, nil
}

func Save(client *elastic.Client, item engine.Item, index string) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}