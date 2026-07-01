package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "queen":
		return 10
	case "king":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	valueCard1 := ParseCard(card1)
	valueCard2 := ParseCard(card2)
	valueDealerCard := ParseCard(dealerCard)
	sumOfCards := valueCard1 + valueCard2

	if valueDealerCard == 10 {
		return "S"
	}
	switch {
	case sumOfCards == 21 && valueDealerCard < 10:
		return "W"
	case sumOfCards == 21 && valueDealerCard >= 10:
		return "S"
	case sumOfCards >= 17 && sumOfCards <= 20:
		return "S"
	case sumOfCards >= 12 && sumOfCards <= 16 && valueDealerCard < 7:
		return "S"
	case sumOfCards >= 12 && sumOfCards <= 16 && valueDealerCard >= 7:
		return "H"
	default:
		return "H"
	}
}
