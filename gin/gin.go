package gin

import (
	"github.com/gin-gonic/gin"
	sbapi "soundboard-api"
)

type API struct {
	EventService sbapi.EventService
}

func (a *API) Start (es sbapi.EventService, host string, port string) {
	// Start, but attach the EventService first.
	a.EventService = es

	// Now load some defaults.
	r := gin.Default()

	// What should the API DO, when localhost:3300/saveEvent is called.
	r.POST("/saveEvent", func(c *gin.Context) {
		evt := sbapi.Event{}
		err := c.BindJSON(&evt)
		if err != nil {
			panic("HELP, I COULD NOT BIND THE RECEIVED DATA TO AN EVENT: " + err.Error())
		}

		// Save the event
		// HOW?
		// That's not GIN's business
		a.EventService.Save(&evt)
	})
}