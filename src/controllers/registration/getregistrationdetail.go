package registration


import (
	"github.com/gin-gonic/gin"
	 resty "github.com/go-resty/resty/v2"
	"log"
	"os"	
)



func GetRegistrationDetail(c *gin.Context){

	campaign_id := c.Query("campaign_id")
	brand := c.Query("brand")
	redemption_token := c.Query("redemption_token")
	action := "registrations"	
	url :=  os.Getenv("base_url") + action
	client := resty.New()
	
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+os.Getenv("app_secret")).	
		SetPathParams(map[string]string{
			   "id":  c.Param("registration_reference"),
		}).
		SetQueryParams(map[string]string{			
			"campaign_id": campaign_id,
			"brand": brand,
			"redemption_token": redemption_token,			
		}).
		Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		c.Data(200, "application/json", resp.Body())
	
}
