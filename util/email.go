package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

var (
	emailrepAddr = "https://emailrep.io/"
)

//EmailReputation holds results from emailrep.io api
type EmailReputationResponse struct {
	Email      string
	Reputation string
	Suspicious bool
	References int
	Details    struct {
		Blacklisted             bool
		MaliciousActivity       bool
		MaliciousActivityRecent bool
		CredentialsLeaked       bool
		CredentialsLeakedRecent bool
		DataBreach              bool
		FirstSeen               string
		LastSeen                string
		DomainExists            bool
		DomainReputation        string
		NewDomain               bool
		DaysSinceDomainCreation int
		SuspiciousTld           bool
		Spam                    bool
		FreeProvider            bool
		Disposable              bool
		Deliverable             bool
		AcceptAll               bool
		ValidMx                 bool
		PrimaryMx               string
		Spoofable               bool
		SpfStrict               bool
		DmarcEnforced           bool
		Profiles                []interface{}
	}
}

//EmailReputation
func EmailReputation(email string) bool {
	req, reqErr := http.NewRequest(
		"GET", fmt.Sprintf("%s/%s", emailrepAddr, email), nil)

	if reqErr != nil {
		return false
	}

	res, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		return false
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		return false
	}
	var reputation EmailReputationResponse
	err := json.Unmarshal(body, &reputation)
	if err != nil {
		return false
	}
	return reputation.Details.Deliverable &&
		!reputation.Details.Blacklisted &&
		!reputation.Suspicious

}

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
	ns, err := net.LookupHost(domain)
	if err != nil {
		return false
	}
	_, err = net.LookupAddr(ns[0])
	if err != nil {
		return false
	}
	return true
}
