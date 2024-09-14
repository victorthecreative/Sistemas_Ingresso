package domain

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "Available"
	SpotStatusSold      SpotStatus = "Sold"
)

type Spots struct {
	ID       string `json:"id"`
	EventID  string `json:"event_id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	TicketID string `json:"ticket_id"`
}
