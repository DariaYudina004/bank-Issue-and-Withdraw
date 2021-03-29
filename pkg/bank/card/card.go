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