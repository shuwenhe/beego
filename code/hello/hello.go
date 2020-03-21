package main

// 1、首先我们引入了包github.com/astaxie/beego,我们知道Go语言里面引入包会深度优先的去执行引入包的初始化(变量和init函数，更多)，beego包中会初始化一个BeeAPP的应用，初始化一些参数。
import "github.com/astaxie/beego" 

// 2、定义Controller，这里我们定义了一个struct为MainController，充分利用了Go语言的组合的概念，匿名包含了beego.Controller，这样我们的MainController就拥有了beego.Controller的所有方法。
type MainController struct{
	beego.Controller
}

// 3、定义RESTFul方法，通过匿名组合之后，其实目前的MainController已经拥有了Get、Post、Delete、Put等方法，这些方法是分别用来对应用户请求的Method函数，如果用户发起的是POST请求，那么就执行Post函数。所以这里我们定义了MainController的Get方法用来重写继承的Get函数，这样当用户GET请求的时候就会执行该函数。
func (this *MainController)Get()  {
	this.Ctx.WriteString("hello world")
}

// 4、定义main函数，所有的Go应用程序和C语言一样都是Main函数作为入口，所以我们这里定义了我们应用的入口。
func main()  {
// 5、Router注册路由，路由就是告诉beego，当用户来请求的时候，该如何去调用相应的Controller，这里我们注册了请求/的时候，请求到MainController。这里我们需要知道，Router函数的两个参数函数，第一个是路径，第二个是Controller的指针。
	beego.Router("/",&MainController{}) // Router函数的两个参数函数，第一个是路径，第二个是Controller的指针
// 6、Run应用，最后一步就是把在1中初始化的BeeApp开启起来，其实就是内部监听了8080端口:Go默认情况会监听你本机所有的IP上面的8080端口	
	beego.Run()
}