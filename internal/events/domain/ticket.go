package domain

type TicketType string

const (
	TicketTypeHalf TicketType = "Half"
	TicketTypeFull TicketType = "Full"
)

type Tickets struct {
	ID         string     `json:"id"`
	EventID    string     `json:"event_id"`
	Spot       *Spots     `json:"spot"`
	TicketType TicketType `json:"ticket_type"`
	Price      float64    `json:"price"`
}
