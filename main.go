package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/namedotcom/go/namecom"
)

func tryToBuy(nc *namecom.NameCom, domain string) bool {
	availability, err := nc.CheckAvailability(&namecom.AvailabilityRequest{
		DomainNames: []string{domain},
		PromoCode:   "",
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "failure: %v\n", err)
		return false
	}

	if availability.Results[0].Purchasable {
		fmt.Printf("can buy for $%f\n", availability.Results[0].PurchasePrice)
	} else {
		fmt.Printf("cannot buy\n")
	}

	return true
}

func main() {
	user := os.Args[1]
	token := os.Args[2]
	domain := os.Args[3]

	fmt.Printf("Polling for %s...\n", domain)

	var nc *namecom.NameCom
	if strings.HasSuffix(user, "-test") {
		nc = namecom.Test(user, token)
	} else {
		nc = namecom.New(user, token)
	}

	helloResp, err := nc.HelloFunc(&namecom.HelloRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println(helloResp.Motd)

	for {
		tryToBuy(nc, domain)
		time.Sleep(30 * time.Second)
	}
}
