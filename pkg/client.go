package pkg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// AcceptEncodingKey  is http header key for Accept-encoding.
	AcceptEncodingKey = "Accept-Encoding"
	// UTF8Encoding is UTF-8 encoding value
	UTF8Encoding = "UTF-8"
	// ContentTypeKey is http header key for content type.
	ContentTypeKey = "Content-Type"
	// ContentTypeValueJSON is http header value for application/json.
	ContentTypeValueJSON = "application/json; charset=utf-8"

	defaultTimeout = 10 * time.Second
)

// ClientCallErr is used if the error has been caused by the external Client request execution
var ClientCallErr = errors.New("failed to execute an external service Client call")

// Client handles http requests to the external services
type Client struct {
	HTTPClient *http.Client
}

func NewDefaultClient() Client {
	return Client{HTTPClient: &http.Client{Timeout: defaultTimeout}}
}

// SendRequest to send request to external services
func (c *Client) SendRequest(req *http.Request) (int, io.ReadCloser, error) {
	c.setHeader(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return 0, nil, err
	}

	if err := handleError(resp); err != nil {
		return resp.StatusCode, nil, err
	}

	return resp.StatusCode, resp.Body, nil
}

func (c *Client) setHeader(req *http.Request) {
	req.Header.Add(AcceptEncodingKey, UTF8Encoding)
	req.Header.Add(ContentTypeKey, ContentTypeValueJSON)
}

func handleError(resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	if resp.Body == nil {
		return ClientCallErr
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return WrapErr(fmt.Sprintf("failed to read the response body, %v", err), ClientCallErr)
	}

	return WrapErr(fmt.Sprintf("status code: %d, body '%s'", resp.StatusCode, string(b)), ClientCallErr)
}
