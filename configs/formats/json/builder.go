package jsonConfig

import (
	"github.com/adverax/metacrm.kernel/access/fetchers/maps/json"
	"github.com/adverax/metacrm.kernel/configs"
)

func NewFileLoaderBuilder() *configs.FileLoaderBuilder {
	return configs.NewFileLoaderBuilder().
		WithSourceBuilder(
			func(fetcher configs.Fetcher) configs.Source {
				return jsonFetcher.New(fetcher)
			},
		)
}
