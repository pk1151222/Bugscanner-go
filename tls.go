package scanner

import (
	"crypto/tls"
	"fmt"
	"net"
)

type ScanResult struct {
	Domain       string
	IP           string
	TLSVersions  []string
	CipherSuites []string
	Server       string
}

// ScanDomain performs SNI and TLS scanning on a domain
func ScanDomain(domain string) ScanResult {
	result := ScanResult{Domain: domain}

	// Resolve IP
	ips, err := net.LookupIP(domain)
	if err == nil && len(ips) > 0 {
		result.IP = ips[0].String()
	}

	// Perform TLS handshake
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{ServerName: domain})
	if err != nil {
		fmt.Printf("TLS error for %s: %v\n", domain, err)
		return result
	}
	defer conn.Close()

	// Collect TLS details
	state := conn.ConnectionState()
	result.TLSVersions = append(result.TLSVersions, state.NegotiatedProtocol)
	result.CipherSuites = append(result.CipherSuites, tls.CipherSuiteName(state.CipherSuite))
	result.Server = state.ServerName

	return result
}
