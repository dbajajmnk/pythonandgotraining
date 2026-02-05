package main

import (
	"fmt"
	"myapp/internal/service"
)

func main() {
	fmt.Println("=== myapp starting ===")

	svc := service.NewService("Team Demo")
	result := svc.AddAndDescribe(10, 20)
	fmt.Println(result)

	fmt.Println("=== myapp finished ===")
}
