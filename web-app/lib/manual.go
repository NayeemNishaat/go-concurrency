package lib

import (
	"fmt"
	"web/model"

	"github.com/go-pdf/fpdf"
)

func (cfg *Config) GenerateManual(u model.User, p *model.Plan) *fpdf.Fpdf {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetMargins(10, 13, 10)
	pdf.SetFont("Arial", "B", 16)
	pdf.SetX(75)
	pdf.SetY(150)

	pdf.MultiCell(0, 4, fmt.Sprintf("%s %s", u.FirstName, u.LastName), "", "C", false)
	pdf.Ln(5)
	pdf.MultiCell(0, 4, fmt.Sprintf("%s User Guide", u.FirstName), "", "C", false)

	return pdf
}
