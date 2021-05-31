// +build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBills(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/bills")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestAddBillGoodRequest(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`{"Payer":"Test", "Price":20}`).Post(BASE_URL + "/api/bills")
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestAddBillBadRequest(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetBody(`bad body`).Post(BASE_URL + "/api/bills")
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode())
}
