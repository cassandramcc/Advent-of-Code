package day6

func SolveOne() int {
	//testTimes := []int{7, 15, 30}
	//testDistances := []int{9, 40, 200}

	times := []int{35, 93, 73, 66}
	distances := []int{212, 2060, 1201, 1044}
	wins := 1
	for i := range times {
		wins *= calculateTotalWins(times[i], distances[i])
	}
	return wins
}

func SolveTwo() int {
	//testTimes := []int{71530}
	//testDistances := []int{940200}

	times := []int{35937366}
	distances := []int{212206012011044}
	wins := 1
	for i := range times {
		wins *= calculateTotalWins(times[i], distances[i])
	}
	return wins
}

func calculateRaceDistance(raceTime, buttonHold int) int {
	return (raceTime - buttonHold) * buttonHold
}

func calculateRaceDistances(raceTime int) []int {
	var distances []int
	for i := 0; i <= raceTime; i++ {
		distances = append(distances, calculateRaceDistance(raceTime, i))
	}
	return distances
}

func calculateTotalWins(raceTime, recordDistance int) int {
	distances := calculateRaceDistances(raceTime)
	var total int

	for _, distance := range distances {
		if distance > recordDistance {
			total += 1
		}
	}
	return total
}
