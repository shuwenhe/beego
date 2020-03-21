package main

// 1、首先我们引入了包github.com/astaxie/beego,我们知道Go语言里面引入包会深度优先的去执行引入包的初始化(变量和init函数，更多)，beego包中会初始化一个BeeAPP的应用，初始化一些参数。
import "github.com/astaxie/beego" 

// 2、定义Controller，这里我们定义了一个struct为MainController，充分利用了Go语言的组合的概念，匿名包含了beego.Controller，这样我们的MainController就拥有了beego.Controller的所有方法。
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