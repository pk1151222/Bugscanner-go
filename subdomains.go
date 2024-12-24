package scanner

import (
	"bufio"
	"os"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

// LoadDomains loads a list of domains from a file
func LoadDomains(filePath string) []string {
	var domains []string
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}
	return domains
}

// FindSubdomains enumerates subdomains for a given domain
func FindSubdomains(domain string) []string {
	options := runner.Options{Domain: domain}
	subfinder, _ := runner.NewRunner(&options)
	return subfinder.RunEnumeration()
}
