package main

import (
	"go-initial/task2/analyzer"
	"go-initial/task2/notifier"
)

func main() {

	slackClient := notifier.NewSlackClient("")
	// testerService := tester.NewTesterService()
	cpuAnalyzer := analyzer.NewCPUAnalyzer(slackClient, 70.0)
	cpuAnalyzer.CheckCPU(90.0)
}
