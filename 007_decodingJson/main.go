//This sample demostrate how to decode JSON response
//using the json package and NewDecoder function.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	// gResult maps to the result document received from the search.
	gResult struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}

	//gResponse contains the top level document.
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

func init() {}

func main() {
	//uri := "http://ajax.googleapis.com/ajax/services/search/web?v=1.0&rsz=8&q=golang"
	uri := "https://www.google.com/search?q=goolang&rlz=1C1CHBF_enIN796IN796&oq=goolang&aqs=chrome..69i57.5883j0j8&sourceid=chrome&ie=UTF-8"

	//Issue the search against the google
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	//Decode the JSON response into our struct type
	var gr gResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	fmt.Println(gr)
}
