package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/SarayutBz/LearnGolang/ballsarayut"
)
func sayHello() {
	fmt.Println("Hello, 世界")
}
func main() {
	id := uuid.New()
	sayHello()
	fmt.Printf("uuid: %s\n", id)
	ballsarayut.SayHelloBallsarayut()
}
