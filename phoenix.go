package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/net/proxy"

	"github.com/y-a-t-s/firebird"
)

const HOST = "kiwifarms.net"
const ONION_HOST = "kiwifarmsaaf4t2h7gc3dfc5ojhmqruw2nit3uejrpiagrxeuxiyxcyd.onion"

func main() {
	p := proxy.FromEnvironment()

	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.DialContext = p.(proxy.ContextDialer).DialContext

	hc := http.Client{}
	hc.Transport = tr

	c, err := firebird.NewChallenge(hc, ONION_HOST)
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
