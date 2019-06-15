package controllers

import (
	"xyzua/services"
	"xyzua/utils"

	"encoding/json"

	"github.com/astaxie/beego"
)

// UtilController operations for info
type UtilController struct {
	beego.Controller
}

// URLMapping ...
func (c *UtilController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Mark
// @router / [post]
func (c *UtilController) Post() {
	var input map[string]interface{}
	var resp utils.SimpleResponse
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &input); err == nil {
		if input["t"] == nil || input["type"] == nil {
			resp.Success = false
			resp.Message = utils.Msg.ParamErr + "(token & type required)"
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
		t := input["t"].(string)
		_type := input["type"].(string)
		// Check admintoken
		if _type != "SELF" {
			if input["admintoken"] == nil {
				resp.Success = false
				resp.Message = utils.Msg.ParamErr + "(admintoken required)"
				c.Data["json"] = resp
				c.ServeJSON()
				return
			}
			admintoken := input["admintoken"].(string)
			admin := services.QueryIdentity(admintoken, false)
			if admin.Role != "ADMIN" {
				resp.Success = false
				resp.Message = utils.Msg.Forbidden
				c.ServeJSON()
				return
			}
		}
		// Mark
		u := services.QueryIdentity(t, false)
		if u != nil {
			resp.Success = services.Mark(u.Identity, _type)
		} else {
			resp.Success = false
			resp.Message = utils.Msg.TokenErr
		}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Get all App Links
// @router / [get]
func (c *UtilController) GetAll() {
	apps := services.QueryApps()
	c.Data["json"] = apps
	c.ServeJSON()
}

// Get Mark by token
// @router /:t [get]
func (c *UtilController) GetOne() {
	t := c.Ctx.Input.Param(":t")
	var resp utils.SimpleResponse
	v := services.QueryIdentity(t, false)
	if v != nil {
		resp.Success = true
		marks := services.QueryMarks(v.Identity)
		resp.Message = marks
	} else {
		resp.Success = false
		resp.Message = utils.Msg.TokenErr
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// nothing
// @router / [put]
func (c *UtilController) Put() {

}

// nothing
// @router / [delete]
func (c *UtilController) Delete() {

}
