package discord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/0xF7A4C6/GoCycle"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/gookit/color"
)

var (
	client  tls_client.HttpClient
	header  http.Header
	proxies *GoCycle.Cycle
)

func InitScraper(token string) {
	header = http.Header{
		`authority`:          {`discord.com`},
		`accept`:             {`*/*`},
		`accept-language`:    {`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		`authorization`:      {token},
		`referer`:            {`https://discord.com/channels/997150957001510923/997150957542588499`},
		`sec-ch-ua`:          {`"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		`sec-ch-ua-mobile`:   {`?0`},
		`sec-ch-ua-platform`: {`"Windows"`},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-origin`},
		`user-agent`:         {`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`},
		`x-debug-options`:    {`bugReporterEnabled`},
		`x-discord-locale`:   {`fr`},
		`x-discord-timezone`: {`Europe/Paris`},
		`x-super-properties`: {`eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImZyLUZSIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzExNC4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTE0LjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjIwNDIwNCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=`},

		http.HeaderOrderKey: {
			`authority`,
			`accept`,
			`accept-language`,
			`authorization`,
			`cookie`,
			`referer`,
			`sec-ch-ua`,
			`sec-ch-ua-mobile`,
			`sec-ch-ua-platform`,
			`sec-fetch-dest`,
			`sec-fetch-mode`,
			`sec-fetch-site`,
			`user-agent`,
			`x-debug-options`,
			`x-discord-locale`,
			`x-discord-timezone`,
			`x-super-properties`,
		},
	}

	p, err := GoCycle.NewFromFile("../../scripts/proxies.txt")
	if err != nil {
		panic(err)
	}
	p.RandomiseIndex()
	proxies = p
}

func ScrapeChannel(channelID string, maxReq int) ([]string, error) {
	client, _ = tls_client.NewHttpClient(tls_client.NewNoopLogger(), []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(tls_client.Chrome_112),
		tls_client.WithProxyUrl(fmt.Sprintf("http://%s", proxies.Next())),
	}...)

	var beforeID string
	ids := []string{}

	scrape := func(url string) ([]string, error) {
		found := []string{}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return ids, err
		}

		req.Header = header

		resp, err := client.Do(req)
		if err != nil {
			return ids, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return ids, fmt.Errorf("cant get channel messages status code %d", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ids, err
		}

		var m Messages
		if err := json.Unmarshal(body, &m); err != nil {
			return ids, err
		}

		for _, mess := range m {
			if mess.Author.ID == "1116776015558086797" {
				continue
			}
			found = append(found, mess.ID)
		}

		return found, nil
	}

	firstScrape, err := scrape(fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages?limit=50", channelID))
	if err != nil {
		return ids, err
	}

	ids = append(ids, firstScrape...)
	beforeID = firstScrape[len(firstScrape)-1]

	for i := 0; i < maxReq; i++ {
		result, err := scrape(fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages?before=%s&limit=50", channelID, beforeID))
		if err != nil {
			return ids, err
		}

		color.Printf("<fg=594cad>[</><fg=67f591>+</><fg=594cad>]</> <fg=594cad>%d</> <fg=e9e4f2>ids (ttl:</> <fg=594cad>%d</><fg=e9e4f2>, round:</> <fg=594cad>%d</><fg=e9e4f2>/</><fg=594cad>%d</>)\n", len(result), len(ids), i, maxReq)

		ids = append(ids, result...)
		beforeID = result[len(result)-1]
	}

	return ids, nil
}
