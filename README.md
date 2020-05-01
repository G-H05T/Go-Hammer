# Hammer
This script is meant to effectively take down Apache websites by sending a ton of POST requests through Tor.
I am NOT resposible for anything this is used for. This is meant for testing and research purposes. I've used it to take down a few phishing sites and I know the potential this has.
If you decide to use this script, you are responsible for anything that happens as a result.
I'm unsure if this works in Windows, but I have tested it on Debian 10. It should work on all Linux distros as long as you have Go installed and should also work on Mac OS X.
If you think this could be improved, let me know or make a fork. I'm still learning Go.
## Getting Started
First, you need to install Go and the dependencies.
I won't walk you through installing Go, but this is how to install the dependencies:

```
go get log net/http net/url time bufio fmt os strings runtime github.com/inancgumus/screen math/rand
```

Make sure you have Tor installed and running before you run this script. NOT the Tor Browser Bundle, but if that's all you have, change the port 9050 in main.go to 9150 on line 18.
### Installing
```
git clone https://github.com/G-H05T/Go-Hammer.git
```

```
go run main.go AA.go UA.go
```

OR

```
go build main.go AA.go UA.go
```

```
mv main hammer
```

```
./hammer
```
### Screencap
![Hammer](https://github.com/G-H05T/Go-Hammer/blob/master/Hammer.png)
