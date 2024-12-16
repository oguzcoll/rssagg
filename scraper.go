package main

import (
	"context"
	"github.com/oguzcoll/rssagg/internal/database"
	"log"
	"sync"
	"time"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Starting scraping with %d concurrency and %s between requests", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Cannot get feeds: %s", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Cannot mark feed as fetched: %s", err)
		return
	}
	rssFeed, err := URLToFeed(feed.Url)
	if err != nil {
		log.Printf("Cannot get feed: %s", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		log.Println("Found Post: ", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
