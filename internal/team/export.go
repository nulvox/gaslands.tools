package team

import (
	"bytes"
	_ "embed"
	"html/template"
	"strings"

	"gaslands.tools/internal/gamedata"
)

//go:embed export_template.html
var exportTemplate string

// exportData is the template data for HTML export.
type exportData struct {
	Team         *Team
	Spent        int
	Remaining    int
	TotalHull    int
	VehicleCount int
	Vehicles     []exportVehicle
	Sponsor      *gamedata.Sponsor
}

// exportVehicle is the per-vehicle template data.
type exportVehicle struct {
	Vehicle    *Vehicle
	VType      *gamedata.VehicleType
	Cost       int
	SlotsUsed  int
	SlotsTotal int
	SlotPct    int
}

// ExportHTML generates a standalone styled HTML file for the team roster.
func ExportHTML(t *Team) (string, error) {
	tmpl, err := template.New("export").Funcs(template.FuncMap{
		"upper": strings.ToUpper,
		"sub":   func(a, b int) int { return a - b },
	}).Parse(exportTemplate)
	if err != nil {
		return "", err
	}

	spent := TeamCost(t)
	data := exportData{
		Team:         t,
		Spent:        spent,
		Remaining:    t.Budget - spent,
		TotalHull:    TeamHull(t),
		VehicleCount: len(t.Vehicles),
	}

	if t.Sponsor.Name != "" && !t.Sponsor.IsCustom {
		s, ok := gamedata.GetSponsor(t.Sponsor.Name)
		if ok {
			data.Sponsor = &s
		}
	}

	for i := range t.Vehicles {
		v := &t.Vehicles[i]
		ev := exportVehicle{
			Vehicle:   v,
			Cost:      VehicleCost(v),
			SlotsUsed: SlotsUsed(v),
		}
		vt, ok := gamedata.GetVehicleType(v.VehicleType)
		if ok {
			ev.VType = &vt
			ev.SlotsTotal = vt.BuildSlots
			if vt.BuildSlots > 0 {
				ev.SlotPct = ev.SlotsUsed * 100 / vt.BuildSlots
			}
		}
		data.Vehicles = append(data.Vehicles, ev)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
