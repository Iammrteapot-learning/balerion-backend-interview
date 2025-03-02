package thaiconverter

import (
	"strings"

	"github.com/iammrteapot-learning/balerion-backend-interview/domain"
	"github.com/shopspring/decimal"
)

var thaiNumbers = map[int]string{
	0: "ศูนย์",
	1: "หนึ่ง",
	2: "สอง",
	3: "สาม",
	4: "สี่",
	5: "ห้า",
	6: "หก",
	7: "เจ็ด",
	8: "แปด",
	9: "เก้า",
}

var thaiUnits = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}

type ThaiConverter struct{}

func NewThaiConverter() *ThaiConverter {
	return &ThaiConverter{}
}

func (c *ThaiConverter) ConvertToThai(money domain.Money) string {
	amount := money.Amount
	if amount.Equal(decimal.Zero) {
		return "ศูนย์บาท"
	}

	// Split into baht and satang
	baht := amount.Floor()
	satang := amount.Sub(baht).Mul(decimal.NewFromInt(100)).Floor()

	result := []string{}

	if baht.GreaterThan(decimal.Zero) {
		result = append(result, convertNumber(int(baht.IntPart())))
		result = append(result, "บาท")
	}

	if satang.GreaterThan(decimal.Zero) {
		result = append(result, convertNumber(int(satang.IntPart())))
		result = append(result, "สตางค์")
	} else {
		result = append(result, "ถ้วน")
	}

	return strings.Join(result, "")
}

func convertNumber(num int) string {
	if num == 0 {
		return thaiNumbers[0]
	}

	result := []string{}
	pos := 0 // หน่วย

	for num > 0 {
		digit := num % 10
		if digit != 0 {
			unit := thaiUnits[pos%7]
			number := thaiNumbers[digit]

			// Special cases
			if pos == 1 { // Tens position
				if digit == 1 {
					number = ""
				} else if digit == 2 {
					number = "ยี่"
				}
			}

			if pos == 0 && digit == 1 { // Ones position
				number = "เอ็ด"
				if num > 100 && (num/10)%10 == 0 {
					number = "หนึ่ง"
				}
			}

			result = append([]string{number + unit}, result...)
		}
		pos++
		num = num / 10
	}

	return strings.Join(result, "")
}
