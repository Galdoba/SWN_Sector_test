package main

import (
	"strconv"
)

type StarSystem struct {
	Name       string
	Coords     []int
	Stars      []Star
	Planets    []Planet
	planetCode []int
}

func NewStarSystem(template int) *StarSystem {
	system := StarSystem{}
	switch template {
	case 3:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.Stars[0].description = "Old Red Giant Star"
		system.planetCode = []int{4, 0, 1, 3}
		system.registerSystem()
		system.fillOrbits()

		//planets :
	default:

	}

	return &system
}

func (system *StarSystem) getStarClasses() string {
	res := ""
	for i := range system.Stars {
		res = res + system.Stars[i].class
	}
	return res
}

func (system *StarSystem) getStarCode() string {
	res := ""
	for i := range system.planetCode {
		res = res + strconv.Itoa(system.planetCode[i])
	}
	return res
}

func (system *StarSystem) registerSystem() {
	system.Name = system.getStarClasses() + "-" + system.getStarCode()
	//"0000-C-4013"
}

func (system *StarSystem) fillOrbits() {
	system.Planets = nil
	plTracker := 0
	for i := range system.planetCode {
		for j := 0; j < system.planetCode[i]; j++ {
			switch i {
			case 0:
				system.Planets = append(system.Planets, *NewHotZonePlanet(1))
			case 1:
				system.Planets = append(system.Planets, *NewGardenPlanet(1))
			case 2:
				system.Planets = append(system.Planets, *NewColdZonePlanet(1))
			case 3:
				system.Planets = append(system.Planets, *NewOuterGasGigant(1))
			default:
			}
			system.Planets[plTracker].Name = system.Name + " " + romanNumberStr(plTracker+1)
			plTracker++
		}
	}
}

func (s *StarSystem) toString() string {
	res := ""
	for i := range s.Stars {
		res = res + s.Stars[i].toString() + "\n"
	}
	for i := range s.Planets {
		res = res + s.Planets[i].toString() + "\n"
	}
	return res
}
