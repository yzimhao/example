package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {
    fmt.Printf("hello world\n")
    // sque()
    // http_request()
    test_structs()
}

func sque() {

    for x := 1; x < 10; x++ {
        for y := 1; y < 10; y++ {
            if x >= y {
                fmt.Printf("%d * %d = %d    ", y, x, x*y)
            }
        }
        fmt.Printf("\n")
    }
}

func http_request(){

    url := "http://www.cnblogs.com/yjf512/archive/2012/06/18/2554066.html"

    response,_ := http.Get(url)
    // defer response.Body.Close()

    body,_ := ioutil.ReadAll(response.Body)
    fmt.Printf(string(body))
}

func test_structs() {
    type userinfo struct {
        name string
        age int
        address string
    }


    user := userinfo{name:"liaodong", age:27, address:"gangdong shenzhen"}
    fmt.Printf("name: %s age: %d address: %s \n", user.name, user.age, user.address)

}