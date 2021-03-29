package card

import ( 
	"github.com/Daria1097/bank/pkg/bank/types"
)

func Issue(currency types.Currency, color string, name string) types.Card {
	card := types.Card {
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:   0,
		Currency:  types.TJS,
		Color:     color,
		Name:      name,
		Active:    true,
	}
	return card
}