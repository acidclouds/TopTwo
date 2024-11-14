package main

import "fmt"
import "agent/common/com"
import "agent/common/sys"

func main() {
	fmt.Println("Hello from Darwin!")
	com.CommFunc()
	sys.GetWifi()
}
