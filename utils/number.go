package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func NullIntScan(a *int) int {
	if a != nil {
		return *a
	}

	return 0
}

func NullFloatScan(a *float64) float64 {
	if a != nil {
		return *a
	}

	return 0.0
}

func NullFloat32Scan(a *float32) float32 {
	if a != nil {
		return *a
	}
	return 0
}

func ScanIntToNullValue(a int) *int {
	if a == 0 {
		return nil
	}

	return &a
}

func NullFloat64ScanFromNullableString(a *string) float64 {
	if a != nil {
		value, err := strconv.ParseFloat(*a, 64)
		if err != nil {
			return 0.0
		}
		return value
	}
	return 0.0
}

func CountTotalPage(total, perPage int) int {
	if (total % perPage) > 0 {
		return (total / perPage) + 1
	} else {
		return total / perPage
	}
}

func CommaSeparated(v float64) string {
	sign := ""

	// Min float64 can't be negated to a usable value, so it has to be special cased.
	if v == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = fmt.Sprintf("%.0f", math.Floor(math.Mod(v, 1000)))
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], ",")
}
