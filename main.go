package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"waydo/routes"
)



func SampleHandlerForRoutes(ctx echo.Context) error {
	type Message struct {
		Text string `json:"text"`
	}
	sampleMessage := &Message{Text: "this route handles by function and waiting for write this section"}
	return ctx.JSON(http.StatusOK, sampleMessage)
}

//Models
type GroupIssue struct {
	ID int not null
}


func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	dsn := "user=postgres password=mje1212 dbname=wd port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("connect to postgres database")

	//Routes
	// index route and help for end user for use this application
	e.GET("/", routes.Home)

	//groups is split your mind to general sections for example: life, work, study, and etc...
	//get all groups
	e.GET("/groups", SampleHandlerForRoutes)
	//get issue by id
	e.GET("/groups/:id", SampleHandlerForRoutes)
	//create issue by id
	e.POST("/groups/:id", SampleHandlerForRoutes)
	//delete issue by id
	e.DELETE("/groups/:id", SampleHandlerForRoutes)
	//update issue by id
	e.PUT("/groups/:id", SampleHandlerForRoutes)

	//groups is split your mind to general sections for example: life, work, study, and etc...
	//get all groups
	e.GET("/parents", SampleHandlerForRoutes)
	//get issue by id
	e.GET("/parents/:id", SampleHandlerForRoutes)
	//create issue by id
	e.POST("/parents/:id", SampleHandlerForRoutes)
	//delete issue by id
	e.DELETE("/parents/:id", SampleHandlerForRoutes)
	//update issue by id
	e.PUT("/parents/:id", SampleHandlerForRoutes)



	e.Logger.Fatal(e.Start(":8001"))
}
