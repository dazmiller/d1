// Package main defines functions to initialize database access via Beego's ORM,
// enables persistent sessions and starts the app's built-in server
package main

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// Preload values used by Beego on startup
	_ "nohassls/models"
	_ "nohassls/routers"

	_ "github.com/go-sql-driver/mysql"
)

// Init resgisters the sixtyplus database with Beego's ORM
// IMPORTANT: Edit nohassls/conf/app.conf changing 'mysqluser', 'mysqlpass'
// and 'mysqldb' to the values you use to access your MySQL database
// If the DB is not on the app's server at localhost:3306, edit
// the information between '@tcp(' and ')/' as appropriate
func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/" +
		beego.AppConfig.String("mysqldb") + "?charset=utf8&parseTime=true&loc=Australia%2FSydney"
	orm.RegisterDataBase("default", "mysql", mysqlReg)
}

// Main is the program's entry point after all imports and
// related init functions have been loaded/executed
// Explicit Beego options can set here before calling Beego.Run()
func main() {

	beego.AddFuncMap("hi", hello)

	//beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

// additional helper functions
func hello(in string) (out string) {
	out = in + "world"
	return
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

var unKind = map[reflect.Kind]bool{
	reflect.Uintptr:       true,
	reflect.Complex64:     true,
	reflect.Complex128:    true,
	reflect.Array:         true,
	reflect.Chan:          true,
	reflect.Func:          true,
	reflect.Map:           true,
	reflect.Ptr:           true,
	reflect.Slice:         true,
	reflect.Struct:        true,
	reflect.UnsafePointer: true,
}

// isValidForInput checks if fType is a valid value for the `type` property of an HTML input element.
func isValidForInput(fType string) bool {
	validInputTypes := strings.Fields("text password checkbox radio submit reset hidden image file button search email url tel number range date month week time datetime datetime-local color")
	for _, validType := range validInputTypes {
		if fType == validType {
			return true
		}
	}
	return false
}

// RenderForm will render object to form html.
// obj must be a struct pointer.
func RenderNoHasslsForm(obj interface{}) template.HTML {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if !isStructPtr(objT) {
		return template.HTML("")
	}
	objT = objT.Elem()
	objV = objV.Elem()

	var raw []string
	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() || unKind[fieldV.Kind()] {
			continue
		}

		fieldT := objT.Field(i)

		label, name, fType, id, class, ignored, required := parseNhFormTag(fieldT)
		if ignored {
			continue
		}

		raw = append(raw, renderNhFormField(label, name, fType, fieldV.Interface(), id, class, required))
	}
	return template.HTML(strings.Join(raw, "</br>"))
}

// renderFormField returns a string containing HTML of a single form field.
func renderNhFormField(label, name, fType string, value interface{}, id string, class string, required bool) string {
	if id != "" {
		id = " id=\"" + id + "\""
	}

	if class != "" {
		class = " class=\"" + class + "\""
	}

	requiredString := ""
	if required {
		requiredString = " required"
	}

	if isValidForInput(fType) {
		return fmt.Sprintf(`%v<input%v%v name="%v" type="%v" value="%v"%v>`, label, id, class, name, fType, value, requiredString)
	}

	return fmt.Sprintf(`%v<%v%v%v name="%v"%v>%v</%v>`, label, fType, id, class, name, requiredString, value, fType)
}

// parseFormTag takes the stuct-tag of a StructField and parses the form value.
// returned are the form label, name-property, type and wether the field should be ignored.
func parseNhFormTag(fieldT reflect.StructField) (label, name, fType string, id string, class string, ignored bool, required bool) {
	tags := strings.Split(fieldT.Tag.Get("form"), ",")
	label = fieldT.Name + ": "
	name = fieldT.Name
	fType = "text"
	ignored = false
	id = fieldT.Tag.Get("id")
	class = fieldT.Tag.Get("class")

	required = false
	required_field := fieldT.Tag.Get("required")
	if required_field != "-" && required_field != "" {
		required, _ = strconv.ParseBool(required_field)
	}

	switch len(tags) {
	case 1:
		if tags[0] == "-" {
			ignored = true
		}
		if len(tags[0]) > 0 {
			name = tags[0]
		}
	case 2:
		if len(tags[0]) > 0 {
			name = tags[0]
		}
		if len(tags[1]) > 0 {
			fType = tags[1]
		}
	case 3:
		if len(tags[0]) > 0 {
			name = tags[0]
		}
		if len(tags[1]) > 0 {
			fType = tags[1]
		}
		if len(tags[2]) > 0 {
			label = tags[2]
		}
	}

	return
}
