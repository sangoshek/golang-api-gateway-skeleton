package registration

import (
	"github.com/gin-gonic/gin"
	 resty "github.com/go-resty/resty/v2"
	"log"
	"os"	
	// "fmt"
)

func GetDates(c *gin.Context){
	
	store_id := c.Query("store_id")
	offer_id := c.Query("offer_id")
	campaign_id := c.Query("campaign_id")
	brand := c.Query("brand")
	locale := c.Query("locale")
	
	action := "dates"	
	url :=  os.Getenv("base_url") + action 
	client := resty.New()
	
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", os.Getenv("app_secret")).			
		SetQueryParams(map[string]string{
			"store_id": store_id,
			"offer_id": offer_id,
			"campaign_id": campaign_id,
			"brand": brand,
			"locale": locale,
		}).		
		Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		c.Data(200, "application/json", resp.Body())
	
}