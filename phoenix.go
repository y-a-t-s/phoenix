package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/proxy"

	"github.com/y-a-t-s/firebird"
)

const ONION_HOST = "kiwifarmsaaf4t2h7gc3dfc5ojhmqruw2nit3uejrpiagrxeuxiyxcyd.onion"

func hostURL() string {
	host := os.Getenv("KF_HOST")
	if host == "" {
		return ONION_HOST
	}

	return host
}

func main() {
	p := proxy.FromEnvironment()

	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DialContext = p.(proxy.ContextDialer).DialContext

	hc := http.Client{}
	hc.Transport = tr

	c, err := firebird.NewChallenge(hc, hostURL())
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s, err := firebird.Solve(ctx, c)
	if err != nil {
		panic(err)
	}

	a, err := firebird.Submit(hc, s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %s\n", a)
}
