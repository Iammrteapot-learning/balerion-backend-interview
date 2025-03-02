package ports

import "github.com/iammrteapot-learning/balerion-backend-interview/domain"

type MoneyConverter interface {
	ConvertToThai(money domain.Money) string
}
