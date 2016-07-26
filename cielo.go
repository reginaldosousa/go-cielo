// Package cielo is used to communicate with Cielo API
package cielo

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/satori/go.uuid"
)

// CieloEcommerce is the interface used to communicate with Cielo v3 API
type Ecommerce interface {
	CreateSale(sale Sale) (Sale, error)
}

type cieloEcommerce struct {
	merchant Merchant
	env      environment
}

func (c *cieloEcommerce) CreateSale(s Sale) (sale Sale, err error) {
	body, err := json.Marshal(s)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPost, c.env.APIURL+"/1/sales", bytes.NewBuffer(body))
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "CieloEcommerce/3.0 Golang client")
	req.Header.Add("MerchantId", c.merchant.ID)
	req.Header.Add("MerchantKey", c.merchant.Key)
	req.Header.Add("RequestId", uuid.NewV4().String())

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	err = json.NewDecoder(response.Body).Decode(&sale)
	return
}

// NewEcommerce create a Cielo v3 API client
func NewEcommerce(m Merchant, env environment) Ecommerce {
	return &cieloEcommerce{
		merchant: m,
		env:      env,
	}
}
