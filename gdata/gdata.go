package main

import (
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"time"
	"errors"
)

type YEntry struct {
	Url string
	Thumb string
	Title string
	Description string
	Category string
	Keywords string
	Published time.Time
	Rating string
}

func ParseEntry(entry string) (ye YEntry, err error){
	urlRe := regexp.MustCompile(`<media:player url='(.*?)&amp;feature=youtube_gdata_player'/>`)
	thumbRe := regexp.MustCompile(`<media:thumbnail url='(.*?)'`)
	titleRe := regexp.MustCompile(`<media:title type='plain'>(.*?)</media:title>`)
	descriptionRe := regexp.MustCompile(`(?s)<media:description type='plain'>(.*?)</media:description>`)
	categoryRe := regexp.MustCompile(`<media:category.*?>(.*?)</media:category>`)	
	keywordsRe := regexp.MustCompile(`<media:keywords>(.*?)</media:keywords>`)
	publishedRe := regexp.MustCompile(`<published>(.*?)</published>`)	
	ratingRe := regexp.MustCompile(`<gd:rating average='(.*?)' .*? numRaters='(.*?)'`)	
	if url := urlRe.FindStringSubmatch(entry); len(url) == 2 {
		ye.Url = url[1]
	} else {
		err = errors.New("Malformed entry: url")
	}
	if thumb := thumbRe.FindStringSubmatch(entry); len(thumb) == 2 {
		ye.Thumb = thumb[1]
	} else {
		err = errors.New("Malformed entry: thumb")
	}
	if title := titleRe.FindStringSubmatch(entry); len(title) == 2 {
		ye.Title = title[1]
	} else {
		err = errors.New("Malformed entry: title")
	}
	if description := descriptionRe.FindStringSubmatch(entry); len(description) == 2 {
		ye.Description = description[1]
	} else {
		err = errors.New("Malformed entry: description")
	}
	if category := categoryRe.FindStringSubmatch(entry); len(category) == 2 {
		ye.Category = category[1]
	} else {
		err = errors.New("Malformed entry: category")
	}
	if keywords := keywordsRe.FindStringSubmatch(entry); len(keywords) == 2 {
		ye.Keywords = keywords[1]
	} else {
		err = errors.New("Malformed entry: keywords")
	}	
	if published := publishedRe.FindStringSubmatch(entry); len(published) == 2 {
		ye.Published, err = time.Parse(time.RFC3339, published[1])
	} else {
		err = errors.New("Malformed entry: published")
	}
	if rating := ratingRe.FindStringSubmatch(entry); len(rating) == 3 {
		ye.Rating = fmt.Sprintf("Rating: %s of 5 stars<br/>%s Votes", rating[1], rating[2])
	} 				
	log.Print(err)
	return 
}

func ParseFeed(feed string) (yentries []YEntry) {
	resp, _ := http.Get(feed)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)	
	entryRe := regexp.MustCompile(`(?s)<entry.*?>(.*?)</entry>`)
	entries := entryRe.FindAllString(string(body), 50)	
	for _, e := range entries {				
		if ye, err := ParseEntry(e); err == nil {		
			yentries = append(yentries, ye)
		}
	}
	return
}

func main(){
	//yes := ParseFeed("https://gdata.youtube.com/feeds/api/standardfeeds/recently_featured")
	//yes := ParseFeed("https://gdata.youtube.com/feeds/api/videos?q=surfing")
	yes := ParseFeed("https://gdata.youtube.com/feeds/api/videos?q=elvis&v=2&max-results=10&category=Music&orderby=published")
	log.Print(yes)	
}
