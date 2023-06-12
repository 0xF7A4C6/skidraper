package main

import (
	"fmt"
	"strconv"

	"github.com/0xF7A4C6/skidraper/internal/console"
	"github.com/0xF7A4C6/skidraper/internal/discord"
	"github.com/0xF7A4C6/skidraper/internal/reporter"
	"github.com/zenthangplus/goccm"
)

func main() {
	console.PrintLogo()
	reporter.InitReporter()
	discord.InitScraper("MTExNzQxNTcwNDMwODc0ODM4MQ.GJPUYq.LSLzomFRlQFDRaJf0QjkIPU9B7K4XuIZt7fIzY")

	guildID := console.Input("guildID")
	channelID := console.Input("channelID")
	
	maxReq, err := strconv.Atoi(console.Input("MaxReq"))
	if err != nil {
		panic(err)
	}

	threadCount, err := strconv.Atoi(console.Input("Threads"))
	if err != nil {
		panic(err)
	}

	fmt.Println("")
	messages, err := discord.ScrapeChannel(channelID, maxReq)

	if len(messages) <= 0 && err != nil {
		panic(err)
	}

	console.Confirm("\npress to send...")

	c := goccm.New(threadCount)

    for i, message := range messages {
		c.Wait()

        go func(i int, messID string) {
            defer c.Done()

			r := &reporter.Report{
				MessageLink: fmt.Sprintf("https://discord.com/channels/%s/%s/%s", guildID, channelID, messID),
				Description: "This user is underage",
				ReportType:  "__dc.ticket_form-tnsv1_cat_-_underage_user__",
				Subject:     "Underage user !!",
				Email:       "khggvupyqwfh@internetkeno.com",
				FormID:      "360000029731", // Underage
			}

			if err := r.Send(); err != nil {
				console.Log(fmt.Sprintf("<fg=594cad>[</><fg=67f591>%d</><fg=594cad>]</> <fg=eb4034>error #</><fg=eb4034>%s</>\n", i, messID))
				return
			}

			console.Log(fmt.Sprintf("<fg=594cad>[</><fg=67f591>%d</><fg=594cad>]</> <fg=e9e4f2>sent #</><fg=594cad>%s</>\n", i, messID))
        }(i, message)
    }
	
	c.WaitAllDone()
	console.Confirm("Done !")
}