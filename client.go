package cryptoclientgo

import "time"

//Multiplier is the multiplier used when dealing with amounts or prices, to avoid floating point errors
const Multiplier = int64(100000000)

//ConvertToFloat converts an int64 to the corresponding float64
func ConvertToFloat(i int64) float64 {
	return float64(i) / float64(Multiplier)
}

//ConvertFromFloat converts an float64 to the corresponding int64
func ConvertFromFloat(f float64) int64 {
	return int64(f) * Multiplier
}

//CryptoClient is the generic crypto currency client
type CryptoClient interface {
	//Public
	GetAccountCurrencies() ([]string, error)
	Tick(CurrencyFrom, CurrencyTo string) (Tick, error)
	GetOrderBook(CurrencyFrom, CurrencyTo string) (OrderBook, error)
	GetRecentTrades(CurrencyFrom, CurrencyTo string, historyAmount int) (RecentTrades, error)
	//Private
	PlaceLimitOrder(CurrencyFrom, CurrencyTo string, amount int64, price int64) (OrderDetails, error)
	PlaceMarketOrder(CurrencyFrom, CurrencyTo string, amount int64) (OrderDetails, error)
	CancelOrder(OrderID int) error
	GetOrderDetails(OrderID int) (OrderDetails, error)
	GetOpenOrders() (OrdersDetails, error)
	GetBalance(Currency string) (AccountBalance, error)
	GetBalances() (AccountBalances, error)
	GetDigitalCurrencyDepositAddress(Currency string) (CurrencyAddress, error)
	WithdrawCurrency(Currency, to string, amount int64) error
	//Custom
	GetTransactionCost(CurrencyFrom, CurrencyTo string) (Cost, error)
	GetWithdrawCost(Currency string) (Cost, error)
	GetDepositCost(Currency string) (Cost, error)
}

//Cost is the cost associated with making a specfic transaction
type Cost struct {
	Flat    int64
	Percent int64
}

//Tick is the current status of the market
type Tick struct {
	PrimaryCurrency         string
	SecondaryCurrency       string
	CurrentHighestBidPrice  int64
	CurrentLowestOfferPrice int64
	LastPrice               int64
}

//OrderBook gets the current open orders
type OrderBook struct {
	PrimaryCurrency   string
	SecondaryCurrency string
	BuyOrders         Orders
	SellOrders        Orders
}

//Orders encapsulates multiple instances of an Order
type Orders []Order

//Order encapsulates the information of the order places
type Order struct {
	Price  int64
	Volume int64
}

//RecentTrades returns the most recent trades between the two specified currencies
type RecentTrades struct {
	PrimaryCurrency   string
	SecondaryCurrency string
	Timestamp         time.Time //ms
	Trades            Trades
}

//Trades is a collection of singular instances of Trade
type Trades []Trade

//Trade is a single trade between two currencies
type Trade struct {
	PrimaryAmount   int64
	SecondaryAmount int64
}

//OrdersDetails is a list of singular OrderDetails
type OrdersDetails []OrderDetails

//OrderDetails encapsulates the details of an order
type OrderDetails struct {
	PrimaryCurrency   string
	SecondaryCurrency string
	OrderID           int64
	Created           time.Time
	//OrderSide Bid/Ask
	OrderSide string
	//OrderType Limit/Market
	OrderType string
}

//AccountBalances is a list of all available AccountBalance(s)
type AccountBalances []AccountBalance

//AccountBalance is used to show the amount available in a specific account
type AccountBalance struct {
	Currency         string
	AvailableBalance int64
	TotalBalance     int64
}

//CurrencyAddress is the address to recieve cryptocurrency in
type CurrencyAddress struct {
	DepositAddress          string
	LastCheckedTimestampUtc time.Time
	NextUpdateTimestampUtc  time.Time
}
