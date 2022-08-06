package buxfer

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL)
	resp, err := http.DefaultTransport.RoundTrip(r)
	fmt.Fprintf(l.logger, "[%s] %s\n", time.Now().Format(time.ANSIC), resp.Status)
	return resp, err
}
