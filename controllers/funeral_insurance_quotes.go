package controllers

import (
	"encoding/json"
	"errors"
	"nohasslsapi/models"
	"nohasslsapi/utils"
	"strconv"
	"strings"

	"github.com/dazmiller/beego"
)

// oprations for FuneralInsuranceQuotes
type FuneralQuotesController struct {
	beego.Controller
}

func (c *FuneralQuotesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create FuneralInsuranceQuotes
// @Param	body		body 	models.FuneralInsuranceQuotes	true		"body for FuneralInsuranceQuotes content"
// @Success 201 {int} models.FuneralInsuranceQuotes
// @Failure 403 body is empty
// @router / [post]
func (c *FuneralQuotesController) Post() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	var v models.FuneralInsuranceQuotes
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddFuneralInsuranceQuotes(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get FuneralInsuranceQuotes by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.FuneralInsuranceQuotes
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FuneralQuotesController) GetOne() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	jsoninfo := c.GetString("id")
	id, _ := strconv.Atoi(jsoninfo)
	v, err := models.GetFuneralInsuranceQuotesById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get FuneralInsuranceQuotes
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.FuneralInsuranceQuotes
// @Failure 403
// @router / [get]
func (c *FuneralQuotesController) GetAll() {

	version := utils.GetVersion(c.Ctx.Input.Header("Accept"))
	version = "v1"
	switch version {
	case "v1":

		c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
		var fields []string
		var sortby []string
		var order []string
		var query map[string]string = make(map[string]string)
		var limit int64 = 10
		var offset int64 = 0

		// fields: col1,col2,entity.col3
		if v := c.GetString("fields"); v != "" {
			fields = strings.Split(v, ",")
		}
		// limit: 10 (default is 10)
		if v, err := c.GetInt64("limit"); err == nil {
			limit = v
		}
		// offset: 0 (default is 0)
		if v, err := c.GetInt64("offset"); err == nil {
			offset = v
		}
		// sortby: col1,col2
		if v := c.GetString("sortby"); v != "" {
			sortby = strings.Split(v, ",")
		}
		// order: desc,asc
		if v := c.GetString("order"); v != "" {
			order = strings.Split(v, ",")
		}
		// query: k:v,k:v
		if v := c.GetString("query"); v != "" {
			for _, cond := range strings.Split(v, ",") {
				kv := strings.Split(cond, ":")
				if len(kv) != 2 {
					c.Data["json"] = errors.New("Error: invalid query key/value pair")
					c.ServeJSON()
					return
				}
				k, v := kv[0], kv[1]
				query[k] = v
			}
		}

		l, err := models.GetAllFuneralInsuranceQuotes(query, fields, sortby, order, offset, limit)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = l
		}
		c.ServeJSON()

		// no version supplied
	default:
		c.Ctx.Output.SetStatus(505)
		c.Data["json"] = "Incorrect or Missing Version Number Supplied: You Supplied : " + version
		c.ServeJSON()
		return
	}

}

// @Title Update
// @Description update the FuneralInsuranceQuotes
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.FuneralInsuranceQuotes	true		"body for FuneralInsuranceQuotes content"
// @Success 200 {object} models.FuneralInsuranceQuotes
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FuneralQuotesController) Put() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.FuneralInsuranceQuotes{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateFuneralInsuranceQuotesById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the FuneralInsuranceQuotes
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FuneralQuotesController) Delete() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteFuneralInsuranceQuotes(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
