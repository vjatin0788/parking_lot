package vehicle

type Vehicle struct {
	RegisterationNumber string
	Color               string
}

func InitVehicle() *Vehicle {
	return &Vehicle{}
}

func (v *Vehicle) AddRegistrationNumber(number string) *Vehicle {
	v.RegisterationNumber = number
	return v
}

func (v *Vehicle) AddColor(color string) *Vehicle {
	v.Color = color
	return v
}
