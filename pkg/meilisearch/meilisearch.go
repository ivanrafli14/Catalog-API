package meilisearch

import (
	"fmt"
	"os"

	"github.com/ivanrafli14/CatalogAPI/entity"
	"github.com/meilisearch/meilisearch-go"
)

type Interface interface {
	SearchQuery(query string, products []*entity.Product) (any,any, error)
}

type MeilisearchClient struct {
	Client meilisearch.ServiceManager
}

func Init() Interface {
	ms_host := os.Getenv("MEILISEARCH_HOST")
	ms_port := os.Getenv("MEILISEARCH_PORT")
	ms_apikey := os.Getenv("MEILISEARCH_API_KEY")
	str := fmt.Sprintf("http://%s:%s", ms_host, ms_port)

	client := meilisearch.New(str, meilisearch.WithAPIKey(ms_apikey))
	return &MeilisearchClient{Client: client}
}

func (ms *MeilisearchClient) SearchQuery(query string, products []*entity.Product) (any,any, error) {
	var documents []map[string]interface{}
	for _, product := range products {
		doc := map[string]interface{}{
			"id":        product.ID,
			"name":      product.Name,
			"image_url": product.ImageUrl,
			"price":     product.Price,
			"stock":     product.Stock,
			"category":  product.Category.Name,
		}
		documents = append(documents, doc)
	}

	_, err := ms.Client.Index("product").AddDocuments(documents, "id")
	
	if err != nil {
		return nil,nil, err
	}

	index := ms.Client.Index("product")

	searchResult, err := index.Search(query, &meilisearch.SearchRequest{
		HitsPerPage: 100,
		Page:        1,
		Facets:      []string{"category"},
	})

	if err != nil {
		return nil,nil, err
	}
	return searchResult.Hits, searchResult.FacetDistribution, nil
}
