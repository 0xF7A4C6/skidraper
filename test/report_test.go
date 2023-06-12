package test

import (
	"github.com/0xF7A4C6/skidraper/internal/reporter"
	"testing"
)

func TestReport_Send(t *testing.T) {
	tests := []struct {
		name    string
		r       *reporter.Report
		wantErr bool
	}{
		{
			name: "report test",
			r: &reporter.Report{
				MessageLink: "https://discord.com/channels/997150957001510923/997150957542588499/1117873779658465291",
				Description: "This user is underage",
				ReportType:  "__dc.ticket_form-tnsv1_cat_-_underage_user__",
				Subject:     "Underage user !!",
				Email:       "ifuckdiscord@internetkeno.com",
				FormID:      "360000029731", // Underage
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Send(); (err != nil) != tt.wantErr {
				t.Errorf("Report.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
