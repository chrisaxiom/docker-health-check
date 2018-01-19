package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Version = "0.0.3"
	app.Name = "Health Checker"
	app.Usage = "Hits an endpoint for you.  healthcheck -url=http://localhost/ping"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, U",
			Usage:       "the url to hit (required)",
			Destination: &url,
		},
		cli.StringSliceFlag{
			Name:  "headers, H",
			Usage: "specify a header and value for the request (-H=key:value)",
		},
		cli.StringFlag{
			Name:        "verb, V",
			Usage:       "the HTTP verb to use",
			Value:       "GET",
			Destination: &httpVerb,
		},
		cli.IntFlag{
			Name:        "code, C",
			Usage:       "expected response code",
			Value:       http.StatusOK,
			Destination: &statusCode,
		},
		// http body not supported yet
		// response body checking not supported yet
	}
	app.Action = actionFunc

	app.Run(os.Args)
}

func actionFunc(c *cli.Context) error {
	if len(url) < 0 {
		return cli.NewExitError("url length must be > 0 ", 1)
	}

	req, err := http.NewRequest(httpVerb, url, nil)

	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	for _, str := range c.StringSlice("headers") {
		kv := strings.Split(str, ":")
		if len(kv) == 2 {
			req.Header.Add(kv[0], kv[1])
		} else {
			return cli.NewExitError("header field must be in the format \"key:value\"", 1)
		}
	}

	req.Close = true

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	if resp != nil {
		defer func(r *http.Response) {
			if r.Body != nil {
				r.Body.Close()
			}
		}(resp)
		if resp.StatusCode != statusCode {
			return cli.NewExitError(fmt.Sprintf("resp code %d didn't match %d", resp.StatusCode, statusCode), 1)
		}
	}
	return nil
}

// globals
var (
	url        string
	httpVerb   string
	statusCode int
)
