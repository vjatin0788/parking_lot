package vehicle

import "testing"

func TestVehicle(t *testing.T) {
	veh := InitVehicle().AddColor("White").AddRegistrationNumber("PB-03-1993")

	if veh != nil && veh.Color != "White" && veh.RegisterationNumber != "PB-03-1993" {
		t.Errorf("Test failed, Wrong vehicle color or registration number")
	}
}
