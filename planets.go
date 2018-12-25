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
	p.Gravity = p.createGravity()
	p.Moons = p.createMoons()
	p.Description = "This planet is too close to the sun to support life"
	return &p
}

func NewColdZonePlanet(template int) *Planet {
	p := Planet{}
	p.Gravity = p.createGravity()
	p.Moons = p.createMoons()
	p.Description = "This rocky planet too far from the sun to support life"
	return &p
}

func NewOuterGasGigant(template int) *Planet {
	p := Planet{}
	p.Description = "Large gaseous planet."
	return &p
}

func (p *Planet) createGravity() float64 {
	return float64(rollXdY(4, 6)-4) / 10
}

func (p *Planet) createMoons() string {
	moonSeed := rollXdY(2, 6)
	switch moonSeed {
	case 7:
		return "1 moon"
	case 8:
		return "2 moons"
	case 9:
		return "3 moons"
	default:
		if moonSeed > 9 { // case  10 || 11 || 12
			return "Rings only"
		}
	}
	return "None"
}

func (p *Planet) toString() string {
	str := ""
	str = str + "Planet Name: " + p.Name + "\n"
	str = str + "Gravity            : " + FloatToString(p.Gravity, 1) + "g\n"
	if p.Climate != "" {
		str = str + "Atmosphere Pressure: " + p.Pressure + "\n"
		str = str + "Climate            : " + p.Climate + "\n"
		str = str + "Hydrosphere        : " + strconv.Itoa(p.Hydrosphere) + " %\n"
	}
	str = str + "Moons              : " + p.Moons + "\n"
	str = str + p.Description + "\n"
	return str
}

func (p *Planet) describe() string {
	if p.Description != "" {
		return p.Description
	}
	return "@@No Description"
}
