package cryptoclientgo

import (
	"errors"
	"sort"
)

//ExpectedMarketValueBuy gets the expected value cost from PrimaryCurrency to achieve the specified amount of SecondaryCurrency
func (c CryptoClient) ExpectedMarketValueBuy(PrimaryCurrency, SecondaryCurrency string, amountOfSecondary int64) (int64, error) {
	return c.expectedMarketValue(PrimaryCurrency, SecondaryCurrency, amountOfSecondary, true)
}

//ExpectedMarketValueSell gets the expected value recieved of PrimaryCurrency by selling the specified amount of SecondaryCurrency
func (c CryptoClient) ExpectedMarketValueSell(PrimaryCurrency, SecondaryCurrency string, amountOfSecondary int64) (int64, error) {
	return c.expectedMarketValue(PrimaryCurrency, SecondaryCurrency, amountOfSecondary, false)
}

func (c CryptoClient) expectedMarketValue(PrimaryCurrency, SecondaryCurrency string, amt int64, buy bool) (int64, error) {
	order, err := c.GetOrderBook(PrimaryCurrency, SecondaryCurrency)
	if err != nil {
		return 0, errors.New("Failed to get open orders;" + err.Error())
	}
	if buy {
		return order.BuyOrders.getBestBuy(amt)
	}
	return order.SellOrders.getBestSell(amt)
}

func (o Orders) getBestBuy(amt int64) (int64, error) {
	if !sort.IsSorted(o) {
		sort.Sort(o)
	}
	return o.getBest(amt)
}
func (o Orders) getBestSell(amt int64) (int64, error) {
	if !sort.IsSorted(sort.Reverse(o)) {
		sort.Sort(sort.Reverse(o))
	}
	return o.getBest(amt)
}

//getBest Assumes sorted then gets the best
func (o Orders) getBest(amt int64) (total int64, err error) {
	for _, order := range o {
		if order.Volume >= amt {
			return total + (((amt / 100) * (order.Price / 100)) / (Multiplier / 10000)), nil
		}
		amt -= order.Volume
		total += (((order.Volume / 100) * (order.Price / 100)) / (Multiplier / 10000))
	}
	return 0, errors.New("Not enough volume in open orders")
}
