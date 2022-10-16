package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func GetScraper() *colly.Collector {
	c := colly.NewCollector()

	c.Async = true

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-B-Token-Long", os.Getenv("BRAINLY_AUTH_TOKEN"))
		r.Headers.Set("Authorization", fmt.Sprintf("Basic %s", os.Getenv("BRAINLY_PROXY_AUTH_PASSWORD")))
	})

	return c
}