package tools

import (
	"HorizonX/model"
	"crypto/rand"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

// RawURLEncode
func RawURLEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

// CalculateDueDate 计算订单的到期时间
func CalculateDueDate(createTime time.Time, billingCycle string) time.Time {
	dueDate := createTime
	switch billingCycle {
	case "monthly":
		return dueDate.AddDate(0, 1, 0)
	case "quarterly":
		return dueDate.AddDate(0, 3, 0)
	case "semiAnnually":
		return dueDate.AddDate(0, 6, 0)
	case "annually":
		return dueDate.AddDate(1, 0, 0)
	default:
		return dueDate
	}
}

// CalculatePrice 计算订单的价格
func CalculatePrice(plan *model.VmPlan, cycle string) int64 {
	switch cycle {
	case "monthly":
		return plan.MonthlyPrice
	case "quarterly":
		return plan.QuarterlyPrice
	case "semiAnnually":
		return plan.SemiAnnuallyPrice
	case "annually":
		return plan.AnnuallyPrice
	default:
		return 0
	}
}

func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func RandomString(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
		'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
		'Y', 'Z'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)

}

// GenerateOrderNo
//
//	eg: 2023080818182604753
func GenerateOrderNo(now time.Time) string {
	date := now.Format("20060102150405")
	r := RandomNumber(5)
	code := fmt.Sprintf("%s%s", date, r)
	return code
}
