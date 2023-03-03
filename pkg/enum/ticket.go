package enum

type TicketStatus string

const (
	Todo       TicketStatus = "todo"
	InProgress TicketStatus = "on-progress"
	Done       TicketStatus = "done"
)

func (t TicketStatus) String() string {
	return string(t)
}
