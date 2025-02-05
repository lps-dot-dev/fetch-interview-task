package receipts

import "regexp"

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func Process(receipt Receipt) int {
	score := 0

	// One point for every alphanumeric character in the retailer name
	score += len(nonAlphanumericRegex.ReplaceAllString(receipt.Retailer, ""))

	return score
}
