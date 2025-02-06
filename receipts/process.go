package receipts

import (
	"math"
	"regexp"
	"strconv"
	"strings"
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

	/**
	 * If the trimmed length of the item description is a multiple of 3,
	 * multiply the price by 0.2 and round up to the nearest integer.
	 * The result is the number of points earned.
	 */
	for _, receiptItem := range receipt.Items {
		trimmedDescription := strings.TrimSpace(receiptItem.ShortDescription)
		if len(trimmedDescription)%3 != 0 {
			continue
		}

		itemPrice, parsingError := strconv.ParseFloat(receiptItem.Price, 32)
		if parsingError != nil {
			return 0, parsingError
		}

		score += int(math.Ceil(itemPrice / 5))
	}

	return score, nil
}
