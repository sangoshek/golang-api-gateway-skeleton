package registration


import (
	"github.com/gin-gonic/gin"
	 resty "github.com/go-resty/resty/v2"
	"log"
	"os"	
	// "encoding/json"
	"fmt"
)

type DataArray struct {
    App_id 	string `json:"app_id"`
    Campaign_id string `json:"campaign_id"`
    Timeslot_id string `json:"timeslot_id"`
    Brand string `json:"brand"`
    Locale string `json:"locale"`
    Last_name string `json:"last_name"`
    First_name string `json:"first_name"`
    Email string `json:"email"`
    Country_code string `json:"country_code"`
    Mobile string `json:"mobile"`
    Has_agreed_terms string `json:"has_agreed_terms"`
    Has_opted_in string `json:"has_opted_in"`
    Utm_source string `json:"utm_source"`
    Utm_medium string `json:"utm_medium"`
    Utm_campaign string `json:"utm_campaign"`
    Utm_term string `json:"utm_term"`
    Utm_content string `json:"utm_content"`
}

type SignArray struct {
    Type 	string `json:"type"`
    Sign    string `json:"sign"`
    
}

type postData struct {
	Data DataArray
	Sign SignArray
}

func Submit(c *gin.Context){
	if c.Request.Method == "POST" {

	var req postData
	c.ShouldBindJSON(&req)	

	// jsonString, _ := json.Marshal(req)
	fmt.Println(req.Data)
	action := "registration"
	url :=  os.Getenv("base_url") + action
	client := resty.New()
	
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", os.Getenv("app_secret")).		
		SetBody(req).
		Get(url)
		if err != nil {
			log.Fatal(err)
		}	
		c.Data(200, "application/json", resp.Body())
	}
}
