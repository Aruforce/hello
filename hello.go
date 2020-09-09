package hello

import "fmt"

func Hello() string {
    return testHello();
}

func testHello() string {
    fmt.Print("this is v2.0.0");
    return "Hello, world.";
}