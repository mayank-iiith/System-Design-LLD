package helper

import (
	"math/rand"
	"time"
)

func SimulateNetworkLatency() {
	sleepIntervalInSeconds := 1 + rand.Float32()*(2) // Random float between 1 and 3
	time.Sleep(time.Duration(sleepIntervalInSeconds) * time.Second)
}
