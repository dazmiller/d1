package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type FuneralInsuranceFrm struct {
	Id              int       `form:"id"`
	Gender          string    `form:"gender"`
	Age             int       `form:"age"`
	Smoker          int       `form:"smoker"`
	Height          int       `form:"height"`
	Weight          int       `form:"weight"`
	CoverAmount     int       `form:"cover_amount"`
	ResidentCountry string    `form:"resident_country"`
	ResidencyStatus string    `form:"residency_status"`
	Dob             time.Time `form:"dob"`
	Created         time.Time `form:"created"`
	Updated         time.Time `form:"updated"`
}

type FuneralInsurance struct {
	Id              int       `orm:"column(id);pk";form:"id"`
	Gender          string    `orm:"column(gender);size(8);null";form:"gender"`
	Age             int       `orm:"column(age);null";form:"age"`
	Smoker          int       `orm:"column(smoker);null"`
	Height          int       `orm:"column(height);null"`
	Weight          int       `orm:"column(weight);null"`
	CoverAmount     int       `orm:"column(cover_amount);null"`
	ResidentCountry string    `orm:"column(resident_country);size(45);null"`
	ResidencyStatus string    `orm:"column(residency_status);size(45);null"`
	Dob             time.Time `orm:"column(dob);type(date);null"`
	Created         time.Time `orm:"column(created);type(datetime);null"`
	Updated         time.Time `orm:"column(updated);type(datetime);null"`
}

func (t *FuneralInsurance) TableName() string {
	return "funeral_insurance"
}

func init() {
	//orm.RegisterModel(new(FuneralInsurance))
}

// AddFuneralInsurance insert a new FuneralInsurance into database and returns
// last inserted Id on success.
func AddFuneralInsurance(m *FuneralInsurance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFuneralInsuranceById retrieves FuneralInsurance by Id. Returns error if
// Id doesn't exist
func GetFuneralInsuranceById(id int) (v *FuneralInsurance, err error) {
	o := orm.NewOrm()
	v = &FuneralInsurance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFuneralInsurance retrieves all FuneralInsurance matches certain condition. Returns empty list if
// no records exist
func GetAllFuneralInsurance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FuneralInsurance))
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

	var l []FuneralInsurance
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

// UpdateFuneralInsurance updates FuneralInsurance by Id and returns error if
// the record to be updated doesn't exist
func UpdateFuneralInsuranceById(m *FuneralInsurance) (err error) {
	o := orm.NewOrm()
	v := FuneralInsurance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFuneralInsurance deletes FuneralInsurance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFuneralInsurance(id int) (err error) {
	o := orm.NewOrm()
	v := FuneralInsurance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FuneralInsurance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
