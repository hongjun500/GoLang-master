# 函数

- 在 `go` 中所有东西都是变量, 函数也是变量的一种，因此函数名的命名规则也和变量一样
- 使用关键字 `func`来声明和定义

```go
func funcName (parameter-list) (result-list) {
body
}
```

- 函数的主要特点
    - 函数可以使用 `...T` 的方式来定义参数，表示参数的个数是可变的
    - 函数可以返回多个值
    - 函数的返回值可以被命名，并且像变量那样使用，即
  > 返回值可以用括号显示声明出来，即一开始就定义好，那么在函数体中可以直接调用，函数最终直接一个 `return`
    - 函数可以作为变量的值
    - 函数可以作为参数和返回值
    - 隐藏函数名的匿名函数可以是闭包的，闭包是一个函数值，*它引用了其函数体之外的变量或者常量的内存地址，而不是它们的值*

- 函数的参数可以是值传递，也可以是引用传递，即
    - 值传递：传递的是值的副本，不会改变原有的值
    - 引用传递：传递的是值的内存地址，会改变原有的值
    - 函数的参数可以是另一个函数,带函数参数的函数示例如下

        ```go
        func funcName (f func (int) int) {
            // do something
        }
        ```

## 匿名函数

- 匿名函数可以定义在函数内部、作为函数的参数或者返回值
- 匿名函数的闭包可以在函数内部引用外部变量或者外部常量，这里的**引用**其实是*引用了其内存地址，而不是引用了其值*。
  这通常不需要通过参数传递，因为匿名函数可以直接使用外部变量或者外部常量

  ```go
    func funcName (parameter-list) (result-list) {
        body
        var fc = func (name string) (int) {
            i := 0
            fmt.Printf("匿名函数接收到的参数是：%s\n", name)
            return i
        }
        // 调用匿名函数
		fc("hello")
        
		// 直接使用匿名函数
        func (name string) (int) {
            i := 0
            fmt.Printf("匿名函数接收到的参数是：%s\n", name)
            return i
        }() // 注意这里的括号，表示直接调用匿名函数
        
		// defer 关键字配合闭包使用
        str := "hello world"
        defer func () { // 这里的 defer 关键字会使此匿名函数延迟调用, 并且 str 的内存地址会被 defer 关键字保留，而并不是保留其值的内容
            fmt.Println("str = ", str) // hello golang 
        }()
        // 这里改变了变量 str 的值，内存地址并没有改变，因此在匿名函数中打印的值是改变后的值
		str = "hello golang" 
        
    }
    ```

> **小写字母开头的函数只能在本包内使用，大写字母开头的函数才可以在其它包中使用，此规则同时适用于变量和类型**
  