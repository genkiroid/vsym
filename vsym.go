package vsym

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/genkiroid/cert"
)

var Issuers = []string{
	"Symantec",
	"GeoTrust",
	"thawte",
	"RapidSSL",
}

type VSym struct {
	certs cert.Certs
}

func NewVSym() (*VSym, error) {
	certs, err := cert.NewCerts(os.Args[1:])
	if err != nil {
		return nil, err
	}
	return &VSym{certs}, nil
}

func symantecIssued(cert *cert.Cert) bool {
	for _, issuer := range Issuers {
		if strings.Contains(strings.ToLower(cert.Issuer), strings.ToLower(issuer)) {
			return true
		}
	}
	return false
}

func chromeVer(cert *cert.Cert) string {
	issuedDate := cert.Certificate.NotBefore.In(time.Local)
	baseDate := time.Date(2016, time.June, 1, 0, 0, 0, 0, time.Local)

	if issuedDate.Before(baseDate) {
		return "v66"
	}
	return "v70"
}

func (v *VSym) Verify() {
	for _, cert := range v.certs {
		if symantecIssued(cert) {
			fmt.Printf(
				"The SSL certificate used on https://%s will be distrusted in Chrome %s.\n",
				cert.DomainName,
				chromeVer(cert),
			)
		}
	}
}
