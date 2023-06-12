package reporter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/0xF7A4C6/GoCycle"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

var proxies *GoCycle.Cycle

func InitReporter() {
	p, err := GoCycle.NewFromFile("../../scripts/proxies.txt")
	if err != nil {
		panic(err)
	}
	p.RandomiseIndex()
	proxies = p
}

func (r *Report) Build(token string) *strings.Reader {
	p := url.Values{}

	for k, v := range map[string]string{
		`utf8`:                                 `✓`,
		`request[ticket_form_id]`:              r.FormID,
		`request[anonymous_requester_email]`:   r.Email,
		`request[custom_fields][360055270593]`: `__dc.ticket_form-tnsv1_report_other_issue__`,
		`request[custom_fields][360054260974]`: r.ReportType,
		`request[custom_fields][360008125792]`: r.MessageLink,
		`request[subject]`:                     r.Subject,
		`request[description]`:                 r.Description,
		`request[description_mimetype]`:        `text/plain`,
		`authenticity_token`:                   token,
	} {
		p.Set(k, v)
	}

	return strings.NewReader(p.Encode())
}

func (r *Report) Send() error {
	headers := http.Header{
		`authority`:                 {`support.discord.com`},
		`accept`:                    {`text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7`},
		`accept-language`:           {`fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7`},
		`cache-control`:             {`max-age=0`},
		`content-type`:              {`application/x-www-form-urlencoded`},
		`origin`:                    {`https://support.discord.com`},
		`referer`:                   {`https://support.discord.com/hc/en-us/requests/new?ticket_form_id=360000029731`},
		`sec-ch-ua`:                 {`"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		`sec-ch-ua-mobile`:          {`?0`},
		`sec-ch-ua-platform`:        {`"Windows"`},
		`sec-fetch-dest`:            {`document`},
		`sec-fetch-mode`:            {`navigate`},
		`sec-fetch-site`:            {`same-origin`},
		`sec-fetch-user`:            {`?1`},
		`upgrade-insecure-requests`: {`1`},
		`user-agent`:                {`Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36`},

		http.HeaderOrderKey: {
			`authority`,
			`accept`,
			`accept-language`,
			`cache-control`,
			`content-type`,
			`origin`,
			`referer`,
			`sec-ch-ua`,
			`sec-ch-ua-mobile`,
			`sec-ch-ua-platform`,
			`sec-fetch-dest`,
			`sec-fetch-mode`,
			`sec-fetch-site`,
			`sec-fetch-user`,
			`upgrade-insecure-requests`,
			`user-agent`,
		},
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(30),
		tls_client.WithClientProfile(tls_client.Chrome_112),
		tls_client.WithProxyUrl(fmt.Sprintf("http://%s", proxies.Next())),
	}...)
	if err != nil {
		return err
	}

	// Get cookie & session token.
	req, err := http.NewRequest(http.MethodGet, "https://support.discord.com/hc/api/internal/csrf_token.json", nil)
	if err != nil {
		return err
	}

	req.Header = headers

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cant get cookies status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var s Session
	if err := json.Unmarshal(body, &s); err != nil {
		return err
	}

	/*
		fmt.Println(s)
		fmt.Println(strings.Split(strings.Split(resp.Header.Get("set-cookie"), "_help_center_session=")[1], "; path=/")[0])
		fmt.Println(r.Build(s.CurrentSession.CsrfToken))
	*/

	// Send report.
	req, err = http.NewRequest(http.MethodPost, "https://support.discord.com/hc/en-us/requests", r.Build(s.CurrentSession.CsrfToken))
	if err != nil {
		return err
	}

	headers.Add("cookie", fmt.Sprintf("_help_center_session=%s", strings.Split(strings.Split(resp.Header.Get("set-cookie"), "_help_center_session=")[1], "; path=/")[0]))
	req.Header = headers

	resp, err = client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("cant send report status code %d", resp.StatusCode)
	}

	return nil
}
