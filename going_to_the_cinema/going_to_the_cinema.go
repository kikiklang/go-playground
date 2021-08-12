package going_to_the_cinema

import "math"

func Movie(card, ticket int, perc float64) int {
	normalPrice := 0
	boughtTickets := 0
	priceWithCard := float64(card)

	for {
		normalPrice += ticket
		boughtTickets += 1
		priceWithCard += float64(ticket) * math.Pow(perc, float64(boughtTickets))

		if math.Ceil(priceWithCard) < float64(normalPrice) {
			return boughtTickets
		}
	}
}
