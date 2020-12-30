package creational

// Creation Pattern - Builder Pattern

/**
1. Abstract complex creations so that object creation is separated from the object
user
2. Create an object step by step by filling its fields and creating the embedded
objects
3. Reuse the object creation algorithm between many objects
*/

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
