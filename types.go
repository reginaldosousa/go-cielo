package cielo

// Merchant identification on Cielo
type Merchant struct {
	ID, Key string
}

// Sale is used to create a sale request
type Sale struct {
	MerchantOrderID string    `json:",omitempty"`
	Customer        *Customer `json:",omitempty"`
	Payment         *Payment  `json:",omitempty"`
}

// Customer handle information about buyer
type Customer struct {
	Name            string   `json:",omitempty"`
	Email           string   `json:",omitempty"`
	BirthDate       string   `json:",omitempty"`
	Identity        string   `json:",omitempty"`
	IdentityType    string   `json:",omitempty"`
	Address         *Address `json:",omitempty"`
	DeliveryAddress *Address `json:",omitempty"`
}

// Address is used to set customer address
type Address struct {
	Street     string `json:",omitempty"`
	Number     string `json:",omitempty"`
	Complement string `json:",omitempty"`
	ZipCode    string `json:",omitempty"`
	City       string `json:",omitempty"`
	State      string `json:",omitempty"`
	Country    string `json:",omitempty"`
}

// Payment represents a payment on Cielo
type Payment struct {
	ServiceTaxAmount    uint32            `json:",omitempty"`
	Installments        uint32            `json:",omitempty"`
	Interest            interface{}       `json:",omitempty"`
	Capture             bool              `json:",omitempty"`
	Authenticate        bool              `json:",omitempty"`
	Recurrent           bool              `json:",omitempty"`
	RecurrentPayment    *RecurrentPayment `json:",omitempty"`
	CreditCard          *CreditCard       `json:",omitempty"`
	Tid                 string            `json:",omitempty"`
	ProofOfSale         string            `json:",omitempty"`
	AuthorizationCode   string            `json:",omitempty"`
	SoftDescriptor      string            `json:",omitempty"`
	ReturnURL           string            `json:",omitempty"`
	Provider            paymentProvider   `json:",omitempty"`
	PaymentID           string            `json:",omitempty"`
	Type                paymentType       `json:",omitempty"`
	Amount              uint32            `json:",omitempty"`
	ReceiveDate         string            `json:",omitempty"`
	CapturedAmount      uint32            `json:",omitempty"`
	CapturedDate        string            `json:",omitempty"`
	Currency            currency          `json:",omitempty"`
	Country             string            `json:",omitempty"`
	ReturnCode          string            `json:",omitempty"`
	ReturnMessage       string            `json:",omitempty"`
	Status              uint32            `json:",omitempty"`
	Links               []interface{}     `json:",omitempty"`
	ExtraDataCollection []interface{}     `json:",omitempty"`
	ExpirationDate      string            `json:",omitempty"`
	URL                 string            `json:",omitempty"`
	Number              string            `json:",omitempty"`
	BarCodeNumber       string            `json:",omitempty"`
	DigitableLine       string            `json:",omitempty"`
	Address             string            `json:",omitempty"`
}

// RecurrentPayment is used to configure recurrent payments
type RecurrentPayment struct {
	AuthorizeNow bool              `json:",omitempty"`
	EndDate      string            `json:",omitempty"`
	Interval     recurrentInterval `json:",omitempty"`
}

type recurrentInterval string
type paymentProvider string
type paymentType string
type currency string
type environment struct {
	APIURL, APIQueryURL string
}

// CreditCard holds credit card informations
type CreditCard struct {
	CardNumber     string `json:",omitempty"`
	Holder         string `json:",omitempty"`
	ExpirationDate string `json:",omitempty"`
	SecurityCode   string `json:",omitempty"`
	SaveCard       bool   `json:",omitempty"`
	Brand          string `json:",omitempty"`
	CardToken      string `json:",omitempty"`
}
