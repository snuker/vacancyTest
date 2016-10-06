package main

import ("github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"strings"
	)
type Request struct {
    Site []string `json:"Site" binding:"required"`
    SearchText string `json:"SearchText" binding:"required"`
}
type Response struct {
    FoundAtSite string `json:"FoundAtSite"`
}
func main() {
    r := gin.Default()
    r.POST("/checkText", func(c *gin.Context) {
        var req Request
	err :=  c.BindJSON(&req)
	if err != nil {
		c.String(400,"Bad request")
		}
		
	for _,url := range req.Site{
		res, err := http.Get(url)
		if err!=nil { 
			continue;
		}
		body,err:=ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err!=nil {
			continue;			
		}
		if strings.Contains(string(body),req.SearchText){
			c.JSON(200,gin.H {"FoundAtSite":url})		
			return;
		}			

	}
	c.String(204,"No Content")
	
    })
    r.Run() 
}
