package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type HomeInsurances struct {
	Id                        int       `orm:"column(id);auto"`
	Address                   string    `orm:"column(address);size(255)"`
	City                      string    `orm:"column(city);size(45);null"`
	Postcode                  int       `orm:"column(postcode)"`
	ProposedStartDate         time.Time `orm:"column(proposed_start_date);type(date)"`
	EstimatedYearConstruction int       `orm:"column(estimated_year_construction)"`
	NumberStoreys             int       `orm:"column(number_storeys)"`
	NumberBathrooms           int       `orm:"column(number_bathrooms)"`
	NumberBedrooms            int       `orm:"column(number_bedrooms)"`
	BedroomSize               string    `orm:"column(bedroom_size);size(24)"`
	HomeDesc                  string    `orm:"column(home_desc);size(255)"`
	ExteriorWalls             string    `orm:"column(exterior_walls);size(32)"`
	RoofMaterial              string    `orm:"column(roof_material);size(32)"`
	HomeOccupied              string    `orm:"column(home_occupied);size(32)"`
	StrataPlan                string    `orm:"column(strata_plan);size(12)"`
	Pool                      string    `orm:"column(pool);size(32)"`
	Name                      string    `orm:"column(name);size(100)"`
	DobOldestPolicyHolder     time.Time `orm:"column(dob_oldest_policy_holder);type(date)"`
	EntitledNoClaim           int8      `orm:"column(entitled_no_claim)"`
	AnyBelowGround            int8      `orm:"column(any_below_ground)"`
	SecurityDevices           string    `orm:"column(security_devices);size(255)"`
	DoorSecurityDevices       string    `orm:"column(door_security_devices);size(100)"`
	WindowSecurityDevices     string    `orm:"column(window_security_devices);size(100)"`
	Alarm                     int8      `orm:"column(alarm)"`
	LandLargerNormal          int8      `orm:"column(land_larger_normal)"`
	HadPolicyCancelled        int8      `orm:"column(had_policy_cancelled)"`
	NumTimesClaimDeclined     int       `orm:"column(num_times_claim_declined)"`
	NumClaimsTotal            int       `orm:"column(num_claims_total)"`
	NumConvictions            int       `orm:"column(num_convictions)"`
	CoverStormDmgGatesFences  int8      `orm:"column(cover_storm_dmg_gates_fences)"`
	CoverAccidentalDmg        int8      `orm:"column(cover_accidental_dmg)"`
	Over50Discount            int8      `orm:"column(over_50_discount)"`
	Carports                  int       `orm:"column(carports)"`
	BalconiesDecks            int       `orm:"column(balconies_decks)"`
	Verandahs                 int       `orm:"column(verandahs)"`
	EuroKitchenAppliances     int8      `orm:"column(euro_kitchen_appliances)"`
	GraniteMarbleTiling       int8      `orm:"column(granite_marble_tiling)"`
	LargeGlazedAreas          int8      `orm:"column(large_glazed_areas)"`
	PlantationShutters        int8      `orm:"column(plantation_shutters)"`
	CurvedWalls               int8      `orm:"column(curved_walls)"`
	DuctedAircon              int8      `orm:"column(ducted_aircon)"`
	TennisCourt               int8      `orm:"column(tennis_court)"`
	IngroundPool              int8      `orm:"column(inground_pool)"`
	Watertight                int8      `orm:"column(watertight)"`
	NewHomeUnderConstruction  int8      `orm:"column(new_home_under_construction)"`
	UnderRenovation           int8      `orm:"column(under_renovation)"`
	HomeUse                   string    `orm:"column(home_use);size(100)"`
	Mortgate                  int8      `orm:"column(mortgate)"`
	InsureNameOrCompany       string    `orm:"column(insure_name_or_company);size(50)"`
	NumOwnersNamedPolicy      int       `orm:"column(num_owners_named_policy)"`
	Comments                  string    `orm:"column(comments);size(255)"`
	Status                    int       `orm:"column(status)"`
	Open                      int       `orm:"column(open)"`
	UserId                    int       `orm:"column(user_id)"`
	CreatedAt                 time.Time `orm:"column(created_at);type(timestamp)"`
	UpdatedAt                 time.Time `orm:"column(updated_at);type(timestamp)"`
}

func (t *HomeInsurances) TableName() string {
	return "home_insurances"
}

func init() {
	//orm.RegisterModel(new(HomeInsurances))
}

// AddHomeInsurances insert a new HomeInsurances into database and returns
// last inserted Id on success.
func AddHomeInsurances(m *HomeInsurances) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHomeInsurancesById retrieves HomeInsurances by Id. Returns error if
// Id doesn't exist
func GetHomeInsurancesById(id int) (v *HomeInsurances, err error) {
	o := orm.NewOrm()
	v = &HomeInsurances{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHomeInsurances retrieves all HomeInsurances matches certain condition. Returns empty list if
// no records exist
func GetAllHomeInsurances(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(HomeInsurances))
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

	var l []HomeInsurances
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

// UpdateHomeInsurances updates HomeInsurances by Id and returns error if
// the record to be updated doesn't exist
func UpdateHomeInsurancesById(m *HomeInsurances) (err error) {
	o := orm.NewOrm()
	v := HomeInsurances{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHomeInsurances deletes HomeInsurances by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHomeInsurances(id int) (err error) {
	o := orm.NewOrm()
	v := HomeInsurances{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&HomeInsurances{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
