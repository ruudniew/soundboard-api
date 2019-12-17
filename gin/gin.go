package gin

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	sbapi "sbapi"
)

type API struct {
	EventService sbapi.EventService
}

func (a *API) Start(host string, port string, corsHosts []string, pass string) {
	// Now load some defaults.
	r := gin.Default()

	// CORS middlware for gin.
	r.Use(cors.New(cors.Config{
		AllowOrigins:     corsHosts,
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// What should the API DO, when POST localhost:3300/event is called?
	r.POST("/"+pass+"/event", func(c *gin.Context) {
		evt := sbapi.Event{}

		err := c.BindJSON(&evt)
		if err != nil {
			log.Printf("HELP, I COULD NOT BIND THE RECEIVED DATA TO AN EVENT: %+v", err)
			c.AbortWithStatus(400) // bad request
			return
		}
		log.Printf("evt received: %+v", evt.Source)

		id, err := a.EventService.Save(&evt)

		if err != nil {
			log.Printf("HELP, I COULD NOT SAVE THIS DATA TO AN EVENT: %+v", err)
			c.AbortWithStatus(400)
			return
		}

		c.JSON(200, id)
	})

	// What should the API DO, when GET localhost:3300/event/:id is called
	r.GET("/"+pass+"/event/:id", func(c *gin.Context) {
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

	r.GET("/"+pass+"/events/:time", func(c *gin.Context) {
		t := c.Param("time")

		if t == "" {
			log.Printf("no time provided, when trying to get event list")
			c.AbortWithStatus(400)
			return
		}

		// Test with time parameter.
		//fifteenMinutesAgo, _ := nd.Parse("15 minutes ago", time.Now())
		//evts := a.EventService.GetList(fifteenMinutesAgo.String())

		evts := a.EventService.GetList(t)

		c.JSON(200, evts)
	})

	r.Run(host + ":" + port)
}
