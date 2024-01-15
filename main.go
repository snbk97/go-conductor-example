package main

import (
	"ccc/utils"
	"ccc/work"
)

func main() {
	utils.LoadEnv()
	utils.SetupLogger()
	work.InitClients()
	work.InitWorkers() // start custom workers
	// --- setup complete ---

	// work.RunRemoteWorkFlow("MakerChecker2", )
	// work.RunCreatedWorkflow("6", "Email")
	work.RunMakerChecker()

	// keep the program running
	c := make(chan int)
	<-c

}
