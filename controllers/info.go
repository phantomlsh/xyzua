package controllers

import (
	"encoding/json"

	"xyzua/services"
	"xyzua/utils"

	"github.com/astaxie/beego"
)

// InfoController operations for info
type InfoController struct {
	beego.Controller
}

// URLMapping ...
func (c *InfoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Create and issue code
// @router / [post]
func (c *InfoController) Post() {
	var input map[string]interface{}
	var resp utils.SimpleResponse
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		if input["token"] == nil || input["appname"] == nil {
			resp.Success = false
			resp.Message = utils.Msg.ParamErr + "(token & appname required)"
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		token := input["token"].(string)
		appname := input["appname"].(string)
		resp.Success, resp.Message = services.CheckToken(token)
		// pass CheckToken
		if resp.Success == true {
			identity := services.QueryIdentity(resp.Message.(string), false)
			callback := services.QueryCode(identity.Id, appname)
			if callback == nil {
				resp.Success = false
			} else {
				resp.Success = true
			}
			callback.Token = resp.Message.(string)
			resp.Message = callback
		}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Testing
// @router / [get]
func (c *InfoController) GetAll() {
	c.Ctx.WriteString("InfoController - XYZUA")
}

// Get Info by token
// @router /:token [get]
func (c *InfoController) GetOne() {
	token := c.Ctx.Input.Param(":token")
	var resp utils.SimpleResponse
	v := services.QueryIdentity(token, true)
	if v != nil {
		resp.Success = true
		resp.Message = v
	} else {
		resp.Success = false
		resp.Message = utils.Msg.TokenErr
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// Change Password
// @router / [put]
func (c *InfoController) Put() {
	var input map[string]interface{}
	var resp utils.SimpleResponse
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		// read random and username from SESSION
		if c.GetSession("challenge") == nil {
			resp.Success = false
			resp.Message = utils.Msg.ContextErr
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		// check param
		if input["password"] == nil || input["newpassword"] == nil {
			resp.Success = false
			resp.Message = utils.Msg.ParamErr + "(password & newpassword required)"
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		challengeInfo := c.GetSession("challenge").(utils.ChallengeInfo)
		if challengeInfo.TimeStamp + 2 > utils.TimeStamp(0) { // time is too close
			resp.Success = false
			resp.Message = utils.Msg.ContextErr
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		random := challengeInfo.Random
		username := challengeInfo.Username
		password := input["password"].(string)
		newpassword := input["newpassword"].(string)

		resp.Success, resp.Message = services.CheckLogin(username, password, random)
		if resp.Success == true { // CheckLogin passed
			token := resp.Message.(string)
			v := services.QueryIdentity(token, false)
			resp.Success = services.ChangePassword(v.Id, newpassword)
			resp.Message = ""
		}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Check Code
// @router / [delete]
func (c *InfoController) Delete() {
	var input map[string]interface{}
	var resp utils.SimpleResponse
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		// check param
		if input["appname"] == nil || input["appsecret"] == nil || input["code"] == nil {
			resp.Success = false
			resp.Message = utils.Msg.ParamErr + "(appname & appsecret & code required)"
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		appname := input["appname"].(string)
		appsecret := input["appsecret"].(string)
		code := input["code"].(string)

		resp.Success, resp.Message = services.CheckCode(code, appname, appsecret)
		if resp.Success == true { // CheckLogin passed
			token := resp.Message.(string)
			v := services.QueryIdentity(token, false)
			resp.Message = v
		}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
