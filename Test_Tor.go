package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Specify Tor proxy ip and port
var torProxy string = "socks5://127.0.0.1:9050" // 9150 w/ Tor Browser

func main() {
//Ask user for URL
	fmt.Println("Enter URL")
	reader := bufio.NewReader(os.Stdin)
	webUrl, _ := reader.ReadString('\n')
	webUrl = strings.Replace(webUrl, "\n", "", -1)

	// Parse Tor proxy URL string to a URL type
	torProxyUrl, err := url.Parse(torProxy)
	if err != nil {
		log.Fatal("Error parsing Tor proxy URL:", torProxy, ".", err)
	}

	// Set up a custom HTTP transport to use the proxy and create the client
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyUrl)}
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 5}

	// Make request
	resp, err := client.Get(webUrl)
	if err != nil {
		log.Fatal("Error making GET request.", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body of response.", err)
	}
	log.Println(string(body))
}
