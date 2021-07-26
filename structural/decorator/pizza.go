package decorator

type Pizza interface {
	GetPrice() int
}

type VeggieMania struct{}

func (m VeggieMania) GetPrice() int {
	return 20
}

type PeppyPaneer struct{}

func (p PeppyPaneer) GetPrice() int {
	return 15
}

type CheeseTopping struct {
	p Pizza
}

func (t CheeseTopping) GetPrice() int {
	price := 7
	if t.p != nil {
		price += t.p.GetPrice()
	}
	return price
}

type MeatTopping struct {
	p Pizza
}

func (t MeatTopping) GetPrice() int {
	price := 10
	if t.p != nil {
		price += t.p.GetPrice()
	}
	return price
}
