package Utils

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

//	显示当前系统基本信息
func ShowSysInf() {

	fmt.Println("★★★★★★★★★★★★★★★★★★★★★★★★")
	fmt.Println("runtime.Version --->", runtime.Version()) //GO的版本
	fmt.Println("runtime.NumCPU --->", runtime.NumCPU())   //CPU核数
	fmt.Println("runtime.GOOS --->", runtime.GOOS)         //运行GO的OS操作系统
	fmt.Println("runtime.GOARCH --->", runtime.GOARCH)     //CPU架构
	fmt.Println("runtime.Version --->", runtime.Version()) //当前GO语言版本
	fmt.Println("time --->", time.Now())                   //系统当前时间
	fmt.Println("★★★★★★★★★★★★★★★★★★★★★★★★")

	//var memStats runtime.MemStats
	//runtime.ReadMemStats(&memStats)
	//fmt.Println("runtime.memStats --->", memStats)

	//获取全部的环境变量
	// data := os.Environ()
	// for _, val := range data {
	//     fmt.Println(val)
	// }

}

//	go不支持三元表达式，可以使用自定义的函数实现
//	例如：max := utils.If(x > y, x, y).(int)
func If(condition bool, trueVal, falseVal interface{}) interface{} {

	if condition {
		return trueVal
	}
	return falseVal
}

/*
	交换int数据：a, b := utils.Swap(2, 9)
	交换字符串数据：A, B := utils.Swap("Li", "Chen")
*/
func Swap(x, y interface{}) (interface{}, interface{}) {
	return y, x
}

//	设置环境变量
func SetEnv(key, value string) error {

	return os.Setenv(key, value)
}

//	取环境变量的值
func GetEnv(key string) string {

	return os.Getenv(key)
}

//取进程ID
func GetPid() int {
	return os.Getpid()
}

func KillByPid(pid int) {
	p, _ := os.FindProcess(pid)
	fmt.Println("KillByPid", p)
	p.Kill()
}

func StartProcessDemo() {
	//例子演示
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	p, _ := os.StartProcess("xxx.exe", []string{"xxx", "1.txt"}, attr)
	p.Release()
	time.Sleep(10000)
	p.Signal(os.Kill)
	os.Exit(10)
}
