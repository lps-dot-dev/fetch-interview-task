package receipts

import (
	"math"
	"regexp"
	"strconv"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func Process(receipt Receipt) (int, error) {
	score := 0

	// One point for every alphanumeric character in the retailer name
	score += len(nonAlphanumericRegex.ReplaceAllString(receipt.Retailer, ""))

	total, parsingError := strconv.ParseFloat(receipt.Total, 32)
	if parsingError != nil {
		return 0, parsingError
	}

	// 50 points if the total is a round dollar amount with no cents
	if math.Floor(total) == total {
		score += 50
	}

	// 25 points if the total is a multiple of `0.25`
	if math.Mod(total, 0.25) == 0 {
		score += 25
	}

	// 5 points for every two items on the receipt
	score += 5 * int(len(receipt.Items)/2)

	return score, nil
}
