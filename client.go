package cryptoclientgo

//CryptoClient is the generic crypto currency client
type CryptoClient struct{ AbstractClient }

//New gets a new cryptoclient
func New(client AbstractClient) CryptoClient {
	return CryptoClient{client}
}

//AbstractClient is an abstract client for basic cryptoClient methods
type AbstractClient interface {
	//Public
	GetPrimaryCurrencies() ([]string, error)
	GetSecondaryCurrencies() ([]string, error)
	Tick(CurrencyFrom, CurrencyTo string) (Tick, error)
	GetOrderBook(CurrencyFrom, CurrencyTo string) (OrderBook, error)
	GetRecentTrades(CurrencyFrom, CurrencyTo string, historyAmount int) (RecentTrades, error)
	//ExpectedMarketValueBuy(PrimaryCurrency, SecondaryCurrency string, amountOfToCurrency int64) (int64, error)
	//ExpectedMarketValueSell(PrimaryCurrency, SecondaryCurrency string, amountOfFromCurrency int64) (int64, error)
	//Private
	PlaceLimitBuyOrder(CurrencyFrom, CurrencyTo string, amount int64, price int64) (PlacedOrder, error)
	PlaceMarketBuyOrder(CurrencyFrom, CurrencyTo string, amount int64) (PlacedOrder, error)
	PlaceLimitSellOrder(CurrencyFrom, CurrencyTo string, amount int64, price int64) (PlacedOrder, error)
	PlaceMarketSellOrder(CurrencyFrom, CurrencyTo string, amount int64) (PlacedOrder, error)
	CancelOrder(OrderID int) error
	GetOrderDetails(OrderID int) (OrderDetails, error)
	GetOpenOrders() (OrdersDetails, error)
	GetBalance(Currency string) (AccountBalance, error)
	GetBalances() (AccountBalances, error)
	GetPrimaryCurrencyDepositAddress(Currency string) (CurrencyAddress, error)
	WithdrawCurrency(Currency, to string, amount int64) error
	//Custom
	GetTransactionCost(CurrencyFrom, CurrencyTo string) (Cost, error)
	GetWithdrawCost(Currency string) (Cost, error)
	GetDepositCost(Currency string) (Cost, error)
}
