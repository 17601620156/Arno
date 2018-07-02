/**
  **日志监控系统
  **读取
  **解析
  **写入
 **/
package main

import (
	"strings"
	"fmt"
	"time"
)

//结构体

type LogProcess struct {
    //channels进行通信
    rc chan string
    wc chan string

    //路径和写入用户名密码
	path string
	influxDBDsn string
	 		
}

//读取模块
//  *为引用  l就可修改自身定义的一些参数
func (l *LogProcess) ReadFromFile() {
	//模拟读取到数据
	line := "hello world"
	l.rc <- line
}

//解析模块
func (l *LogProcess) Process() {
    //接收传过来的数据
    data := <-l.rc
    //转化为大写
    l.wc <- strings.ToUpper(data)

}

//写入模块
func (l *LogProcess) WriteToInfluxDB() {

	fmt.Print(<-l.wc)

}

func main()  {

	 lp := &LogProcess{

	 	rc: make(chan string),
	 	wc: make(chan string),
	 	path: "/tmp/access.log",
	 	influxDBDsn: "username&password...",

	 }

	 //创建3个goroutine
	 go lp.ReadFromFile()
	 go lp.Process()
	 go lp.WriteToInfluxDB()

	 time.Sleep(1*time.Second)//当创建完3个程序就会自动退出  不会等待这些goroutine执行完才退出  所以让等待1秒


}
