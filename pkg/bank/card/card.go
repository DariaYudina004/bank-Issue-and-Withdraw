package card

import ( 
	"github.com/Daria1097/bank/pkg/bank/types"
)

 
//Withdraw рассчитывает баланс после снятия денег
func Withdraw(card *types.Card,amount types.Money){
	const withdrawLimit = 20_000_00
    if !card.Active {
		return 
	}
	if amount >card.Balance {
		return 
	}
	if amount>  withdrawLimit  {
		return 
	}
	if amount < 0 {
		return 
	}
	card.Balance-=amount
}

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