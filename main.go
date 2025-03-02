package main

import (
	"fmt"

	"github.com/iammrteapot-learning/balerion-backend-interview/adapters/thaiconverter"
	"github.com/iammrteapot-learning/balerion-backend-interview/domain"
	"github.com/shopspring/decimal"
)

func main() {
	converter := thaiconverter.NewThaiConverter()

	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
		decimal.NewFromFloat(1201.50),
		decimal.NewFromFloat(1220.50),
		decimal.NewFromInt(3222),
		decimal.NewFromFloat(1111.0),
		decimal.NewFromInt(1110),
	}

	for _, input := range inputs {
		money := domain.Money{Amount: input}
		fmt.Println(input)
		thaiText := converter.ConvertToThai(money)
		fmt.Println(thaiText)
	}
}
