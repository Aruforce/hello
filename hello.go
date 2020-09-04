package hello

import "fmt"

func Hello() string {
    return testHello();
}

func testHello() string {
    fmt.Print("this is v1.0.2");
    return "Hello, world.";
}