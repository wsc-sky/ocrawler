package persist

import (
	"ocrawler/crawler/engine"
	"github.com/olivere/elastic"
	"ocrawler/crawler/persist"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save (item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	log.Printf("Item %v saved.", item)
	if err != nil {
		log.Printf("Error saving item %v: %v", item, err)
		return err
	}
	*result = "ok"
	return nil
}