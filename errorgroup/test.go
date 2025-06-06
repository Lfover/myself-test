package main

import (
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	g := new(errgroup.Group)
	urls := []string{"http://www.golang.org", "http://www.google.com", "http://www.somestupidname.com"}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err == nil {
		println("Successfully fetched all URLs.")
	} else {
		println("Failed to fetch URLs:", err)
	}
}
