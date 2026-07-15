package mining

import "time"

const TargetBlockTime = 10 * time.Second

// AdjustDifficulty adjusts mining difficulty based on the previous block time.
func AdjustDifficulty(currentDifficulty int, previousTime, currentTime time.Time) int {

	elapsed := currentTime.Sub(previousTime)

	if elapsed < TargetBlockTime {
		return currentDifficulty + 1
	}

	if elapsed > TargetBlockTime && currentDifficulty > 1 {
		return currentDifficulty - 1
	}

	return currentDifficulty
}
