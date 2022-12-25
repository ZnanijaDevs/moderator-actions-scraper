package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func GetScraper(authToken string) *colly.Collector {
	c := colly.NewCollector()

	c.Async = true
	c.ParseHTTPErrorResponse = true

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("X-B-Token-Long", authToken)
		r.Headers.Set("Authorization", fmt.Sprintf("Basic %s", os.Getenv("BRAINLY_PROXY_AUTH_PASSWORD")))
	})

	return c
}