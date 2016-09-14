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
	merchant                cielo.Merchant
	ecommerce               cielo.Ecommerce
)

func init() {
	merchantID = os.Getenv("CIELO_MERCHANT_ID")
	merchantKey = os.Getenv("CIELO_MERCHANT_KEY")
	if merchantID == "" || merchantKey == "" {
		log.Fatal("You must provide CIELO_MERCHANT_ID and CIELO_MERCHANT_KEY environment variables to run tests")
	}
	merchant = cielo.Merchant{
		ID:  merchantID,
		Key: merchantKey,
	}
	ecommerce = cielo.NewEcommerce(merchant, cielo.SandboxEnvironment)
}

func Test_CieloEcommerce_CreateSale(t *testing.T) {
	testCases := []struct {
		name    string
		getSale func() cielo.Sale
	}{
		{
			name: "basic sale with success",
			getSale: func() cielo.Sale {
				payment := newPayment(15700, newCreditCard("1234123412341231", "12/2020", "123"))
				customer := newCustomer("John")
				sale := cielo.Sale{
					Payment:         &payment,
					MerchantOrderID: "123A",
					Customer:        &customer,
				}
				return sale
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := ecommerce.CreateSale(tc.getSale())
			assert.NoError(t, err)
			assert.NotNil(t, s.Payment.PaymentID)
		})
	}
}

func newPayment(amount uint32, creditCard cielo.CreditCard) cielo.Payment {
	return cielo.Payment{
		Type:         cielo.CreditCardPayment,
		Provider:     cielo.SimuladoProvider,
		Amount:       amount,
		Interest:     "ByMerchant",
		Installments: 1,
		CreditCard:   &creditCard,
		Country:      "BRA",
	}
}

func newCreditCard(number, expiration, securityCode string) cielo.CreditCard {
	return cielo.CreditCard{
		CardNumber:     number,
		Holder:         "John Bla",
		ExpirationDate: expiration,
		SecurityCode:   securityCode,
		Brand:          "Visa",
	}
}

func newCustomer(name string) cielo.Customer {
	return cielo.Customer{
		Name: name,
	}
}
