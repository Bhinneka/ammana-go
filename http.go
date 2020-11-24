package ammana

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// httpRequest data model
type ammanaHttpClient struct {
	httpClient *http.Client
}

// newRequest function for intialize httpRequest object
// Paramter, timeout in time.Duration
func newRequest(timeout time.Duration) *ammanaHttpClient {
	return &ammanaHttpClient{
		httpClient: &http.Client{Timeout: time.Second * timeout},
	}
}

// newReq function for initalize http request,
// paramters, http method, uri path, body, and headers
func (c *ammanaHttpClient) newReq(method string, fullPath string, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// exec private function for call http request
func (c *ammanaHttpClient) exec(method, path string, body io.Reader, v interface{}, headers map[string]string) error {
	req, err := c.newReq(method, path, body, headers)

	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	if v != nil {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}

// execAsync private function for call http request with async
func (c *ammanaHttpClient) execAsync(method, path string, body io.Reader, v interface{}, headers map[string]string) <-chan error {
	output := make(chan error, 1)
	go func() {
		req, err := c.newReq(method, path, body, headers)

		if err != nil {
			output <- err
			return
		}

		res, err := c.httpClient.Do(req)
		defer res.Body.Close()

		if err != nil {
			output <- err
			return
		}

		if v != nil {
			output <- json.NewDecoder(res.Body).Decode(v)
			return
		}

	}()
	return output
}

