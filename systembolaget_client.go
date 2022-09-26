package systembolaget

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// https://www.systembolaget.se
// /api/gateway/productsearch/search/
// baseURL https://api-extern.systembolaget.se/sb-api-ecommerce/v1

//https://www.systembolaget.se/api/gateway/productsearch/search/?page=1&size=30&sortBy=Score&sortDirection=Ascending&categoryLevel1=Öl

//headers:
//BaseURL: https://api-extern.systembolaget.se/sb-api-ecommerce/v1
// referer: https://www.systembolaget.se/sortiment/ol/
//Content-Type: application/x-www-form-urlencoded
//

func (c *Client) GetProductPages(category1 Category1, category2 Category2) []ProductPage {

	productPages := []ProductPage{}
	productStartPage := c.getProductPage(string(URL), 1, category1, category2)
	productPages = append(productPages, productStartPage)
	hasNextPage := c.nextPageExists(productStartPage)
	nextPageNumber := productStartPage.Metadata.NextPage

	for {
		if hasNextPage {
			productPage := c.getProductPage(string(URL), nextPageNumber, category1, category2)
			productPages = append(productPages, productPage)
			hasNextPage = c.nextPageExists(productPage)
			if hasNextPage {
				nextPageNumber = productPage.Metadata.NextPage
			}

		} else {
			break
		}
	}
	return productPages
}

func (cc *Client) nextPageExists(productPage ProductPage) bool {
	return productPage.Metadata.NextPage != -1
}

// func (cc *Client) getNextPage(productPage ProductPage) int {
// 	return productPage.Metadata.NextPage
// }

func (cc *Client) getProductPage(URL string, pageNumber int, category1 Category1, category2 Category2) ProductPage {

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	req.Header.Add("BaseURL", "https://api-extern.systembolaget.se/sb-api-ecommerce/v1")
	req.Header.Add("Accept", "*/*")

	//Query string parameters
	q := req.URL.Query()
	q.Add("page", strconv.Itoa(pageNumber))
	q.Add("size", "30")
	q.Add("sortBy", "Score")
	q.Add("sortDirection", "Ascending")
	if category1 != "" {
		q.Add("categoryLevel1", string(category1))

		if category1 == "Vin" {
			q.Add("categoryLevel2", string(category2))
		}
	}
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var productPage *ProductPage
	err = json.Unmarshal(resBody, &productPage)
	if err != nil {
		fmt.Printf("client: could not unmarshal response body: %s\n", err)
	}

	return *productPage
}
