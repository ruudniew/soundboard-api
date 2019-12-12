package gin

import (
  "log"
  "strconv"

	"github.com/gin-gonic/gin"
	sbapi "sbapi"
)

type API struct {
	EventService sbapi.EventService
}

func (a *API) Start (host string, port string) {
	// Now load some defaults.
	r := gin.Default()

	// What should the API DO, when POST localhost:3300/event is called?
	r.POST("/event", func(c *gin.Context) {
		evt := sbapi.Event{}
		err := c.BindJSON(&evt)
		if err != nil {
			log.Printf("HELP, I COULD NOT BIND THE RECEIVED DATA TO AN EVENT: %+v", err)
		}
	})

  // What should the API DO, when GET localhost:3300/event is called
  r.GET("/event/:id", func(c *gin.Context) {
    id := c.Param("id")
    
    if id == "" {
      log.Printf("empty ID provided when trying to get event")
      c.AbortWithStatus(400)
      return
    }

    idNumber, err := strconv.Atoi(id)

    if err != nil {
      log.Printf("non-numerical ID provided when trying to get event")
      c.AbortWithStatus(400)
      return
    }

    evt := a.EventService.Get(strconv.Itoa(idNumber))
    c.JSON(200, evt)
  })

  r.Run(host + ":" + port)
}