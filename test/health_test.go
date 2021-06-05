// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")
	client := resty.New()
	resp, err := client.R().Get(BaseUrl + "/api/health")
	retryCount := 20
	for {
		if err != nil {
			if retryCount == 0 {
				break
			}
			resp, err := client.R().Get(BaseUrl + "/api/health")
			time.Sleep(3 * time.Second)
			retryCount--
		} else {
			break
		}
	}
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}
