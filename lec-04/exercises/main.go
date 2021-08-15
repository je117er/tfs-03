package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"./utils"

	"golang.org/x/net/html"
)

var seen = make(map[string]bool)
var movies []utils.Movie

func main() {

	url := os.Args[1]

	movieExplorer := make(chan string)
	movieToJson := make(chan string)

	go func() {
		wg := sync.WaitGroup{}
		defer close(movieExplorer)
		n := 9
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				ExtractMovieLink(url, movieExplorer)
			}()
		}
		wg.Wait()
	}()

	go func() {
		wg := sync.WaitGroup{}
		defer close(movieToJson)
		n := 18
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func() {
				defer wg.Done()
				ExtractMovieDetail(movieToJson, movieExplorer)
			}()
		}
		wg.Wait()
	}()
	AddToList(movieToJson)
	fmt.Println(len(movies))
}

func AddToList(in <-chan string) {
	var movie utils.Movie
	for movieString := range in {
		//movieString := <-in
		if err := json.Unmarshal([]byte(movieString), &movie); err != nil {
			log.Printf("JSON unmarshaling failed: %s", err)
		}
		movies = append(movies, movie)
		log.Printf("number of crawled movies so far: %d", len(movies))
	}
}

func ExtractMovieLink(url string, out chan<- string) {
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Printf("getting %s: %s", url, resp.Status)
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Printf("parsing %s as HTML: %v", url, err)
		}

		visitNode := func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" && n.Parent.Type == html.ElementNode && n.Parent.Attr[0].Val == "titleColumn" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						link, err := resp.Request.URL.Parse(a.Val)

						if err != nil {
							continue
						}

						currentLink := removeQuery(link.String())
						if !seen[currentLink] {
							seen[currentLink] = true
							log.Printf("unique: %d", len(seen))
							out <- currentLink
						}
					}
				}
			}
		}
		forEachNode(doc, visitNode)
}

func ExtractMovieDetail(out chan<- string, in <-chan string) {
	for movieLink := range in {
		//movieLink := <-in
		resp, err := http.Get(movieLink)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Printf("getting %s: %s", movieLink, resp.Status)
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Printf("parsing %s as HTML: %v", movieLink, err)
		}
		visitNode := func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "meta" && n.Attr[0].Val == "description" {
				//s = n.NextSibling.Attr[0].Val
				out <- n.NextSibling.FirstChild.Data
				log.Printf("we're on movie %dth", len(movies))
				return
			}
		}

		forEachNode(doc, visitNode)
	}
}

func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}

func removeQuery(url string) string {
	return strings.Split(url, "?")[0]
}