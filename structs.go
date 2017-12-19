package cryptoclientgo

import (
	"sort"
	"time"
)

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

type TickFloat struct {
	PrimaryCurrency			string
	SecondaryCurrency		string
	CurrentHighestBidPrice	float64
	CurrentLowestOfferPrice float64
	LastPrice				float64
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

//SortBuy is a helper method to sort the orders from cheapest to most expensive
func (o *Orders) SortBuy() {
	sort.Sort(o)
}

//SortSell is a helper mthod to sort the orders from most expensive to cheapest
func (o *Orders) SortSell() {
	sort.Sort(sort.Reverse(o))
}

// Len is the number of elements in the collection.
func (o Orders) Len() int {
	return len(o)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (o Orders) Less(i, j int) bool {
	return o[i].Price < o[j].Price
}

// Swap swaps the elements with indexes i and j.
func (o Orders) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

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
	Amount int64
	Price  int64
}

//OrdersDetails is a list of singular OrderDetails
type OrdersDetails []OrderDetails

//OrderDetails encapsulates the details of an order
type OrderDetails struct {
	PrimaryCurrency   string
	SecondaryCurrency string
	OrderID           int64
	Created           time.Time
	VolumeOrdered     int64
	VolumeFilled      int64
	Price             int64
	//OrderSide Bid/Ask
	OrderSide string
	//OrderType Limit/Market
	OrderType string
}

//PlacedOrder is a basic order which contains the order id from a successful request
type PlacedOrder struct {
	OrderID int
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
