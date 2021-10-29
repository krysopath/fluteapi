package util

import (
	"log"
	"net"
	"strings"
)

// IsResolvableEmail returns true if the email can resolve is supported
func IsResolvableEmail(email string) bool {
	var domain string
	at := strings.LastIndex(email, "@")
	if at >= 0 {
		_, domain = email[:at], email[at+1:]
	} else {
		log.Printf("Error: %s is an invalid email address\n", email)
		return false
	}

	// Resolve cname
	//cname, _ := net.LookupCNAME(domain)
	//log.Println("domain has cname:", cname)

	// Resolve ip address
	ns, err := net.LookupHost(domain)
	if err != nil {
		//log.Printf("Err: %s", err.Error())
		return false
	}

	// Reverse Resolution (Host must be able to resolve to address)
	_, err = net.LookupAddr(ns[0])
	if err != nil {
		//log.Printf("Err: %s", err.Error())
		return false
	}
	//log.Println("domain can reverse-resolved", dnsname)

	return true
}
