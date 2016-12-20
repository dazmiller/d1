package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MortgageInsuranceFrm struct {
	Id                                  int       `form:"id"`
	BatchId                             int       `form:"batch_id"`
	Gender                              string    `form:"gender"`
	Smoker                              int       `form:"smoker"`
	Age                                 int       `form:"age"`
	MortgageRecent                      int       `form:"mortgage_recent"`
	MortgageAmount                      int       `form:"mortgage_amount"`
	LenderName                          string    `form:"lender_name"`
	LenderBranch                        string    `form:"lender_branch"`
	LenderCity                          string    `form:"lender_city"`
	CountryResident                     string    `form:"country_resident"`
	ResidentStatus                      string    `form:"resident_status"`
	Dob                                 time.Time `form:"dob"`
	Height                              int       `form:"height"`
	Weight                              int       `form:"weight"`
	CancerStatus                        int       `form:"cancer_status"`
	CancerTypes                         string    `form:"cancer_types"`
	DiabetesStatus                      int       `form:"diabetes_status"`
	BloodDisorderStatus                 int       `form:"blood_disorder_status"`
	DiabetiesComplications              int       `form:"diabeties_complications"`
	DiabetesComplicationsDetails        string    `form:"diabetes_complications_details"`
	HighBloodPressure                   int       `form:"high_blood_pressure"`
	HighCholesterol                     int       `form:"high_cholesterol"`
	HeartProblems                       int       `form:"heart_problems"`
	HighBloodPressureTreatment          string    `form:"high_blood_pressure_treatment"`
	HighBloodPressureHeartKidneyProblem int       `form:"high_blood_pressure_heart_kidney_problem"`
	HighCholesterolTreatment            string    `form:"high_cholesterol_treatment"`
	CholesterolChecked12m               int       `form:"cholesterol_checked_12m"`
	LastCholesterolCheckStatus          string    `form:"last_cholesterol_check_status"`
	LastCholesterolCheckReading         string    `form:"last_cholesterol_check_reading"`
	HeartProblemsDetails                string    `form:"heart_problems_details"`
	GastroProblemsStatus                string    `form:"gastro_problems_status"`
	KidneyBladderProblemsStatus         string    `form:"kidney_bladder_problems_status"`
	BreathingLungProblemStatus          string    `form:"breathing_lung_problem_status"`
	GastroProblemsDetails               string    `form:"gastro_problems_details"`
	HerniaRepaired                      int       `form:"hernia_repaired"`
	NumGastritis12m                     string    `form:"num_gastritis_12m"`
	GallstonesRemoved                   int       `form:"gallstones_removed"`
	GallstonesOngoingIssues             int       `form:"gallstones_ongoing_issues"`
	UlcerHospitlisation2d               int       `form:"ulcer_hospitlisation_2d"`
	OtherGastroProblems12m              int       `form:"other_gastro_problems_12m"`
	NumGastroIssues5y                   int       `form:"num_gastro_issues_5y"`
	KidneyProblemDetails                string    `form:"kidney_problem_details"`
	BreathingProblemsDetails            string    `form:"breathing_problems_details"`
	RiskyOccupation                     string    `form:"risky_occupation"`
	RecreationalActivities              string    `form:"recreational_activities"`
	TypeOfCancerWhenDiagnosed           string    `form:"type_of_cancer_when_diagnosed"`
	CancerMedicatedTreated              string    `form:"cancer_medicated_treated"`
	CancerStateRemission                string    `form:"cancer_state_remission"`
	DiabetesFirstDiagnosed              string    `form:"diabetes_first_diagnosed"`
	DiabetesMedicationControl           string    `form:"diabetes_medication_control"`
	DiabetesTests12m                    string    `form:"diabetes_tests_12m"`
	BloodDisorderDiagnosed              string    `form:"blood_disorder_diagnosed"`
	BloodDisorderMedication             string    `form:"blood_disorder_medication"`
	BloodDisorderTests12m               string    `form:"blood_disorder_tests_12m"`
	BloodPressureHighDiagnosed          string    `form:"blood_pressure_high_diagnosed"`
	BloodPressureMedication             string    `form:"blood_pressure_medication"`
	BloodPressureTest                   string    `form:"blood_pressure_test"`
	CholesterolDiagnosed                string    `form:"cholesterol_diagnosed"`
	CholestrolMedication                string    `form:"cholestrol_medication"`
	CholesterolTest                     string    `form:"cholesterol_test"`
	HeartConditionDiagnosed             string    `form:"heart_condition_diagnosed"`
	HeartConditionMedication            string    `form:"heart_condition_medication"`
	HeartConditionTest12m               string    `form:"heart_condition_test_12m"`
	GastroIntestinalDiagnosed           string    `form:"gastro_intestinal_diagnosed"`
	GastroIntestinalMedication          string    `form:"gastro_intestinal_medication"`
	GastroIntestinalCurrentState        string    `form:"gastro_intestinal_current_state"`
	KidneyBladderDiagnosed              string    `form:"kidney_bladder_diagnosed"`
	KidneyBladderMedication             string    `form:"kidney_bladder_medication"`
	KidneyBladderCurrentState           string    `form:"kidney_bladder_current_state"`
	LungProblemDiagnosed                string    `form:"lung_problem_diagnosed"`
	LungProblemMedication               string    `form:"lung_problem_medication"`
	LungProblemCurrentState             string    `form:"lung_problem_current_state"`
	NeurologicalDiagnosed               string    `form:"neurological_diagnosed"`
	NeurologicalMedication              string    `form:"neurological_medication"`
	NeurologicalCurrentState            string    `form:"neurological_current_state"`
	MuscularSkeletalDiagnosed           string    `form:"muscular_skeletal_diagnosed"`
	MuscularSkeletalMedication          string    `form:"muscular_skeletal_medication"`
	MuscularSkeletalCurrentStatus       string    `form:"muscular_skeletal_current_status"`
	MentalHealthDiagnosed               string    `form:"mental_health_diagnosed"`
	MentalHealthCausedEventIllness      string    `form:"mental_health_caused_event_illness"`
	MentalHealthMedication              string    `form:"mental_health_medication"`
	AlcoholTypicalUse                   string    `form:"alcohol_typical_use"`
	AlcoholConsultedProfessional        string    `form:"alcohol_consulted_professional"`
	AlcoholAlcoholismOrCounselling      string    `form:"alcohol_alcoholism_or_counselling"`
	IllegalDrugsUse                     string    `form:"illegal_drugs_use"`
	IllegalDrugsConsultedProfessional   string    `form:"illegal_drugs_consulted_professional"`
	IllegalDrugsHaveStoppedPeriod       string    `form:"illegal_drugs_have_stopped_period"`
	CurrentMedicalAdviceSymptoms        string    `form:"current_medical_advice_symptoms"`
	CurrentMedicalAdviceProfession      string    `form:"current_medical_advice_profession"`
	CurrentMedicalAdviceTests           string    `form:"current_medical_advice_tests"`
	FamilyMemberIllnessDetails          string    `form:"family_member_illness_details"`
	Family2OrMoreHeartDisease           int       `form:"family_2_or_more_heart_disease"`
	Family2OrMoreStroke                 int       `form:"family_2_or_more_stroke"`
	KidneyDiseasePolycystic             int       `form:"kidney_disease_polycystic"`
	Family2OrMoreDiabetes               int       `form:"family_2_or_more_diabetes"`
	CancerTypeDiagnosed                 string    `form:"cancer_type_diagnosed"`
	FamilyCountColonRectal60            int       `form:"family_count_colon_rectal_60"`
	ColonRectalBefore50                 int       `form:"colon_rectal_before_50"`
	FamilyCountProstateCancer60         int       `form:"family_count_prostate_cancer_60"`
	ProstateCancerBefore50              int       `form:"prostate_cancer_before_50"`
	Family2MoreSameCancer60             int       `form:"family_2_more_same_cancer_60"`
	CancerBefore50                      int       `form:"cancer_before_50"`
	FamilyCountTesticularCancer60       int       `form:"family_count_testicular_cancer_60"`
	TesticularCancerBefore50            int       `form:"testicular_cancer_before_50"`
	FamilyCountMs60                     int       `form:"family_count_ms_60"`
	FamilyCountParkinsons60             int       `form:"family_count_parkinsons_60"`
	FamilyCountMotorNeurone60           int       `form:"family_count_motor_neurone_60"`
}

type MortgageInsurance struct {
	Id                                  int       `orm:"column(id);auto"`
	BatchId                             int       `orm:"column(batch_id);null"`
	Gender                              string    `orm:"column(gender);size(8);null"`
	Smoker                              int       `orm:"column(smoker);null"`
	Age                                 int       `orm:"column(age);null"`
	MortgageRecent                      int       `orm:"column(mortgage_recent);null"`
	MortgageAmount                      int       `orm:"column(mortgage_amount);null"`
	LenderName                          string    `orm:"column(lender_name);size(45);null"`
	LenderBranch                        string    `orm:"column(lender_branch);size(45);null"`
	LenderCity                          string    `orm:"column(lender_city);size(45);null"`
	CountryResident                     string    `orm:"column(country_resident);size(45);null"`
	ResidentStatus                      string    `orm:"column(resident_status);size(45);null"`
	Dob                                 time.Time `orm:"column(dob);type(date);null"`
	Height                              int       `orm:"column(height);null"`
	Weight                              int       `orm:"column(weight);null"`
	CancerStatus                        int       `orm:"column(cancer_status);null"`
	CancerTypes                         string    `orm:"column(cancer_types);size(255);null"`
	DiabetesStatus                      int       `orm:"column(diabetes_status);null"`
	BloodDisorderStatus                 int       `orm:"column(blood_disorder_status);null"`
	DiabetiesComplications              int       `orm:"column(diabeties_complications);null"`
	DiabetesComplicationsDetails        string    `orm:"column(diabetes_complications_details);size(255);null"`
	HighBloodPressure                   int       `orm:"column(high_blood_pressure);null"`
	HighCholesterol                     int       `orm:"column(high_cholesterol);null"`
	HeartProblems                       int       `orm:"column(heart_problems);null"`
	HighBloodPressureTreatment          string    `orm:"column(high_blood_pressure_treatment);size(45);null"`
	HighBloodPressureHeartKidneyProblem int       `orm:"column(high_blood_pressure_heart_kidney_problem);null"`
	HighCholesterolTreatment            string    `orm:"column(high_cholesterol_treatment);size(255);null"`
	CholesterolChecked12m               int       `orm:"column(cholesterol_checked_12m);null"`
	LastCholesterolCheckStatus          string    `orm:"column(last_cholesterol_check_status);size(45);null"`
	LastCholesterolCheckReading         string    `orm:"column(last_cholesterol_check_reading);size(45);null"`
	HeartProblemsDetails                string    `orm:"column(heart_problems_details);size(255);null"`
	GastroProblemsStatus                string    `orm:"column(gastro_problems_status);size(255);null"`
	KidneyBladderProblemsStatus         string    `orm:"column(kidney_bladder_problems_status);size(255);null"`
	BreathingLungProblemStatus          string    `orm:"column(breathing_lung_problem_status);size(255);null"`
	GastroProblemsDetails               string    `orm:"column(gastro_problems_details);size(255);null"`
	HerniaRepaired                      int       `orm:"column(hernia_repaired);null"`
	NumGastritis12m                     string    `orm:"column(num_gastritis_12m);size(8);null"`
	GallstonesRemoved                   int       `orm:"column(gallstones_removed);null"`
	GallstonesOngoingIssues             int       `orm:"column(gallstones_ongoing_issues);null"`
	UlcerHospitlisation2d               int       `orm:"column(ulcer_hospitlisation_2d);null"`
	OtherGastroProblems12m              int       `orm:"column(other_gastro_problems_12m);null"`
	NumGastroIssues5y                   int       `orm:"column(num_gastro_issues_5y);null"`
	KidneyProblemDetails                string    `orm:"column(kidney_problem_details);size(255);null"`
	BreathingProblemsDetails            string    `orm:"column(breathing_problems_details);size(255);null"`
	RiskyOccupation                     string    `orm:"column(risky_occupation);size(255);null"`
	RecreationalActivities              string    `orm:"column(recreational_activities);size(255);null"`
	TypeOfCancerWhenDiagnosed           string    `orm:"column(type_of_cancer_when_diagnosed);size(255);null"`
	CancerMedicatedTreated              string    `orm:"column(cancer_medicated_treated);size(255);null"`
	CancerStateRemission                string    `orm:"column(cancer_state_remission);size(255);null"`
	DiabetesFirstDiagnosed              string    `orm:"column(diabetes_first_diagnosed);size(255);null"`
	DiabetesMedicationControl           string    `orm:"column(diabetes_medication_control);size(255);null"`
	DiabetesTests12m                    string    `orm:"column(diabetes_tests_12m);size(255);null"`
	BloodDisorderDiagnosed              string    `orm:"column(blood_disorder_diagnosed);size(255);null"`
	BloodDisorderMedication             string    `orm:"column(blood_disorder_medication);size(255);null"`
	BloodDisorderTests12m               string    `orm:"column(blood_disorder_tests_12m);size(255);null"`
	BloodPressureHighDiagnosed          string    `orm:"column(blood_pressure_high_diagnosed);size(255);null"`
	BloodPressureMedication             string    `orm:"column(blood_pressure_medication);size(255);null"`
	BloodPressureTest                   string    `orm:"column(blood_pressure_test);size(255);null"`
	CholesterolDiagnosed                string    `orm:"column(cholesterol_diagnosed);size(255);null"`
	CholestrolMedication                string    `orm:"column(cholestrol_medication);size(255);null"`
	CholesterolTest                     string    `orm:"column(cholesterol_test);size(255);null"`
	HeartConditionDiagnosed             string    `orm:"column(heart_condition_diagnosed);size(255);null"`
	HeartConditionMedication            string    `orm:"column(heart_condition_medication);size(255);null"`
	HeartConditionTest12m               string    `orm:"column(heart_condition_test_12m);size(255);null"`
	GastroIntestinalDiagnosed           string    `orm:"column(gastro_intestinal_diagnosed);size(255);null"`
	GastroIntestinalMedication          string    `orm:"column(gastro_intestinal_medication);size(255);null"`
	GastroIntestinalCurrentState        string    `orm:"column(gastro_intestinal_current_state);size(255);null"`
	KidneyBladderDiagnosed              string    `orm:"column(kidney_bladder_diagnosed);size(255);null"`
	KidneyBladderMedication             string    `orm:"column(kidney_bladder_medication);size(255);null"`
	KidneyBladderCurrentState           string    `orm:"column(kidney_bladder_current_state);size(255);null"`
	LungProblemDiagnosed                string    `orm:"column(lung_problem_diagnosed);size(255);null"`
	LungProblemMedication               string    `orm:"column(lung_problem_medication);size(255);null"`
	LungProblemCurrentState             string    `orm:"column(lung_problem_current_state);size(255);null"`
	NeurologicalDiagnosed               string    `orm:"column(neurological_diagnosed);size(255);null"`
	NeurologicalMedication              string    `orm:"column(neurological_medication);size(255);null"`
	NeurologicalCurrentState            string    `orm:"column(neurological_current_state);size(255);null"`
	MuscularSkeletalDiagnosed           string    `orm:"column(muscular_skeletal_diagnosed);size(255);null"`
	MuscularSkeletalMedication          string    `orm:"column(muscular_skeletal_medication);size(255);null"`
	MuscularSkeletalCurrentStatus       string    `orm:"column(muscular_skeletal_current_status);size(255);null"`
	MentalHealthDiagnosed               string    `orm:"column(mental_health_diagnosed);size(255);null"`
	MentalHealthCausedEventIllness      string    `orm:"column(mental_health_caused_event_illness);size(255);null"`
	MentalHealthMedication              string    `orm:"column(mental_health_medication);size(255);null"`
	AlcoholTypicalUse                   string    `orm:"column(alcohol_typical_use);size(255);null"`
	AlcoholConsultedProfessional        string    `orm:"column(alcohol_consulted_professional);size(255);null"`
	AlcoholAlcoholismOrCounselling      string    `orm:"column(alcohol_alcoholism_or_counselling);size(255);null"`
	IllegalDrugsUse                     string    `orm:"column(illegal_drugs_use);size(255);null"`
	IllegalDrugsConsultedProfessional   string    `orm:"column(illegal_drugs_consulted_professional);size(255);null"`
	IllegalDrugsHaveStoppedPeriod       string    `orm:"column(illegal_drugs_have_stopped_period);size(255);null"`
	CurrentMedicalAdviceSymptoms        string    `orm:"column(current_medical_advice_symptoms);size(255);null"`
	CurrentMedicalAdviceProfession      string    `orm:"column(current_medical_advice_profession);size(255);null"`
	CurrentMedicalAdviceTests           string    `orm:"column(current_medical_advice_tests);size(255);null"`
	FamilyMemberIllnessDetails          string    `orm:"column(family_member_illness_details);size(255);null"`
	Family2OrMoreHeartDisease           int       `orm:"column(family_2_or_more_heart_disease);null"`
	Family2OrMoreStroke                 int       `orm:"column(family_2_or_more_stroke);null"`
	KidneyDiseasePolycystic             int       `orm:"column(kidney_disease_polycystic);null"`
	Family2OrMoreDiabetes               int       `orm:"column(family_2_or_more_diabetes);null"`
	CancerTypeDiagnosed                 string    `orm:"column(cancer_type_diagnosed);size(255);null"`
	FamilyCountColonRectal60            int       `orm:"column(family_count_colon_rectal_60);null"`
	ColonRectalBefore50                 int       `orm:"column(colon_rectal_before_50);null"`
	FamilyCountProstateCancer60         int       `orm:"column(family_count_prostate_cancer_60);null"`
	ProstateCancerBefore50              int       `orm:"column(prostate_cancer_before_50);null"`
	Family2MoreSameCancer60             int       `orm:"column(family_2_more_same_cancer_60);null"`
	CancerBefore50                      int       `orm:"column(cancer_before_50);null"`
	FamilyCountTesticularCancer60       int       `orm:"column(family_count_testicular_cancer_60);null"`
	TesticularCancerBefore50            int       `orm:"column(testicular_cancer_before_50);null"`
	FamilyCountMs60                     int       `orm:"column(family_count_ms_60);null"`
	FamilyCountParkinsons60             int       `orm:"column(family_count_parkinsons_60);null"`
	FamilyCountMotorNeurone60           int       `orm:"column(family_count_motor_neurone_60);null"`
}

func (t *MortgageInsurance) TableName() string {
	return "mortgage_insurance"
}

func init() {
	//orm.RegisterModel(new(MortgageInsurance))
}

// AddMortgageInsurance insert a new MortgageInsurance into database and returns
// last inserted Id on success.
func AddMortgageInsurance(m *MortgageInsurance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMortgageInsuranceById retrieves MortgageInsurance by Id. Returns error if
// Id doesn't exist
func GetMortgageInsuranceById(id int) (v *MortgageInsurance, err error) {
	o := orm.NewOrm()
	v = &MortgageInsurance{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMortgageInsurance retrieves all MortgageInsurance matches certain condition. Returns empty list if
// no records exist
func GetAllMortgageInsurance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MortgageInsurance))
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

	var l []MortgageInsurance
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

// UpdateMortgageInsurance updates MortgageInsurance by Id and returns error if
// the record to be updated doesn't exist
func UpdateMortgageInsuranceById(m *MortgageInsurance) (err error) {
	o := orm.NewOrm()
	v := MortgageInsurance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMortgageInsurance deletes MortgageInsurance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMortgageInsurance(id int) (err error) {
	o := orm.NewOrm()
	v := MortgageInsurance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MortgageInsurance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
