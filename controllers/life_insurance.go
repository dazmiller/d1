package controllers

import (
	"encoding/json"
	"errors"
	"nohasslsapi/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for LifeInsurance
type LifeController struct {
	beego.Controller
}

func (c *LifeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create LifeInsurance
// @Param	body		body 	models.LifeInsurance	true		"body for LifeInsurance content"
// @Success 201 {int} models.LifeInsurance
// @Failure 403 body is empty
// @router / [post]
func (c *LifeController) Post() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	var v models.LifeInsurance
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddLifeInsurance(&v); err == nil {
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
// @Description get LifeInsurance by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LifeInsurance
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LifeController) GetOne() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//jsoninfo := c.GetString("id")
	//id, _ := strconv.Atoi(jsoninfo)
	v, err := models.GetLifeInsuranceById(1)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {

		rec := []*models.LifeInsurance{}
		rec = append(rec, v)
		response := map[string]interface{}{

			"data": rec,
		}

		c.Data["json"] = response
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get LifeInsurance
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Param	since   query   datetime false  "Datetime to return records after. Checks creation "
// @Success 200 {object} models.LifeInsurance
// @Failure 403
// @router / [get]
func (c *LifeController) GetAll() {

	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	var fields []string
	var sortby []string
	//var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := c.GetString("f"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("l"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("o"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("s"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	//if v := c.GetString("order"); v != "" {
	//	order = strings.Split(v, ",")
	//}
	// query: k:v,k:v
	if v := c.GetString("q"); v != "" {
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

	l, err := models.GetAllLifeInsurance(query, fields, sortby, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		//places := []*models.LifeInsurance{}
		//places = l
		response := map[string]interface{}{
			"data": l,
		}

		c.Data["json"] = response

		//c.Data["json"] = l
	}
	c.ServeJSON()

}

// @Title Update
// @Description update the LifeInsurance
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.LifeInsurance	true		"body for LifeInsurance content"
// @Success 200 {object} models.LifeInsurance
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LifeController) Put() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.LifeInsurance{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateLifeInsuranceById(&v); err == nil {
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
// @Description delete the LifeInsurance
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LifeController) Delete() {
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteLifeInsurance(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
