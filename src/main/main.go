package main

import (
	c "PBP-API/src/controllers"
)

func main() {
	//go routine
	c.TaskScheduler()
	select {}

}
