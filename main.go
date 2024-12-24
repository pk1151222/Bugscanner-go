package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"bug-scanner/scanner"
)

func main() {
	// Define command-line flags
	domain := flag.String("domain", "", "Single domain to scan")
	domainList := flag.String("domains", "", "File with list of domains")
	output := flag.String("output", "output.json", "File to save the results")
	pdf := flag.String("pdf", "report.pdf", "PDF summary report")
	flag.Parse()

	// Input validation
	if *domain == "" && *domainList == "" {
		log.Fatal("Please provide a domain or domain list using -domain or -domains")
	}

	// Load domains
	var domains []string
	if *domainList != "" {
		domains = scanner.LoadDomains(*domainList)
	} else {
		domains = []string{*domain}
	}

	// Start scanning
	var wg sync.WaitGroup
	results := make(chan scanner.ScanResult, len(domains))

	for _, d := range domains {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			result := scanner.ScanDomain(domain)
			results <- result
		}(d)
	}

	wg.Wait()
	close(results)

	// Save results
	var allResults []scanner.ScanResult
	for r := range results {
		allResults = append(allResults, r)
	}

	err := scanner.SaveResults(allResults, *output)
	if err != nil {
		log.Fatalf("Error saving results: %v", err)
	}

	// Generate PDF summary
	err = scanner.GeneratePDF(allResults, *pdf)
	if err != nil {
		log.Fatalf("Error generating PDF: %v", err)
	}

	fmt.Printf("Scanning completed. Results saved in %s and PDF report generated in %s\n", *output, *pdf)
}
