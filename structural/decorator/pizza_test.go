package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecorator(t *testing.T) {

	mania := VeggieMania{}
	assert.Equal(t, mania.GetPrice(), 20)

	maniaWithMeat := MeatTopping{
		p: mania,
	}
	assert.Equal(t, maniaWithMeat.GetPrice(), 30)

	maniaWithMeatCheese := CheeseTopping{
		p: maniaWithMeat,
	}
	assert.Equal(t, maniaWithMeatCheese.GetPrice(), 37)

	peppy := PeppyPaneer{}
	assert.Equal(t, peppy.GetPrice(), 15)

	peppyWithCheese := CheeseTopping{
		p: peppy,
	}
	assert.Equal(t, peppyWithCheese.GetPrice(), 22)
}
