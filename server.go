package main
 
import (
	"flag"
    "github.com/astaxie/beego"  //使用beego框架
)
 
type MainController struct {
    beego.Controller
}
 
func (this *MainController) Get() {
	name := this.Ctx.Input.Param(":name")  //获取路由信息
    this.Ctx.WriteString("hello "+ name)   //写入
}
 
func main() {
	port := flag.String("port", "", "port:default is 8080")  //传入端口号
	flag.Parse()
    beego.Router("/hello/:name", &MainController{})  //路由设置
    beego.Run(":"+ *port)  //运行
}
