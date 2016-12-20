package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Provider struct {
	ID             int       `orm:"column(ID);auto"`
	CompanyName    string    `orm:"column(CompanyName);size(45);null"`
	CompanyCode    string    `orm:"column(CompanyCode);size(45);null"`
	Phone1         string    `orm:"column(Phone1);size(45);null"`
	Phone2         string    `orm:"column(Phone2);size(45);null"`
	Email          string    `orm:"column(Email);size(45);null"`
	Fax            string    `orm:"column(Fax);size(45);null"`
	Website        string    `orm:"column(Website);size(45);null"`
	Information    string    `orm:"column(Information);size(255);null"`
	SectorCoverage string    `orm:"column(SectorCoverage);size(255);null"`
	Active         int       `orm:"column(active);null"`
	Country        string    `orm:"column(Country);size(45);null"`
	MailAddress    string    `orm:"column(MailAddress);size(255);null"`
	OperationHours string    `orm:"column(OperationHours);size(255);null"`
	OfficeAddress  string    `orm:"column(OfficeAddress);size(255);null"`
	Phone3         string    `orm:"column(Phone3);size(45);null"`
	MemberSince    time.Time `orm:"column(MemberSince);type(date);null"`
}

func (t *Provider) TableName() string {
	return "provider"
}

func init() {
	//orm.RegisterModel(new(Provider))
}

// AddProvider insert a new Provider into database and returns
// last inserted Id on success.
func AddProvider(m *Provider) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProviderById retrieves Provider by Id. Returns error if
// Id doesn't exist
func GetProviderById(id int) (v *Provider, err error) {
	o := orm.NewOrm()
	v = &Provider{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProvider retrieves all Provider matches certain condition. Returns empty list if
// no records exist
func GetAllProvider(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Provider))
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

	var l []Provider
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

// UpdateProvider updates Provider by Id and returns error if
// the record to be updated doesn't exist
func UpdateProviderById(m *Provider) (err error) {

	o := orm.NewOrm()
	v := Provider{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProvider deletes Provider by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProvider(id int) (err error) {
	o := orm.NewOrm()
	v := Provider{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Provider{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
