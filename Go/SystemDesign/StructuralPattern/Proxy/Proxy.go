package main

import "fmt"

type Internet interface{
	Connect(url string)
}

type RealInternet struct{
	
}
func (r *RealInternet) Connect(url string) {
	fmt.Println("Connecting to:", url)
}

var blocked = map[string]bool{
	"facebook.com": true,
	"twitter.com":  true,
}

type ProxyInternet struct{
	internet *RealInternet
}

func (p *ProxyInternet) Connect(url string) {

	if blocked[url] {
		fmt.Println("Access Denied:", url)
		return
	}

	p.internet.Connect(url)
}

func main() {

	proxy := &ProxyInternet{
		internet: &RealInternet{},
	}

	proxy.Connect("google.com")
	proxy.Connect("facebook.com")
}