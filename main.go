package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"main/src/controllers/campaign"
	"main/src/controllers/customer"
	"main/src/controllers/registration"
	"main/src/helpers/sha1hash"	
	"os"
	"fmt"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	app_id := os.Getenv("app_id")
	r := gin.Default()
	sha := sha1hash.Make(app_id)
	fmt.Println(sha)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/campaigns", campaign.GetCampaigns)
		v1.GET("/campaigns/:id", campaign.GetCampaign)		
		v1.GET("/offers", registration.GetOffers)
		v1.GET("/stores", registration.GetStores)
		v1.GET("/dates", registration.GetDates)
		v1.GET("/times", registration.GetTimes)
		v1.GET("/registrations", customer.GetRegistrationRecord)
		v1.GET("/registrations/:registration_reference", registration.GetRegistrationDetail)
		v1.POST("/registration", registration.Submit)
	}	

	r.Run(":9000")
}