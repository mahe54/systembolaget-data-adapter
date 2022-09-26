package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	p "github.com/mahe54/systembolaget-data-adapter/internal/product"
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

type Information struct {
	link  string
	title string
}

func main() {

	webUrl := "https://www.systembolaget.se/api/gateway/productsearch/search/"
	categories := []string{"Öl", "Sprit", "Cider & blanddrycker", "Alkoholfritt", "Presentartiklar", "Vin"}
	vinSubCategories := []string{"Rött vin", "Vitt vin", "Mousserande vin", "Rosévin", "Vinlåda", "Starkvin"}
	var productPages []p.ProductPage

	for _, category := range categories {
		if category == "Vin" {
			for _, subCategory := range vinSubCategories {
				productPages = append(productPages, getAllProductPages(webUrl, 1, category, subCategory)...)
			}
		} else {
			productPages = append(productPages, getAllProductPages(webUrl, 1, category, "")...)
		}
		saveProductPagesToFile(productPages, category)
	}

	// for _, productPage := range productPages {
	// 	for _, product := range productPage.Products {
	// 		fmt.Println(product.ProductNameBold)
	// 	}
	// }

}

func saveProductPagesToFile(productPages []p.ProductPage, category string) {
	productPagesJson, _ := json.Marshal(productPages)
	filename := "productPages_" + category + ".json"
	err := ioutil.WriteFile(filename, productPagesJson, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func getAllProductPages(URL string, startPage int, category, subCategory string) []p.ProductPage {

	productPages := []p.ProductPage{}
	productStartPage := getProductPage(URL, startPage, category, subCategory)
	productPages = append(productPages, productStartPage)
	hasNextPage := nextPageExists(productStartPage)
	nextPageNumber := productStartPage.Metadata.NextPage

	for {
		if hasNextPage {
			productPage := getProductPage(URL, nextPageNumber, category, subCategory)
			productPages = append(productPages, productPage)
			hasNextPage = nextPageExists(productPage)
			if hasNextPage {
				nextPageNumber = productPage.Metadata.NextPage
			}

		} else {
			break
		}
	}
	return productPages
}

func nextPageExists(productPage p.ProductPage) bool {
	return productPage.Metadata.NextPage != -1
}

func getNextPage(productPage p.ProductPage) int {
	return productPage.Metadata.NextPage
}

func getProductPage(URL string, pageNumber int, category, subCategory string) p.ProductPage {

	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	//req.uri
	req.Header.Add("BaseURL", "https://api-extern.systembolaget.se/sb-api-ecommerce/v1")
	req.Header.Add("Accept", "*/*")

	//Query string parameters
	q := req.URL.Query()
	q.Add("page", strconv.Itoa(pageNumber))
	q.Add("size", "30")
	q.Add("sortBy", "Score")
	q.Add("sortDirection", "Ascending")
	if category != "" {
		q.Add("categoryLevel1", category)

		if category == "Vin" {
			q.Add("categoryLevel2", subCategory)
		}
	}
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	var productPage *p.ProductPage
	_ = json.Unmarshal(resBody, &productPage)

	totalpages := productPage.Metadata.DocCount/30 + 1

	//fmt.Print("Next page: ", productPage.Metadata.NextPage, " ")
	fmt.Printf("\rStatus code: %d | category: %s | page: %d/%d", res.StatusCode, category, pageNumber, totalpages)
	return *productPage
}
