package MapTravel_test

import "testing"

func TestMapTravel(t *testing.T) {
	map2 := map[int]int{1: 2, 3: 12, 6: 88}
	for key, value := range map2 {
		t.Log(key, value)
	}
}
