package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"im/tool"
)

func Login(c *gin.Context) {
	params := make(map[string]string)
	data, _ := c.GetRawData()
	fmt.Println(data)
	_ = json.Unmarshal(data, &params)
	t, _ := tool.GenerateToke(params["nickname"])
	r := make(map[string]interface{})
	r["token"] = t
	r["uid"] = 111
	r["nickname"] = "阿斯顿"
	r["avatar"] = "https://c-ssl.duitang.com/uploads/item/202105/29/20210529001057_aSeLB.jpeg"
	tool.Success("登录成功", r, c)
}

func Info(c *gin.Context) {
	r := make(map[string]interface{})
	r["uid"] = 111
	r["is_qiye"] = true
	r["email"] = "asds"
	r["gender"] = "asds"
	r["moto"] = "13"
	r["nickname"] = "阿斯顿"
	r["avatar"] = "https://c-ssl.duitang.com/uploads/item/202105/29/20210529001057_aSeLB.jpeg"
	x := make(map[string]interface{})
	x["user_info"] = r
	tool.Success("登录成功", x, c)

}
