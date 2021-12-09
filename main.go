package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Version: "0.0.5",
		Name:    "Health Checker",
		Usage:   "Hits an endpoint for you.  healthcheck http://localhost/ping",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "headers, H",
				Usage: "specify a header and value for the request (-H=key:value)",
			},
			&cli.StringFlag{
				Name:        "verb, V",
				Usage:       "the HTTP verb to use",
				Value:       "GET",
				EnvVars:     []string{"HEALTHCHECK_VERB"},
				Destination: &httpVerb,
			},
			&cli.IntFlag{
				Name:        "code, C",
				Usage:       "expected response code",
				Value:       http.StatusOK,
				EnvVars:     []string{"HEALTHCHECK_RESPONSECODE"},
				Destination: &statusCode,
			},
			&cli.IntFlag{
				Name:        "timeout, T",
				Usage:       "timeout for HTTP connection",
				Value:       0,
				EnvVars:     []string{"HEALTHCHECK_TIMEOUT"},
				Destination: &timeOut,
			},
			// http body not supported yet
			// response body checking not supported yet
		},
		Action: actionFunc,
	}
	// app.Action, = actionFunc

	app.Run(os.Args)
}

func actionFunc(c *cli.Context) error {
	url = c.Args().Get(0)
	if len(url) == 0 {
		return cli.Exit("url length must be > 0 ", 1)
	}

	req, err := http.NewRequest(httpVerb, url, nil)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	for _, str := range c.StringSlice("headers") {
		kv := strings.Split(str, ":")
		if len(kv) == 2 {
			req.Header.Add(kv[0], kv[1])
		} else {
			return cli.Exit("header field must be in the format \"key:value\"", 1)
		}
	}
	req.Close = true

	var client *http.Client
	if timeOut > 0 {
		timeout := time.Duration(timeOut) * time.Second
		client = &http.Client{Timeout: timeout}
	} else {
		client = &http.Client{}
	}
	resp, err := client.Do(req)
	if err != nil {
		return cli.Exit(err.Error(), 1)
	}
	if resp != nil {
		defer func(r *http.Response) {
			if r.Body != nil {
				r.Body.Close()
			}
		}(resp)
		if resp.StatusCode != statusCode {
			return cli.Exit(fmt.Sprintf("resp code %d didn't match %d", resp.StatusCode, statusCode), 1)
		}
	}
	return nil
}

// globals
var (
	url        string
	httpVerb   string
	statusCode int
	timeOut    int
)
