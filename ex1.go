package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {

	var len = len(droneIds)
	var missing int
	var found int = 0
	for i := 0; i < len; i++ {
		found = 0
		for j := 0; j < len; j++ {
			if droneIds[i] == droneIds[j] && i != j {
				found = 1
				break
			}
		}
		if found == 0 {
			missing = droneIds[i]
			break
		}
	}
	return missing
}
