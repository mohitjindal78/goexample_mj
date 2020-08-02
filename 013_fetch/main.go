package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:\t%v\n", err)
			os.Exit(1)
		}

		fmt.Println("Status :\t", resp.Status)
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s:\t%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)

		//Execrcise 1.7
		//_, err = io.Copy(os.Stdout, resp.Body)
		//resp.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: reading %s:\t%v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Println("")
	}
}
