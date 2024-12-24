package scanner

import (
	"github.com/phpdave11/gofpdf"
)

func GeneratePDF(results []ScanResult, fileName string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Bug Scanner Report")

	for _, result := range results {
		pdf.Ln(10)
		pdf.Cell(40, 10, "Domain: "+result.Domain)
		pdf.Cell(40, 10, "IP: "+result.IP)
		pdf.Cell(40, 10, "Server: "+result.Server)
	}

	return pdf.OutputFileAndClose(fileName)
}
