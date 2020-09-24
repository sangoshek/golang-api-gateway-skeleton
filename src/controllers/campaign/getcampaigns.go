package campaign

import (
	"github.com/gin-gonic/gin"
	 resty "github.com/go-resty/resty/v2"
	"log"
	"os"
	// "fmt"
)

const action string = "campaigns"

func GetCampaigns(c *gin.Context){
		
	url :=  os.Getenv("base_url") + action
	client := resty.New()
	
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", os.Getenv("app_secret")).		
		Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		c.Data(200, "application/json", resp.Body())
	
}


func GetCampaign(c *gin.Context){
	brand := c.Query("brand")
	locale := c.Query("locale")
	url :=  os.Getenv("base_url") + action

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", os.Getenv("app_secret")).		
		SetPathParams(map[string]string{
			   "id":  c.Param("id"),
		}).
		SetQueryParams(map[string]string{			
			"brand": brand,
			"locale": locale,
		}).
		Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		c.Data(200, "application/json", resp.Body())
	
}

