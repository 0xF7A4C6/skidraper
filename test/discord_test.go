package test

import (
	"fmt"
	"testing"

	"github.com/0xF7A4C6/skidraper/internal/discord"
)

func TestScrapeChannel(t *testing.T) {
	discord.InitScraper("MTExNzQxNTcwNDMwODc0ODM4MQ.GJPUYq.LSLzomFRlQFDRaJf0QjkIPU9B7K4XuIZt7fIzY")

	type args struct {
		channelID string
		maxRound  int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "scrape-test",
			args: args{
				channelID: "997150957542588499",
				maxRound:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := discord.ScrapeChannel(tt.args.channelID, tt.args.maxRound)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScrapeChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(got)
		})
	}
}
