package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type QuoteContactRequestFrm struct {
	Id                      int       `form:"id"`
	UserId                  int       `form:"user_id"`
	QuoteType               string    `form:"quote_type"`
	QuoteId                 int       `form:"quote_id"`
	ProviderId              int       `form:"provider_id"`
	ClientRequestDateTime   time.Time `form:"client_request_date_time"`
	ProviderConfirmDateTime time.Time `form:"provider_confirm_date_time"`
	CreatedAt               time.Time `form:"created_at"`
	LastUpdated             time.Time `form:"last_updated"`
	Open                    int       `form:"open"`
}

type QuoteContactRequest struct {
	Id                      int       `orm:"column(id);pk"`
	UserId                  int       `orm:"column(user_id);null"`
	QuoteType               string    `orm:"column(quote_type);size(45);null"`
	QuoteId                 int       `orm:"column(quote_id);null"`
	ProviderId              int       `orm:"column(provider_id);null"`
	ClientRequestDateTime   time.Time `orm:"column(client_request_date_time);type(datetime);null"`
	ProviderConfirmDateTime time.Time `orm:"column(provider_confirm_date_time);type(datetime);null"`
	CreatedAt               time.Time `orm:"column(created_at);type(datetime);null"`
	LastUpdated             time.Time `orm:"column(last_updated);type(datetime);null"`
	Open                    int       `orm:"column(open);null"`
}

func (t *QuoteContactRequest) TableName() string {
	return "quote_contact_request"
}

func init() {
	//orm.RegisterModel(new(QuoteContactRequest))
}

// AddQuoteContactRequest insert a new QuoteContactRequest into database and returns
// last inserted Id on success.
func AddQuoteContactRequest(m *QuoteContactRequest) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetQuoteContactRequestById retrieves QuoteContactRequest by Id. Returns error if
// Id doesn't exist
func GetQuoteContactRequestById(id int) (v *QuoteContactRequest, err error) {
	o := orm.NewOrm()
	v = &QuoteContactRequest{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllQuoteContactRequest retrieves all QuoteContactRequest matches certain condition. Returns empty list if
// no records exist
func GetAllQuoteContactRequest(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(QuoteContactRequest))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []QuoteContactRequest
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateQuoteContactRequest updates QuoteContactRequest by Id and returns error if
// the record to be updated doesn't exist
func UpdateQuoteContactRequestById(m *QuoteContactRequest) (err error) {
	o := orm.NewOrm()
	v := QuoteContactRequest{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteQuoteContactRequest deletes QuoteContactRequest by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQuoteContactRequest(id int) (err error) {
	o := orm.NewOrm()
	v := QuoteContactRequest{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&QuoteContactRequest{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
