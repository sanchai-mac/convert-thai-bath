package convert

import (
	"fmt"
	"math"
	"strings"
)

func ThaiBahtText(amount float64) string {
	amount = math.Round(amount*100) / 100 //TODO: Round to 2 places
	parts := strings.Split(fmt.Sprintf("%.2f", amount), ".")
	bahtText := integerToThaiText(parts[0]) + "บาท"

	if parts[1] == "00" {
		bahtText += "ถ้วน"
	} else {
		satangText := convertGroup(parts[1]) + "สตางค์"
		bahtText += satangText
	}

	return bahtText
}

// TODO: Convert integers to Thai text (supports trillions)
func integerToThaiText(number string) string {
	if number == "0" {
		return "ศูนย์"
	}
	var result string
	groups := []string{}
	for len(number) > 0 {
		end := len(number)
		start := end - 6
		if start < 0 {
			start = 0
		}
		groups = append([]string{number[start:end]}, groups...)
		number = number[:start]
	}

	for i, group := range groups {
		groupText := convertGroup(group)
		if groupText != "" {
			result += groupText
			if i < len(groups)-1 {
				result += UnitMillion
			}
		}
	}
	return result
}

// TODO: Convert groups of 6 or less digits to text.
func convertGroup(number string) string {
	n := len(number)
	result := ""
	for i, ch := range number {
		digit := int(ch - '0')
		pos := n - i - 1
		if digit == 0 {
			continue
		}
		switch {
		case pos == 1 && digit == 1:
			result += "สิบ"
		case pos == 1 && digit == 2:
			result += "ยี่สิบ"
		case pos == 1:
			result += ThaiDigits[digit] + "สิบ"
		case pos == 0 && digit == 1 && n > 1:
			result += "เอ็ด"
		default:
			result += ThaiDigits[digit] + ThaiPlaces[pos]
		}
	}
	return result
}
