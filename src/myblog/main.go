//package main

//import (
//	_ "myblog/routers"

//	"github.com/astaxie/beego"
//)

//func main() {

//	beego.Run()
//}
package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
