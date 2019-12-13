package sbapi

type API interface {
	Start(host string, port string, corsHosts []string)
}

type Event struct {
	ID        string       `json:"id"`
	CreatedAt string       `json:"created_at"`
	Title     string       `json:"title"` // Example: slow-clap
	Source    *EventSource `json:"source"`
}

type EventSource struct {
	Name        string `json:"name"`
	TriggeredBy string `json:"triggered_by"`
}

type EventService interface {
	Get(id string) *Event
	GetList(time string) []*Event
	Save(evt *Event) (id string, err error)
}
