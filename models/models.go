/* Copyright 2014 60plusadventures Author. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

// Packate models defines types and methods used to interact with user
// input data and data associated with a table in the DB
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// User struct used to read and write to DB table "users"
// see const category in home.go for Interest codes
type User struct {
	Id       int       // Database primary key. AutoIncrement value
	Username string    `orm:"size(30);unique"`
	Fullname string    `orm:"size(60)"`
	Email    string    `orm:"size(256)"` // encoded email address
	Password string    `orm:"size(128)"` // password hash value
	AutoLog  bool      // true if user selected "Remember Me"
	Interest uint8     // User's primary interest as a category code
	RegKey   string    `orm:"size(60)"` // used to confirm registration email
	ResetKey string    `orm:"size(60)"` // used to confirm password reset
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

// User database CRUD methods include Insert, Read, Update and Delete
func (usr *User) Insert() error {
	if _, err := orm.NewOrm().Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Read(fields ...string) error {

	if err := orm.NewOrm().Read(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(usr, fields...); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if _, err := orm.NewOrm().Delete(usr); err != nil {
		return err
	}
	return nil
}

// RegFrm struct to hold Registration page and Profile page form values
type RegFrm struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Fullname string `form:"fullname"`
	Password string `form:"pw1"`
	Confirm  string `form:"pw2"`
	AutoLog  string `form:"autolog"`
}

type UserProfileFrm struct {
	Id               int       `form:"id"`
	UserId           int       `form:"user_id"`
	Firstname        string    `form:"firstname"`
	Middlename       string    `form:"middlename"`
	Surname          string    `form:"surname"`
	Sex              string    `form:"sex"`
	Birthdate        time.Time `form:"birthdate"`
	Height           int       `form:"height"`
	Weight           int       `form:"weight"`
	Smoker           uint64    `form:"smoker"`
	Avdrinks         int       `form:"avdrinks"`
	Address1         string    `form:"address1"`
	Address2         string    `form:"address2"`
	Address3         string    `form:"address3"`
	Suburb           string    `form:"suburb"`
	State            string    `form:"state"`
	Postcode         int       `form:"postcode"`
	Email1           string    `form:"email1"`
	Email2           string    `form:"email2"`
	PhoneHome        string    `form:"phone_home"`
	PhoneMob         string    `form:"phone_mob"`
	Msn              string    `form:"msn"`
	Twitter          string    `form:"twitter"`
	Facebook         string    `form:"facebook"`
	Wechat           string    `form:"wechat"`
	Profession       string    `form:"profession"`
	Profilecol       string    `form:"profilecol"`
	Title            string    `form:"title"`
	Child1Name       string    `form:"child1_name"`
	Child1Birthdate  time.Time `form:"child1_birthdate"`
	Child1Sex        string    `form:"child1_sex"`
	Child2Name       string    `form:"child2_name"`
	Child2Birthdate  time.Time `form:"child2_birthdate"`
	Child2Sex        string    `form:"child2_sex"`
	Child3Name       string    `form:"child3_name"`
	Child3Birthdate  time.Time `form:"child3_birthdate"`
	Child3Sex        string    `form:"child3_sex"`
	Child4Name       string    `form:"child4_name"`
	Child4Birthdatae time.Time `form:"child4_birthdatae"`
	Child4Sex        string    `form:"child4_sex"`
	Child5Name       string    `form:"child5_name"`
	Child5Birthdate  time.Time `form:"child5_birthdate"`
	Child5Sex        string    `form:"child5_sex"`
	NoChildren       int       `form:"no_children"`
	NoPets           int       `form:"no_pets"`
	Pet1Name         string    `form:"pet1_name"`
	Pet1Type         string    `form:"pet1_type"`
	Pet1Sex          string    `form:"pet1_sex"`
	Pet2Name         string    `form:"pet2_name"`
	Pet2Type         string    `form:"pet2_type"`
	Pet2Sex          string    `form:"pet2_sex"`
	Pet3Name         string    `form:"pet3_name"`
	Pet3Sex          string    `form:"pet3_sex"`
	Pet3Type         string    `form:"pet3_type"`
	LastUpdated      time.Time `form:"last_updatedtime"`
}

// Comment struct used to read and write to DB table "comment"
// see const category in home.go for Category codes
type Comment struct {
	Id       int    // Database primary key. AutoIncrement value
	User     *User  `orm:"rel(fk);index"` // Indexed Foreign Key -> User.Id
	Txt      string `orm:"size(4096)"`
	Category uint8  // Indicates page used to post the comment
	Archived bool
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

// Comment database access methods include Insert and Read
// TODO Add Archive and review need for Delete
func (cmnt *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(cmnt); err != nil {
		return err
	}
	return nil
}

func (cmnt *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(cmnt, fields...); err != nil {
		return err
	}
	return nil
}

// Register User and Comment models with the Beego ORM
func init() {
	orm.RegisterModel(new(User))
}
