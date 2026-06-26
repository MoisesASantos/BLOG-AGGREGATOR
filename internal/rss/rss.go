package rss

import (
	"context"
	"encoding/xml"
	"net/http"
	"errors"
	"io"
	"fmt"
	"time"
	"html"
	//"bytes"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var data RSSFeed

	client := &http.Client {
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
    	return &RSSFeed{}, fmt.Errorf("Error to make a request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Pau dos paus")
    	return &RSSFeed{}, fmt.Errorf("Error to make a request: %w", err)
	}
	defer res.Body.Close()
	fmt.Println("Status:", res.Status)


	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error to try read the body response")
		return &RSSFeed{}, errors.New("Error to try read the body response")
	}
	
	/*fmt.Println(string(body[:300]))
	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.Strict = false*/
	err = xml.Unmarshal(body, &data)
	if err != nil {
    	return &RSSFeed{}, fmt.Errorf("xml unmarshal failed: %w", err)
	}

	data.Channel.Title = html.UnescapeString(data.Channel.Title)
	data.Channel.Description = html.UnescapeString(data.Channel.Description)
	for i := range data.Channel.Item {
		data.Channel.Item[i].Title = html.UnescapeString(data.Channel.Item[i].Title)
		data.Channel.Item[i].Description = html.UnescapeString(data.Channel.Item[i].Description)
	}

	return &data, nil
}
