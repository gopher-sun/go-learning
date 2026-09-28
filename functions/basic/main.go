package main

import "fmt"

func sayhello() {
	fmt.Println("hello")
}
func param1(id string) {
	fmt.Println(id)
}

func add(a1 int, a2 int) {
	fmt.Println(a1, a2)
}
func add1(a1, a2 int) {
	fmt.Println(a1, a2)
}
func add2(numberList ...int) { //这种写法是任意数量传值
	var sum int
	for _, i := range numberList {
		sum += i
	}
	fmt.Println(sum)
}
func r1() bool {
	return true
} //return一个类型
/*func r1() bool{
var ok bool
return ok
}
*/
func r2() (string, bool) {

	return "", true
}
func r3() (val string, ok bool) {
	if 2 > 1 {
		val = "222"
		return val, ok
	}

	return

}
func main() {
	sayhello()
	param1("123")
	add2(2, 4, 5, 6, 111, 7, 8, 0)
	fmt.Println(r3())
	//匿名函数
	var getname = func() string {
		return "lele"
	}
	var setname = func(name string) {
		fmt.Println(name)

	}
	fmt.Println(getname())
	setname("lele")
	var index int

	fmt.Println("1.登陆")
	fmt.Println("2.注册")
	fmt.Println("3.个人中心")
	// switch index{
	//case 1:log()
	//case 2:reg()
	//case 3:user()
	//}
	fmt.Scan(&index)
	var funMap = map[int]func(){
		1: log,
		2: reg,
		3: user,
	}
	fun, ok := funMap[index]
	if ok {
		fun()
	}
}
func log() {
	fmt.Println("登陆")
}
func reg() {
	fmt.Println("注册")
}
func user() {
	fmt.Println("个人中心")
}
