package main

import (
	"fmt"
	"net"
)

func main() {
	var domain string = "bigm-bd.com"

	mxRecords, _ := net.LookupMX(domain)
	if len(mxRecords) <= 0 {
		fmt.Println("MX Record not found")
		return
	}

	for _, mx := range mxRecords {
		fmt.Println(mx.Host)
		fmt.Println(mx.Pref)
	}
}
