package analyzer

type ChannelSender interface {
	SendToChannel(channel string, message string) error
}

type CPUAnalyzer struct {
	sender   ChannelSender
	cpuLimit float64
}

func NewCPUAnalyzer(sender ChannelSender, cpuLimit float64) *CPUAnalyzer {

	return &CPUAnalyzer{
		sender:   sender,
		cpuLimit: cpuLimit,
	}
}

func (c *CPUAnalyzer) CheckCPU(currentCPULoad float64) {

	if currentCPULoad >= c.cpuLimit {
		c.sender.SendToChannel("#alerts", "CPU usage is high")

	}

}
