package controllers

import (
	"encoding/json"

	"xyzua/services"
	"xyzua/utils"

	"github.com/astaxie/beego"
)

// IdentityController operations for identity
type IdentityController struct {
	beego.Controller
}

// URLMapping ...
func (c *IdentityController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// RESPONSE of HMAC login
// @router / [post]
func (c *IdentityController) Post() {
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
		if input["password"] == nil {
			resp.Success = false
			resp.Message = utils.Msg.ParamErr + "(password required)"
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		challengeInfo := c.GetSession("challenge").(utils.ChallengeInfo)
		if challengeInfo.TimeStamp+2 > utils.TimeStamp(0) { // time is too close
			resp.Success = false
			resp.Message = utils.Msg.ContextErr
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		random := challengeInfo.Random
		username := challengeInfo.Username
		password := input["password"].(string)

		resp.Success, resp.Message = services.CheckLogin(username, password, random)
		c.Data["json"] = resp
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Testing
// @router / [get]
func (c *IdentityController) GetAll() {
	c.Ctx.WriteString("IdentityController - XYZUA")
}

// CHALLENGE of HMAC login
// @router /:username [get]
func (c *IdentityController) GetOne() {
	var challengeInfo utils.ChallengeInfo
	challengeInfo.Random = utils.RandomString(16)
	challengeInfo.Username = c.Ctx.Input.Param(":username")
	challengeInfo.TimeStamp = utils.TimeStamp(0)
	c.SetSession("challenge", challengeInfo)
	c.Ctx.WriteString(challengeInfo.Random)
}

// Exchange Token
// @router /:token [put]
func (c *IdentityController) Put() {
	token := c.Ctx.Input.Param(":token")
	var resp utils.SimpleResponse
	resp.Success, resp.Message = services.CheckToken(token)
	c.Data["json"] = resp
	c.ServeJSON()
}

// Logout
// @router /:token [delete]
func (c *IdentityController) Delete() {
	token := c.Ctx.Input.Param(":token")
	var resp utils.SimpleResponse
	resp.Success, resp.Message = services.CheckToken(token)
	if resp.Success == true {
		resp.Message = "" // hide new token
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
