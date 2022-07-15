package api

import (
	"github.com/gin-gonic/gin"
	"im/tool"
)

func List(c *gin.Context) {
	x := make(map[string]interface{})
	items := make([]map[string]interface{}, 2)
	x["avatar"] = "https://c-ssl.duitang.com/uploads/item/202105/29/20210529001057_aSeLB.jpeg"
	x["id"] = 22
	x["is_disturb"] = 0
	x["is_online"] = 0
	x["is_robot"] = 1
	x["is_top"] = 1
	x["msg_text"] = "[系统通知] 账号登录提醒！"
	x["name"] = "登录助手"
	x["receiver_id"] = 4257
	x["remark_name"] = ""
	x["talk_type"] = 1
	x["unread_num"] = 0
	x["updated_at"] = "2022-07-15 23:49:30"
	items[0] = x
	r := make(map[string]interface{})
	r["items"] = x

	tool.Success("登录成功", r, c)
}
