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

// Package routers associates defined routes with controller functions
// that can handle those routes. Note that the Beego framework can handle
// complex routes, but so far this app uses only simple routes.
package routers

import (
	"nohassls/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.AutoRouter(&controllers.LifeController{})

	beego.Router("/life/getall", &controllers.LifeController{}, "get:GetAll")
	beego.Router("/life/get/:id", &controllers.LifeController{}, "get:GetOne")
	beego.Router("/car/getall", &controllers.CarInsurancesController{}, "get:GetAll")
	beego.Router("/car/get/:id", &controllers.CarInsurancesController{}, "get:GetOne")
	beego.Router("/home/getall", &controllers.HomeInsurancesController{}, "get:GetAll")
	beego.Router("/home/get/:id", &controllers.HomeInsurancesController{}, "get:GetOne")

	beego.Router("/", &controllers.HomeController{}, "get:ShowIndex")
	beego.Router("/f", &controllers.HomeController{}, "get:ShowForms")

	beego.Router("/life", &controllers.HomeController{}, "get:ShowLifeInsuranceProfile")
	beego.Router("/user", &controllers.HomeController{}, "get:ShowUserProfile")
	beego.Router("/form", &controllers.HomeController{}, "get:ShowFormValidation")
	beego.Router("/s", &controllers.HomeController{}, "get:ShowSteps")
	beego.Router("/cards", &controllers.HomeController{}, "get:ShowCards")
	beego.Router("/ribbons", &controllers.HomeController{}, "get:ShowRibbons")
	beego.Router("/lists", &controllers.HomeController{}, "get:ShowLists")
	beego.Router("/overlay", &controllers.HomeController{}, "get:ShowOverlay")
	beego.Router("/t1", &controllers.HomeController{}, "get:ShowTable")

	beego.Router("/register", &controllers.HomeController{}, "get,post:Register")
	beego.Router("/login", &controllers.HomeController{}, "get,post:Login")
	beego.Router("/profile", &controllers.HomeController{}, "get,post:Profile")
	beego.Router("/logout", &controllers.HomeController{}, "get:Logout")
	beego.Router("/delete", &controllers.HomeController{}, "get:Delete")
	beego.Router("/restricted", &controllers.HomeController{}, "get:Restricted")

	// TODO add email service in order to provide email verification
	//	and password reset functionality
	//	beego.Router("/verify/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", &controllers.HomeController{}, "get:Verify")
	//	beego.Router("/forgot", &controllers.HomeController{}, "get,post:Forgot")
	//	beego.Router("/reset/:uuid({[0-9A-F]{8}-[0-9A-F]{4}-4[0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}})", &controllers.HomeController{}, "get,post:Reset")
}
