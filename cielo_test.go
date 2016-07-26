package cielo_test

import (
	"log"
	"os"
	"testing"

	"github.com/reginaldosousa/go-cielo"
	"github.com/stretchr/testify/assert"
)

var (
	merchantID, merchantKey string
)

func init() {
	merchantID = os.Getenv("CIELO_MERCHANT_ID")
	merchantKey = os.Getenv("CIELO_MERCHANT_KEY")
	if merchantID == "" || merchantKey == "" {
		log.Fatal("You must provide CIELO_MERCHANT_ID and CIELO_MERCHANT_KEY environment variables to run tests")
	}
}

func Test_CieloEcommerce_CreateSale(t *testing.T) {
	m := cielo.Merchant{
		ID:  merchantID,
		Key: merchantKey,
	}
	c := cielo.NewEcommerce(m, cielo.SandboxEnvironment)
	sale := cielo.Sale{
		Payment: &cielo.Payment{
			Type:         cielo.CreditCardPayment,
			Provider:     cielo.SimuladoProvider,
			Amount:       15700,
			Interest:     "ByMerchant",
			Installments: 1,
			CreditCard: &cielo.CreditCard{
				CardNumber:     "1234123412341231",
				Holder:         "John Bla",
				ExpirationDate: "12/2020",
				SecurityCode:   "123",
				Brand:          "Visa",
			},
			Country: "BRA",
		},
		MerchantOrderID: "123A",
		Customer: &cielo.Customer{
			Name: "John",
		},
	}
	s, err := c.CreateSale(sale)
	assert.NoError(t, err)
	assert.NotNil(t, s.Payment.PaymentID)
}
