package cryptoclientgo

import "testing"

func TestGetBest(t *testing.T) {
	o := getOrders()
	amt, err := o.getBestBuy(4 * Multiplier)
	if err != nil {
		t.Error("Err not nil")
	}
	want := (1*1 + 1*2 + 1*2) * Multiplier
	if amt != want {
		t.Errorf("Amt:\t%+v\nWant:\t%+v", amt, want)
	}
}

func TestGetSell(t *testing.T) {
	o := getOrders()
	amt, err := o.getBestSell(5 * Multiplier)
	if err != nil {
		t.Error("Err not nil")
	}
	want := (5*4 + 2*1) * Multiplier
	if amt != want {
		t.Errorf("Amt:\t%+v\nWant:\t%+v", amt, want)
	}
}

func getOrders() Orders {
	o := Orders{
		{1, 1},
		{1, 2},
		{5, 4},
		{2, 2},
	}

	for i := range o {
		o[i].Price *= Multiplier
		o[i].Volume *= Multiplier
	}
	return o
}
