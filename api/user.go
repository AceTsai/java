package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	params := make(map[string]interface{})
	data, _ := c.GetRawData()
	fmt.Println(data)
	_ = json.Unmarshal(data, &params)

	fmt.Println(params)
}
