package fynereader

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Title, Link string
	Items       []*FeedItem
}

type FeedItem struct {
	Title, Description string
	Link               string
}

func newFeed(href string) (*Feed, error) {
	fp := gofeed.NewParser()
	resp, err := fp.ParseURL(href)

	if err != nil {
		return nil, err
	}

	ret := &Feed{Title: resp.Title, Link: href}
	for _, item := range resp.Items {
		content := item.Content
		if content == "" {
			content = item.Description
		}
		retItem := &FeedItem{item.Title, content,
			fmt.Sprintf("%s#about", item.Link)}

		ret.Items = append(ret.Items, retItem)
	}
	return ret, nil
}
