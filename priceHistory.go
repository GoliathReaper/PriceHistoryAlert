package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type PriceCheck struct {
	url       string
	maxPrice  int
	userAgent string
} 

func NewPriceCheck(url string, maxPrice int) *PriceCheck {
	return &PriceCheck{
		url:       url,
		maxPrice:  maxPrice,
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	}
}

func (pc *PriceCheck) fetchPage() (*goquery.Document, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", pc.url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", pc.userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to retrieve the webpage: status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (pc *PriceCheck) parsePrice(doc *goquery.Document) (int, error) {
	priceElement := doc.Find(".text-2xl")
	if priceElement.Length() == 0 {
		return 0, fmt.Errorf("price not found")
	}

	priceText := priceElement.Text()
	re := regexp.MustCompile(`[^\d]`)
	priceText = re.ReplaceAllString(priceText, "")

	price, err := strconv.Atoi(priceText)
	
	if err != nil {
		return 0, fmt.Errorf("error parsing the price")
	}

	return price, nil
}

func (pc *PriceCheck) checkPrice() {
	doc, err := pc.fetchPage()
	if err != nil {
		fmt.Println(err)
		return
	}

	price, err := pc.parsePrice(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Price: %d\n", price)
	if price <= pc.maxPrice {
		fmt.Println("You can buy it!")
		formattedText := fmt.Sprintf("Price: %d is below set price %d for URL: %s\n", price, pc.maxPrice, pc.url)
		telegramNotifications(formattedText)
	} else {
		fmt.Println("Price is high")
	}
}

// func main() {

// 	checker := NewPriceCheck("https://pricehistoryapp.com/product/lg-ultragear-21-9-curved-gaming-monitor-86-42-cm-34-inch-qhd-3440-x-1440-5ms-160hz-amd-freesynctm-premium-hdr-10-srgb-99-typ-height-adjust-stand-dp-hdmi-speaker-headphone-out-34gp63a", 
// 	30000)
// 	checker.checkPrice()
// }
