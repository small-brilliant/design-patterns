package Factory

import "testing"

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("l").GetFood()
	NewRestaurant("x").GetFood()
}
