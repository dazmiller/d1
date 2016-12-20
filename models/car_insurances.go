package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type CarInsurances struct {
	Id                        int       `orm:"column(id);auto"`
	Make                      string    `orm:"column(make);size(32)"`
	Model                     string    `orm:"column(model);size(100)"`
	Year                      int       `orm:"column(year)"`
	BodyType                  string    `orm:"column(body_type);size(32)"`
	Transmission              string    `orm:"column(transmission);size(16)"`
	Colour                    int       `orm:"column(colour)"`
	PlanStartInsurance        int       `orm:"column(plan_start_insurance)"`
	Dob                       time.Time `orm:"column(dob);type(date)"`
	Gender                    string    `orm:"column(gender);size(32)"`
	AgeDriversLicenceObtained int       `orm:"column(age_drivers_licence_obtained)"`
	DriverOwner               int8      `orm:"column(driver_owner)"`
	OwnAnotherVehicle         int8      `orm:"column(own_another_vehicle)"`
	AnyAccidents              int8      `orm:"column(any_accidents)"`
	AccidentYear              int       `orm:"column(accident_year)"`
	ClaimType                 string    `orm:"column(claim_type);size(32)"`
	EngineSizeMin             int       `orm:"column(engine_size_min)"`
	EngineSizeMax             int       `orm:"column(engine_size_max)"`
	Cylinders                 int       `orm:"column(cylinders)"`
	EngineType                string    `orm:"column(engine_type);size(32)"`
	InductionTurbo            string    `orm:"column(induction_turbo);size(50)"`
	PowerMin                  int       `orm:"column(power_min)"`
	PowerMax                  int       `orm:"column(power_max)"`
	TowBrakedMin              int       `orm:"column(tow_braked_min)"`
	TowBrakedMax              int       `orm:"column(tow_braked_max)"`
	AncapSafetyRating         string    `orm:"column(ancap_safety_rating);size(32)"`
	GreenStarRating           string    `orm:"column(green_star_rating);size(32)"`
	PPlateApproved            string    `orm:"column(p_plate_approved);size(16)"`
	DealerPostcodes           string    `orm:"column(dealer_postcodes);size(100)"`
	DealerDistance            int       `orm:"column(dealer_distance)"`
	Status                    int       `orm:"column(status)"`
	Open                      int       `orm:"column(open)"`
	UserId                    int       `orm:"column(user_id)"`
	CreatedAt                 time.Time `orm:"column(created_at);type(timestamp)"`
	UpdatedAt                 time.Time `orm:"column(updated_at);type(timestamp)"`
}

func (t *CarInsurances) TableName() string {
	return "car_insurances"
}

func init() {
	//orm.RegisterModel(new(CarInsurances))
}

// AddCarInsurances insert a new CarInsurances into database and returns
// last inserted Id on success.
func AddCarInsurances(m *CarInsurances) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCarInsurancesById retrieves CarInsurances by Id. Returns error if
// Id doesn't exist
func GetCarInsurancesById(id int) (v *CarInsurances, err error) {
	o := orm.NewOrm()
	v = &CarInsurances{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCarInsurances retrieves all CarInsurances matches certain condition. Returns empty list if
// no records exist
func GetAllCarInsurances(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CarInsurances))
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

	var l []CarInsurances
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateCarInsurances updates CarInsurances by Id and returns error if
// the record to be updated doesn't exist
func UpdateCarInsurancesById(m *CarInsurances) (err error) {
	o := orm.NewOrm()
	v := CarInsurances{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCarInsurances deletes CarInsurances by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCarInsurances(id int) (err error) {
	o := orm.NewOrm()
	v := CarInsurances{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CarInsurances{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
