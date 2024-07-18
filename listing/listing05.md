Что выведет программа? Объяснить вывод программы.
```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```
Ответ:
```
Программа выведет:
"error"
Когда функция test возвращает nil типа *customError, это значение присваивается переменной err, но тип информации (*customError) сохраняется в интерфейсе. Поэтому, несмотря на то, что значение внутри интерфейса nil, сам интерфейс err не равен nil.
Так как err содержит интерфейсный тип, который не равен nil (несмотря на то, что значение внутри nil), условие if err != nil будет истинным.
```