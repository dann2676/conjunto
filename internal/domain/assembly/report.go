package assembly

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
	"golang.org/x/text/encoding/charmap"
)

func (s *service) GenerateReport(ctx context.Context, assemblyID int) ([]byte, error) {
	assembly, err := s.repo.Get(ctx, assemblyID)
	if err != nil {
		return nil, err
	}

	items, err := s.repo.GetAgendaItems(ctx, assemblyID)
	if err != nil {
		return nil, err
	}

	attendance, err := s.repo.GetAttendance(ctx, assemblyID)
	if err != nil {
		return nil, err
	}

	quorum, err := s.GetQuorum(ctx, assemblyID)
	if err != nil {
		return nil, err
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(20, 20, 20)

	// Encabezado
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "ACTA DE ASAMBLEA", "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "B", 13)
	pdf.CellFormat(0, 8, assembly.Title, "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 11)
	tipoStr := "Ordinaria"
	if assembly.Type == "extraordinaria" {
		tipoStr = "Extraordinaria"
	}
	pdf.CellFormat(0, 7, fmt.Sprintf("Fecha: %s  |  Tipo: %s",
		assembly.Date.Format("02/01/2006"), tipoStr), "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 7, "Modalidad: Virtual/Presencial  |  Edificio Marquez del Prado", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Verificación de identidad
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(0, 8, toIso("1. MÉTODO DE VERIFICACIÓN DE IDENTIDAD"), "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 6,
		"La identidad de cada asistente fue verificada mediante número de cédula de ciudadanía "+
			"cruzado contra el registro de propietarios del edificio. Los asistentes no registrados "+
			"en el sistema ingresaron su número de cédula y nombre completo, quedando registrados "+
			"como apoderados externos con la hora exacta de ingreso.", "", "L", false)
	pdf.Ln(4)

	// Quórum
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(0, 8, "2. VERIFICACIÓN DE QUÓRUM", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 6, fmt.Sprintf("Quórum requerido: %.1f%%", assembly.QuorumRequired), "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 6, fmt.Sprintf("Quórum alcanzado: %.1f%%", quorum), "", 1, "L", false, 0, "")
	alcanzado := "NO"
	if quorum >= assembly.QuorumRequired {
		alcanzado = "SÍ"
	}
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(0, 6, fmt.Sprintf("Quórum suficiente para sesionar: %s", alcanzado), "", 1, "L", false, 0, "")
	pdf.Ln(4)

	// Lista de asistentes
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(0, 8, "3. LISTA DE ASISTENTES", "", 1, "L", false, 0, "")

	// Encabezado tabla
	pdf.SetFont("Arial", "B", 9)
	pdf.SetFillColor(50, 50, 50)
	pdf.SetTextColor(255, 255, 255)
	pdf.CellFormat(10, 7, "#", "1", 0, "C", true, 0, "")
	pdf.CellFormat(35, 7, "Cédula", "1", 0, "C", true, 0, "")
	pdf.CellFormat(55, 7, "Nombre", "1", 0, "C", true, 0, "")
	pdf.CellFormat(20, 7, "Unidad", "1", 0, "C", true, 0, "")
	pdf.CellFormat(25, 7, "Coeficiente", "1", 0, "C", true, 0, "")
	pdf.CellFormat(25, 7, "Calidad", "1", 0, "C", true, 0, "")
	pdf.CellFormat(0, 7, "Hora ingreso", "1", 1, "C", true, 0, "")

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 9)
	fill := false
	for i, a := range attendance {
		pdf.SetFillColor(240, 240, 240)
		calidad := "Propietario"
		if a.IsProxy {
			calidad = "Apoderado"
		}
		pdf.CellFormat(10, 6, fmt.Sprintf("%d", i+1), "1", 0, "C", fill, 0, "")
		pdf.CellFormat(35, 6, a.AttendedByID, "1", 0, "L", fill, 0, "")
		pdf.CellFormat(55, 6, a.AttendedBy, "1", 0, "L", fill, 0, "")
		pdf.CellFormat(20, 6, fmt.Sprintf("%d", a.UnitNumber), "1", 0, "C", fill, 0, "")
		pdf.CellFormat(25, 6, fmt.Sprintf("%.4f", a.Coeficient), "1", 0, "C", fill, 0, "")
		pdf.CellFormat(25, 6, calidad, "1", 0, "C", fill, 0, "")
		pdf.CellFormat(0, 6, "Ver registro digital", "1", 1, "L", fill, 0, "")
		fill = !fill
	}
	pdf.Ln(5)

	// Orden del día y votaciones
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(0, 8, "4. ORDEN DEL DÍA Y VOTACIONES", "", 1, "L", false, 0, "")

	for _, item := range items {
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(0, 7, fmt.Sprintf("%d. %s", item.Order, item.Title), "", 1, "L", false, 0, "")
		if item.Description != "" {
			pdf.SetFont("Arial", "I", 9)
			pdf.MultiCell(0, 5, item.Description, "", "L", false)
		}

		if item.Status == "closed" {
			votes, err := s.repo.GetVotes(ctx, item.ID)
			if err == nil && len(votes) > 0 {
				results := map[string]float32{"yes": 0, "no": 0, "abstain": 0}
				counts := map[string]int{"yes": 0, "no": 0, "abstain": 0}
				for _, v := range votes {
					results[v.Value] += v.Coeficient
					counts[v.Value]++
				}
				pdf.SetFont("Arial", "", 9)
				pdf.CellFormat(0, 5, fmt.Sprintf(
					"  A favor: %.1f%% (%d votos)  |  En contra: %.1f%% (%d votos)  |  Abstención: %.1f%% (%d votos)",
					results["yes"]*100, counts["yes"],
					results["no"]*100, counts["no"],
					results["abstain"]*100, counts["abstain"],
				), "", 1, "L", false, 0, "")
			}
		} else {
			pdf.SetFont("Arial", "I", 9)
			pdf.CellFormat(0, 5, "  (Sin votación registrada)", "", 1, "L", false, 0, "")
		}
		pdf.Ln(2)
	}

	// Firmas
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(0, 8, "5. FIRMAS", "", 1, "L", false, 0, "")
	pdf.Ln(15)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(80, 0, "_______________________", "", 0, "C", false, 0, "")
	pdf.CellFormat(0, 0, "_______________________", "", 1, "C", false, 0, "")
	pdf.CellFormat(80, 6, "Presidente de la Asamblea", "", 0, "C", false, 0, "")
	pdf.CellFormat(0, 6, "Secretario de la Asamblea", "", 1, "C", false, 0, "")

	// Pie de página
	pdf.SetFont("Arial", "I", 8)
	pdf.SetY(-20)
	pdf.CellFormat(0, 5, fmt.Sprintf(
		"Reporte generado digitalmente el %s — Edificio Marquez del Prado",
		time.Now().Format("02/01/2006 15:04"),
	), "", 1, "C", false, 0, "")

	var buf bytes.Buffer
	if err = pdf.Output(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func toIso(s string) string {
	encoder := charmap.ISO8859_1.NewEncoder()
	result, _ := encoder.String(s)
	return result
}
