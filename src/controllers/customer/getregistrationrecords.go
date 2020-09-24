package customer


import (
	"github.com/gin-gonic/gin"
	 resty "github.com/go-resty/resty/v2"
	"log"
	"os"	
	// "fmt"
)


func GetRegistrationRecord(c *gin.Context){

	country_code := c.Query("country_code")
	mobile := c.Query("mobile")
	action := "registrations"	
	url :=  os.Getenv("base_url") + action
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", os.Getenv("app_secret")).			
		SetQueryParams(map[string]string{			
			"country_code": country_code,
			"mobile": mobile,			
		}).
		Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		c.Data(200, "application/json", resp.Body())
	
}
