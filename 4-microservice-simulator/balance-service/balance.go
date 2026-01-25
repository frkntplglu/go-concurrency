package balanceservice

import (
	"math/rand"
	"time"
)

type BalanceService struct {
}

func New() *BalanceService {
	return &BalanceService{}
}

func (s *BalanceService) CheckBalance(balance chan string) {
	n := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(n))

	balance <- "Your balance is 1500€"

}
