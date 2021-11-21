# Golang

## 语言：

### 命令：

- go build 编译.go文件
  - -o 指定生成文件的命令
- go run 运行.go文件
  - go run *.go 直接运行程序，不会编译成exe文件

### 规则：

#### 变量大小写：

- 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

#### 匿名导入(起别名)：

- 在go中，如果只是导入一个包而并不使用导入的包将会导致一个编译错误。

- 但是有时候我们只是想利用导入包而产生的副作用，如执行包内的init()函数。这时候可以在引入的包名前加下划线 _ 。
- import _ "imgae/png"

#### 导入包中的全部方法：

- import . "test" 导入test包中的全部方法给本包使用。

### 变量：

#### 声明：

```go
//不赋值时，默认值是0
var a int 

var b int = 100

//在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
var c = 100 

//省去var关键字，自动匹配（常用）
//但是这种方式不能声明全局变量，只能在函数体内声明。
e := 100

//声明多个变量
var x,y,z int = 100,200,300
var kk,ll = 100,"abc"
var (
	vv  int = 100
    jj bool = true
)
```

#### 常量（只读）：

```go
const length int = 10

//const 来定义枚举类型
const (
	BEIJING = 0
    SHANGHAI = 1
    SHENZHEN = 2
)

//可以在const()添加一个关键字 iota，每行的iota都会累加1，第一行的iota默认是0
const (
	BEIJING = iota
    SHANGHAI
    SHENZHEN
)
```

#### 获取变量的数据类型：

```go
var a int
fmt.Printf("type of a = %T\n",a)
```

### 函数：

```go
//多个返回值,匿名的
fnuc foo1(a string,b string) (int, int){
    return 666,777
}

//多个返回值，有形参名称的
func foo2(a string,b string) (r1 int, r2 int){
    r1 = 100
    r2 = 200
    return 
}
```

### defer:

- Go语言的 defer 语句会将其后面跟随的语句进行延迟处理，在 defer 归属的函数即将返回时，将延迟处理的语句按 defer 的逆序进行执行 (栈)，也就是说，先被 defer 的语句最后被执行，最后被 defer 的语句，最先被执行。

- ```go
  func main() {
      fmt.Println("defer begin")
      // 将defer放入延迟调用栈
      defer fmt.Println(1)
      defer fmt.Println(2)
      // 最后一个放入, 位于栈顶, 最先调用
      defer fmt.Println(3)
      fmt.Println("defer end")
  }
  //打印 defer begin -> defer end -> 3 -> 2-> 1
  ```

- return 先于 defer执行。

### 数组：

#### 定义：

```go
//固定长度的数组
var arr1 [10]int
var arr2 := [10]int{1,2,3,4}

//动态数组，切片slice
arr3 := []int{1,2,3,4}

//声明slice1是一个切片，但是并没有给slice1分配空间
var slice1 = []int
slice1 = make([]int, 3) //开辟3个空间，默认值是0 
slice1 := make([]int,3,5) //创建一个长度为3，容量为5的切片
```

#### 遍历：

```go
for i := 0; i < len(arr1); i++ {
	fmt.Println(arr1[i])
}

for index, value := range arr2{
    //若不使用index，可以改写为匿名变量：for _, value := range arr2
    fmt.Println("index = ", index, "value = ", value)
}
```

#### 操作：

```go
//获取长度
len(arr1)

//获取容量
cap(arr1)

//追加元素
append(arr1,123)

//切片截取
newArr := arr1[0:2] //截取arr1的下标0-2的元素
newArr := arr1[0:] //截取arr1从下标0一直到arr1的最后一个元素
newArr := arr1[:2] //截取arr1的第一个元素到arr1的最后一个元素
newArr := arr1[:] //截取全部元素，即初始化newArr是arr1的引用

//copy 可以将底层数组的slice一起进行拷贝，复制值，指向新地址
copy(newArr,arr1)
```

- 容量cap是指底层数组的大小，长度len指可以使用的大小。

  容量的用处在哪？在与当你用 appen d扩展长度时，如果新的长度小于容量，不会更换底层数组，否则，go 会新申请一个底层数组（容量为旧容量 * 2），拷贝这边的值过去，把原来的数组丢掉。也就是说，容量的用途是：在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡。

  而长度，则是为了帮助你限制切片可用成员的数量，提供边界查询的。所以用 make 申请好空间后，需要注意不要越界。

#### 作为函数参数：

```go
//只能传递数组长度为4的整型数组，值传递
func printArray(myArray [4]int){
	...
}

//动态数组作为参数，引用传递
func printArray2(myArray []int){
    ...
}
```

### map：

#### 定义：

```go
var map1 map[string]string //key为string，value为string的map
make(map[string]string, 10)
map1["one"] = "java"
map1["two"] = "c++"

map2 := make(map[int]string)
map2[1] = "python"

map3 := map([string]string){
    "one" = "typescript"
    "two" = "golang"
}
```

#### 操作：

```go
 //遍历
for key, value := range map1{
	...
}

//删除
delete(map1, "one")

//修改
map1["one"] = "c#"

//作为函数参数（引用传递）
func test(map1 map[string]string){
    ...
}
```

### 结构体：

#### 定义：

```go
//声明一种行的数据类型myint，是int的一个别名
type myint int

//定义一个结构体
type Book struct {
	title string
	auth string
}
```

#### 转为JSON:

```go
import "encoding/json"

type Movie struct{
    Title string `json:"title"`
    Year int `json:"year"`
    Price int `json:'rmb'`
    Actors []string `json:"actors"`
    Name string `json:"-"` //在JSON编码时，这个编码不参与
    Count int `json:"age,string"` //在JSON编码时，将Count起别名为age，转为string类型。名字,类型。中间不能有空格
    Address string `json:"address,omitempty"` //在JSON编码时，如果是空的，则不解析
    date string //以小写字母开头的，在json编码时会忽略掉
}

func main(){
    movie := Movie{"COCO",2017,40,[]string{"mama","coco"}}
    
    //结构体 -> JSON
    jsonStr, err := json.Marshal(movie)
    if err != nil{
        retrun
    }
    fmt.Println(jsonStr) 
    
    // JSON -> 结构体
    myMovie := Movie{}
    err = json.Unmarshal(jsonStr, &myMovie) //将JSON解析到myMovie
    if err != nil{
        return
    }
    fmt.Println("%v\n",myMovie)
}
```



### 面向对象：

#### 类的表示和封装：

```go
type Person struct {
    Name string
    Age int
}

func (this Person) ShowName() {
    fmt.Println("Name = ", this.Name)
}

func (this Person) SetName(newName string) {
    // this 是调用该方法的对象的一个拷贝
    //改进：传入 this *Student
    this.Name = newName
}

func main() {
    //创建对象
    person1 := Person{Name: "Mike", age: 20}
    person1.ShowName()
}
```

#### 继承：

```go	
type Student struct {
    Person //继承了Person类的属性和方法                              
    gender string
}

func (this *Student) Eat(){
    fmt.Println("Eating...")
}

func main(){
    s1 := Student{Person{"John",18},"male"}
    s1.ShowName() //父类的方法
    s1.Eat() //子类的方法
    
    //或者
    var s2 Student
    s2.name = "Lisa"
    s2.age = 44
    s2.gender = "female"
}
```

#### 多态：

```go
//本质是一个指针
type IAnimal interface{
    Sleep()
    GetColor() string
    GetType() string //获取动物种类
}

//具体的类
type Cat struct {
    color string
}

func (this *Cat) Sleep(){
    fmt.Println("Cat is Sleeping")
}

func (this *Cat) GetColor() string{
    return this.color
}

func (this *Cat) GetType() string{
    return "Cat"
}

//对象作为形参实现多态
func showAnimal(animal IAnimal){
    animal.Sleep() //多态
    fmt.Println("color = ",animal.GetColor())
    fmt.Pringtln("type = ",animal.GetType())
}

func main(){
    var animal IAnimal //接口的数据类型，父类指针
    animal = &Cat{"Green"} //接收一个对象的引用
    animal.Sleep() //调用Cat的Slepp()方法
    
    cat := Cat{"Red"}
    showAnimal(&cat)
}
```

### interface空接口万能类型：

```go
//interface{} 是万能数据类型
func test(arg interface{}){
    fmt.Println(arg)
    
    //给interface{} 提供“类型断言” 的机制
    value, ok := arg.(string)
    if !ok {
        fmt.Println("arg is not string type")
    }else{
        fmt.Println("arg is string type, value = ",value )
    }
}

type Book struct{
    auth string
}

func main(){
    book := Book{"Golang"}
    test(book)
}
```

### 反射：

#### 获取变量的type和value：

```go
import "reflect"

//通过反射获取变量的type和value
func test(arg interface{}){
    fmt.Println("type : ", reflect.Typeof(arg))
    fmt.Println("value : ", reflect.Valueof(arg))
}

func main(){
    var num float64 = 1.2345
    test(num)
}
```

#### 获取结构体元素的type和value：

```go
type Student struct{
	Id int
	Name string
}

func (this Student) Call(){}

func DoFiledAndMethod(input interface{}){
    //获取input的type
    inputType := reflect.TypeOf(input)
    fmt.Println(inputType.Name()) //Student
    
    //获取input的value
    inputValue := reflect.ValueOf(input)
    fmt.Println(inputValue) //{1 Mike}
    
    //通过type获取里面的字段
    //1. 获取interface的reflect.Type，通过Type得到NumFiled，进行遍历
    //2. 得到每个filed，数据类型
    //3. 通过filed有一个Interface()方法得到对应的value
    for i := 0; i < inputType.NumField(); i++{
        field := inputType.Filed(i)
        value := inputValue.Filed(i).Interface()
        
        fmt.Printf(filed.Name, filed.Type, value) //Id: int = 1, Name: string = Mike
    }
    //通过Type获取方法
    for i := 0; i < inputType.NumMethod(); i++{
        m := inputType.Method(i)
        fmt.Printf(m.Name, m.Type) //Call: func(main.Student)
    }
}

func main(){
    s1 := Student{1,"Mike"}
    DoFiledAndMethod(s1)
}
```

获取结构体的tag：

```go
type resume struct{
    Name string `info:"name" doc:"我的名字"`
    Gender string `info: "gender"`
}

func findTag(str interface{}){
    t := reflect.Typeof(str).Elem()
    for i := 0; i < t.NumFiled(); i++{
        tagstring := t.Filed(i).Tag.Get("info")
        fmt.Println(tagstring) //name, gender
    }
}
```

### goroutine：

```GO	
func main(){
    //用go创建一个匿名自执行函数
    go func(){
        defer fmt.Println("A.defer")
        
        func(){
            defer fmt.Println("B.defer")
            //可以通过runtime.Goexit()提前退出当前goroutine
            runtime.Goexit()
            fmt.Println("B")
        }()
        
        fmt.Println("A")
    }()
}
```

### channel：

#### 无缓冲的channel：

- ```go
  func main(){
  	//定义一个无缓冲channel
      c := make(chan int)
      
      go func(){
          defer fmt.Println("goroutine结束") 
          c <- 666 //将666 发送给c
          //若接收c数据的操作未执行，则会阻塞。等待num := <- c执行完毕
      }
      
      num := <-   c //从c中接收数据，并发送给num
      //若协程中还未给c发送值。给num发送的操作会阻塞，等待有值
  }
  ```

#### 有缓冲的channel：

- 不会有阻塞的限制。
- 一个goroutine可以不断向通道发送新值，而接收的goroutine可以不断从通道接收。两个操作既不是同步的，也不会相互阻塞。

- ```go
  func main(){
      c := make(chan int, 3) //定义一个有缓冲的channel
      
      go func(){
          for i := 0; i < 4; i++{
              c <- i //channel的最大容量为3，发送3个后，需要等main先拿走3个，才能放入第4个
          }
      }
      
      time.Sleep(2 * time.Second)
      
      for i := 0; i < 4; i++{
          num := <-c
      }
  }
  ```

#### channel的关闭：  

- ```go
  func main(){
      c := make(chan int)
      go func(){
      for i := -; i < 5; i++{
          c <- i
      }
      
      close(c) //close可以关闭channel
  	} 
  
  	for{
     		 //ok为true表示channel没有关闭
      	if data, ok := <-c; ok {
      	    fmt.Println(data)
      	}else{
      	    break
     	    }
  	}
  }
  ```

- channel不像文件已一样需要经常关闭，只有我们确实没有任何发送数据了，或者想显式地结束range循环之类的情况，才去关闭channel。

- 关闭channel后，无法再向channel发送数据。

- 关闭channel后，可以继续从channel接收数据。

#### range：

```go
func main(){
    c := make(chan int)
    go func(){
    for i := -; i < 5; i++{
        c <- i
    }
    
    close(c) //close可以关闭channel
	} 

    //可以使用range来迭代不断操作channel
    for data := range c{
        fmt.Println(data)
    }
}
```

#### select：

```go
//单流程下一个go只能监控一个channel的状态，select可以完成监控多个channel的状态
select{
    case <- chan1:
    	//如果chan1成功读到数据，则进行该case处理语句
	case chan2 <-1:
    	//如果成功向chan2写入数据，则进行该case语句
    default:
    	//如果上面的都没有成功，则额进入default处理流程
}
```

### Go Modules：

#### 命令：

```go
go mod init //生成go.mod文件
go mod download //下载go.mod文件中指明的所有依赖
go mod tidy //整理现有的依赖
go mod graph //查看现有的依赖结构
go mod edit //编辑go.mod文件
go mod vendor //导出项目所有依赖到vendor目录
go mod verify //校验一个模块是否被篡改过
go mod why //查看为什么需要依赖某模块
```



