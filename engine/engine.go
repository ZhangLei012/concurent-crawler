package engine

import (
	"crawler/fetcher"
	"log"
	"time"
)

//SimpleEngine is a single goroutine engine
type SimpleEngine struct {
}

// Run will continuously handler requests
// if the request queue is not empty, and will add new request to the queue
func (e *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	time.Sleep(time.Minute)
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)

		parserResult, err := worker(r)
		if err != nil {
			log.Printf("Worker: error handling request %v : %v", r, err)
			continue
		}
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %s\n", item)
		}
		log.Printf("len(requests) = %v", len(requests))
	}
	log.Println("Run finish")
}

// Worker is to fetch URL content and parse content
func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s\n", r.URL)
	content, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v\n", r.URL, err)
		return ParserResult{}, err
	}

	return r.ParserFunc(content), nil
}

// MultiRun is to fetch and parse with goroutine
func (e *SimpleEngine) MultiRun(seeds ...Request) {
	requests := make(chan Request, 1)
	for _, r := range seeds {
		requests <- r
	}
	contentCh := make(chan []byte)

	for req := range requests {
		r := req
		log.Printf("Fetching %s", r.URL)
		go func() {
			content, err := fetcher.Fetch(r.URL)
			if err != nil {
				log.Printf("Fetcher: error fetching url %s: %v", r.URL, err)
				return
			}
			contentCh <- content
		}()
		go func() {
			for content := range contentCh {
				parserResult := r.ParserFunc(content)
				for _, re := range parserResult.Requests {
					requests <- re
				}

				for _, item := range parserResult.Items {
					log.Printf("Got item %s\n", item)
				}
				log.Printf("len(requests) = %v", len(requests))
			}
		}()
	}
	time.Sleep(time.Hour)
	log.Println("Run finish")
}
