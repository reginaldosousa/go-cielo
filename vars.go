package cielo

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

var (
	// ProductionEnvironment sets the environment to production
	ProductionEnvironment = environment{
		APIURL:      "https://api.cieloecommerce.cielo.com.br/",
		APIQueryURL: "https://apiquery.cieloecommerce.cielo.com.br/",
	}
	// SandboxEnvironment sets the environment to sandbox
	SandboxEnvironment = environment{
		APIURL:      "https://apisandbox.cieloecommerce.cielo.com.br/",
		APIQueryURL: "https://apiquerysandbox.cieloecommerce.cielo.com.br/",
	}
)
