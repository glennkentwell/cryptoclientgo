package cryptoclientgo

import (
	"errors"
	"math/big"
	"sort"
)

var ErrMath = errors.New("MATH ERROR, most likely BigInt too large for int64")

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
		//if we are buying we want best sell orders
		return order.SellOrders.getBestSell(amt)
	}
	return order.BuyOrders.getBestBuy(amt)
}

func (o Orders) getBestBuy(amt int64) (int64, error) {
	if !sort.IsSorted(sort.Reverse(o)) {
		sort.Sort(sort.Reverse(o))
	}
	return o.getBest(amt)
}

func (o Orders) getBestSell(amt int64) (int64, error) {
	if !sort.IsSorted(o) {
		sort.Sort(o)
	}
	return o.getBestSell2(amt)
}

func (o Orders) getBestSell2(amt int64) (total int64, err error) {
	BIGamt := big.NewInt(amt)
	BIGtotal := big.NewInt(total)
	temp := big.NewInt(0)
	for _, order := range o {
		BIGVolume := big.NewInt(order.Volume)
		BIGPrice := big.NewInt(order.Price)
		//fmt.Printf("Volume\t\tPrice\t\tamt\t\ttotal\n%+v\t%+v\t%+v\t%+v\n\n", order.Volume, order.Price, BIGamt, total)
		if temp.Div(temp.Mul(BIGVolume, BIGPrice), bigMultiplier).Cmp(BIGamt) >= 0 {
			BIGtotal.Add(BIGtotal, temp.Div(temp.Mul(BIGamt, bigMultiplier), BIGPrice))
			if BIGtotal.IsInt64() {
				return BIGtotal.Int64(), nil
			}
			return 0, ErrMath
		}
		BIGamt.Sub(BIGamt, temp.Div(temp.Mul(BIGVolume, BIGPrice), bigMultiplier))
		BIGtotal.Add(BIGtotal, BIGVolume)
	}
	return 0, errors.New("Not enough volume in open orders")
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
