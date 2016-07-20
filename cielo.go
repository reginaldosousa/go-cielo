// Package cielo is used to communicate with Cielo API
package cielo

// Merchant identification on Cielo
type Merchant struct {
	ID, Key string
}

// Environment configuration interface
type Environment interface {
	GetApiUrl() string
	GetQueryApiUrl() string
}

// Sale is used to create a sale request
type Sale struct {
	MerchantOrderID string
	Customer        Customer
	Payment         Payment
}

// Customer handle information about buyer
type Customer struct {
	Name            string
	Email           string
	BirthDate       string
	Identity        string
	IdentityType    string
	Address         Address
	DeliveryAddress Address
}

// Address is used to set customer address
type Address struct {
	Street     string
	Number     string
	Complement string
	ZipCode    string
	City       string
	State      string
	Country    string
}

// Payment represents a payment on Cielo
type Payment struct {
	ServiceTaxAmount    uint32
	Installments        uint32
	Interest            string
	Capture             bool
	Authenticate        bool
	Recurrent           bool
	RecurrentPayment    RecurrentPayment
	CreditCard          CreditCard
	Tid                 string
	ProofOfSale         string
	AuthorizationCode   string
	SoftDescriptor      string
	ReturnURL           string
	Provider            paymentProvider
	PaymentID           string
	Type                paymentType
	Amount              uint32
	ReceiveDate         string
	CapturedAmount      uint32
	CapturedDate        string
	Currency            currency
	Country             string
	ReturnCode          string
	ReturnMessage       string
	Status              uint32
	Links               []interface{}
	ExtraDataCollection []interface{}
	ExpirationDate      string
	URL                 string
	Number              string
	BarCodeNumber       string
	DigitableLine       string
	Address             string
}

// RecurrentPayment is used to configure recurrent payments
type RecurrentPayment struct {
	AuthorizeNow bool
	EndDate      string
	Interval     recurrentInterval
}

type recurrentInterval string
type paymentProvider string
type paymentType string
type currency string

const (
	// MonthlyInterval configures recurrent payment interval to Montly
	MonthlyInterval = recurrentInterval("Monthly")
	// BimonthlyInterval configures recurrent payment interval to Bimonthly
	BimonthlyInterval = recurrentInterval("Bimonthly")
	// QuarterlyInterval configures recurrent payment interval to Quarterly
	QuarterlyInterval = recurrentInterval("Quarterly")
	// SemiAnnualInterval configures recurrent payment interval to SemiAnnual
	SemiAnnualInterval = recurrentInterval("SemiAnnual")
	// AnnualInterval configures recurrent payment interval to Annual
	AnnualInterval = recurrentInterval("Annual")

	// BradescoProvider is the Bradesco payment provider
	BradescoProvider = paymentProvider("Bradesco")
	// BancoDoBrasilProvider is the Banco do Brasil payment provider
	BancoDoBrasilProvider = paymentProvider("BancoDoBrasil")
	// SimuladoProvider is the Simulated payment provider
	SimuladoProvider = paymentProvider("Simulado")

	// CreditCardPayment is used to set payment method to Credit Card
	CreditCardPayment = paymentType("CreditCard")
	// DebitCardPayment is used to set payment method to Debit Card
	DebitCardPayment = paymentType("DebitCard")
	// ElectronicTransferPayment is used to set payment method to Electronic Transfer
	ElectronicTransferPayment = paymentType("ElectronicTransfer")
	// BoletoPayment is used to set payment method to Boleto
	BoletoPayment = paymentType("Boleto")

	// CurrencyBRL set current to Brazilian Real
	CurrencyBRL = currency("BRL")
	// CurrencyUSD set current to USA Dollar
	CurrencyUSD = currency("USD")
	// CurrencyMXN set current to Mexican Peso
	CurrencyMXN = currency("MXN")
	// CurrencyCOP set current to Colombian Peso
	CurrencyCOP = currency("COP")
	// CurrencyCLP set current to Chilean Peso
	CurrencyCLP = currency("CLP")
	// CurrencyARS set current to Argentine Peso
	CurrencyARS = currency("ARS")
	// CurrencyPEN set current to Peruvian Sol
	CurrencyPEN = currency("PEN")
	// CurrencyEUR set current to Euro
	CurrencyEUR = currency("EUR")
	// CurrencyPYN set current to PYN
	CurrencyPYN = currency("PYN")
	// CurrencyUYU set current to Uruguayan Peso
	CurrencyUYU = currency("UYU")
	// CurrencyVEB set current to Venezuelan Bolivar
	CurrencyVEB = currency("VEB")
	// CurrencyVEF set current to Venezuelan Bolivar
	CurrencyVEF = currency("VEF")
	// CurrencyGBP set current to British Pound
	CurrencyGBP = currency("GBP")
)

// CreditCard holds credit card informations
type CreditCard struct {
	CardNumber     string
	Holder         string
	ExpirationDate string
	SecurityCode   string
	SaveCard       bool
	Brand          string
	CardToken      string
}
