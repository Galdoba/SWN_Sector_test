package main

import "strconv"

type Planet struct {
	Name        string
	Gravity     float64
	Moons       string
	Pressure    string
	Climate     string
	Hydrosphere int
	Ecosystem   string
	EModRating  int
	Description string
	nextStep    string
}

func NewGardenPlanet(template int) *Planet {
	p := Planet{}
	switch template {
	case 1:
		p.Gravity = 0.1
		p.Moons = "None"
		p.Pressure = "Trace"
		p.Climate = "Frigid (< -50c)"
		p.Hydrosphere = 0
		p.Ecosystem = "Lifeless"
		p.EModRating = 0
		p.nextStep = "H"

	default:
	}
	return &p
}

func NewHotZonePlanet(template int) *Planet {
	p := Planet{}
	p.Description = "This planet is too close to the sun to support life"
	return &p
}

func NewColdZonePlanet(template int) *Planet {
	p := Planet{}
	p.Description = "This rocky planet too far from the sun to support life"
	return &p
}

func NewOuterGasGigant(template int) *Planet {
	p := Planet{}
	p.Description = "Large gaseous planet."
	return &p
}

func (p *Planet) toString() string {
	str := ""
	str = str + "Planet Name: " + p.Name + "\n"
	if p.Climate != "" {
		str = str + "Gravity            : " + FloatToString(p.Gravity, 1) + "g\n"
		str = str + "Atmosphere Pressure: " + p.Pressure + "\n"
		str = str + "Climate            : " + p.Climate + "\n"
		str = str + "Hydrosphere        : " + strconv.Itoa(p.Hydrosphere) + " %\n"
	}
	str = str + p.Description + "\n"
	return str
}

func (p *Planet) describe() string {
	if p.Description != "" {
		return p.Description
	}
	return "@@No Description"
}
