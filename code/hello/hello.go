package main

import "github.com/astaxie/beego"

type MainController struct{
	beego.Controller
}

func (this *MainController)Get()  {
	this.Ctx.WriteString("hello world")
}

func main()  {
	beego.Router("/",&MainController{}) // Router函数的两个参数函数，第一个是路径，第二个是Controller的指针
	beego.Run()
}