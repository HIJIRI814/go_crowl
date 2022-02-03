package main

import (
	"fmt"
	"time"

	"github.com/sclevine/agouti"
)

const CROWLING_URL = "https://www.youtube.com/"

type Crowler struct {
	driver *agouti.WebDriver
}

func (c *Crowler) Start() (err error) {
	c.driver.Start()
	return
}

func (c *Crowler) Stop() {
	c.driver.Stop()
}

func (c *Crowler) crawling(url string) (page *agouti.Page, err error) {
	page, err = c.driver.NewPage()
	page.Navigate(url)
	time.Sleep(1 * time.Second)
	return
}

func main() {
	c := &Crowler{driver: agouti.ChromeDriver()}
	defer c.Stop()
	if err := c.Start(); err != nil {
		fmt.Println(err)
		return
	}

	page, err := c.crawling(CROWLING_URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(page.Title())
}
