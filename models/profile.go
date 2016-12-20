package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Profile struct {
	Id               int       `orm:"column(id);auto"`
	UserId           int       `orm:"column(user_id);null"`
	Firstname        string    `orm:"column(firstname);size(45);null"`
	Middlename       string    `orm:"column(middlename);size(45);null"`
	Surname          string    `orm:"column(surname);size(45);null"`
	Sex              string    `orm:"column(sex);size(45);null"`
	Birthdate        time.Time `orm:"column(birthdate);type(date);null"`
	Height           int       `orm:"column(height);null"`
	Weight           int       `orm:"column(weight);null"`
	Smoker           uint64    `orm:"column(smoker);size(1);null"`
	Avdrinks         int       `orm:"column(avdrinks);null"`
	Address1         string    `orm:"column(address1);size(90);null"`
	Address2         string    `orm:"column(address2);size(90);null"`
	Address3         string    `orm:"column(address3);size(90);null"`
	Suburb           string    `orm:"column(suburb);size(45);null"`
	State            string    `orm:"column(state);size(10);null"`
	Postcode         int       `orm:"column(postcode);null"`
	Email1           string    `orm:"column(email1);size(90);null"`
	Email2           string    `orm:"column(email2);size(90);null"`
	PhoneHome        string    `orm:"column(phone_home);size(45);null"`
	PhoneMob         string    `orm:"column(phone_mob);size(45);null"`
	Msn              string    `orm:"column(msn);size(90);null"`
	Twitter          string    `orm:"column(twitter);size(90);null"`
	Facebook         string    `orm:"column(facebook);size(90);null"`
	Wechat           string    `orm:"column(wechat);size(90);null"`
	Profession       string    `orm:"column(profession);size(90);null"`
	Profilecol       string    `orm:"column(profilecol);size(45);null"`
	Title            string    `orm:"column(title);size(45);null"`
	Child1Name       string    `orm:"column(child1_name);size(45);null"`
	Child1Birthdate  time.Time `orm:"column(child1_birthdate);type(date);null"`
	Child1Sex        string    `orm:"column(child1_sex);size(10);null"`
	Child2Name       string    `orm:"column(child2_name);size(45);null"`
	Child2Birthdate  time.Time `orm:"column(child2_birthdate);type(date);null"`
	Child2Sex        string    `orm:"column(child2_sex);size(10);null"`
	Child3Name       string    `orm:"column(child3_name);size(45);null"`
	Child3Birthdate  time.Time `orm:"column(child3_birthdate);type(date);null"`
	Child3Sex        string    `orm:"column(child3_sex);size(10);null"`
	Child4Name       string    `orm:"column(child4_name);size(45);null"`
	Child4Birthdatae time.Time `orm:"column(child4_birthdatae);type(date);null"`
	Child4Sex        string    `orm:"column(child4_sex);size(10);null"`
	Child5Name       string    `orm:"column(child5_name);size(45);null"`
	Child5Birthdate  time.Time `orm:"column(child5_birthdate);type(date);null"`
	Child5Sex        string    `orm:"column(child5_sex);size(10);null"`
	NoChildren       int       `orm:"column(no_children);null"`
	NoPets           int       `orm:"column(no_pets);null"`
	Pet1Name         string    `orm:"column(pet1_name);size(45);null"`
	Pet1Type         string    `orm:"column(pet1_type);size(45);null"`
	Pet1Sex          string    `orm:"column(pet1_sex);size(10);null"`
	Pet2Name         string    `orm:"column(pet2_name);size(45);null"`
	Pet2Type         string    `orm:"column(pet2_type);size(45);null"`
	Pet2Sex          string    `orm:"column(pet2_sex);size(10);null"`
	Pet3Name         string    `orm:"column(pet3_name);size(45);null"`
	Pet3Sex          string    `orm:"column(pet3_sex);size(10);null"`
	Pet3Type         string    `orm:"column(pet3_type);size(45);null"`
	LastUpdated      time.Time `orm:"column(last_updated);type(datetime);null"`
	OptionsRadios    string    `orm:"column(options_radios);size(191);null"`
	A                string    `orm:"column(a);size(191);null"`
}

func (t *Profile) TableName() string {
	return "profile"
}

func init() {
	//orm.RegisterModel(new(Profile))
}

// AddProfile insert a new Profile into database and returns
// last inserted Id on success.
func AddProfile(m *Profile) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProfileById retrieves Profile by Id. Returns error if
// Id doesn't exist
func GetProfileById(id int) (v *Profile, err error) {
	o := orm.NewOrm()
	v = &Profile{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProfile retrieves all Profile matches certain condition. Returns empty list if
// no records exist
func GetAllProfile(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Profile))
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

	var l []Profile
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

// UpdateProfile updates Profile by Id and returns error if
// the record to be updated doesn't exist
func UpdateProfileById(m *Profile) (err error) {
	o := orm.NewOrm()
	v := Profile{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProfile deletes Profile by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProfile(id int) (err error) {
	o := orm.NewOrm()
	v := Profile{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Profile{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
