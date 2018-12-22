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
	Note       string
}

func NewStarSystem(template int) *StarSystem {
	system := StarSystem{}
	switch template {
	case 0:
		system.Stars = append(system.Stars, *NewStar("Black Hole"))
	case 1:
		system.Stars = append(system.Stars, *NewStar("Neutron Star"))
	case 2:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.Stars[0].description = "Gigantic dying star on the verge of collapsing into a supernova or black hole. Dull red in color. The star has grown so large, it has engulfed most of its inner planets. One remaining inner planet shows signs of ruins on its surface, but its oceans and atmosphere have long since boiled off into space. There are also two gas giants further out."
		system.planetCode = []int{1, 0, 0, 2}
	case 3:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.planetCode = []int{4, 0, 1, 3}
		system.Stars[0].description = "Old Red Giant Star"
	case 4:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.planetCode = []int{5, 1, 0, 4}
		system.Stars[0].description = "Old Red Giant Star"
	case 5:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.planetCode = []int{3, 1, 2, 3}
		system.Stars[0].description = "Old Red Giant Star"
	case 6:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.planetCode = []int{4, 0, 1, 5}
		system.Stars[0].description = "Old Red Giant Star"
	case 7:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.Stars[0].description = "Old Red Giant Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf companion"
		system.planetCode = []int{6, 0, 1, 3}

	case 8:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.Stars[0].description = "Old Red Giant Star"
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow companion"
		system.Note = "Asteroid Belt"
		system.planetCode = []int{7, 0, 0, 1}
	case 9:
		system.Stars = append(system.Stars, *NewStar("O"))
		system.Stars[0].description = "Huge Bright Blue Star"
		system.Note = "Large Asteroid Belt"
		system.planetCode = []int{2, 0, 0, 1}
	case 10:
		system.Stars = append(system.Stars, *NewStar("O"))
		system.Stars[0].description = "Large Bright Blue Star"
		system.Note = "Large Asteroid Belt"
		system.planetCode = []int{1, 0, 0, 0}
	case 11:
		system.Stars = append(system.Stars, *NewStar("O"))
		system.Stars[0].description = "Large Bright Blue Star"
		system.Note = "Large Asteroid Belt"
		system.planetCode = []int{0, 0, 0, 1}
	case 12:
		system.Stars = append(system.Stars, *NewStar("B"))
		system.Stars[0].description = "Big Blue White Star"
		system.Note = "Large Asteroid Belt"
		system.planetCode = []int{3, 1, 1, 4}
	case 13:
		system.Stars = append(system.Stars, *NewStar("A"))
		system.Stars[0].description = "White Star"
		system.planetCode = []int{3, 1, 1, 4}
	case 14:
		system.Stars = append(system.Stars, *NewStar("A"))
		system.Stars[0].description = "White Star"
		system.planetCode = []int{2, 1, 2, 3}
	case 15:
		system.Stars = append(system.Stars, *NewStar("A"))
		system.Stars[0].description = "White Star"
		system.planetCode = []int{3, 2, 2, 5}
	case 16:
		system.Stars = append(system.Stars, *NewStar("F"))
		system.Stars[0].description = "Yellow-White Star"
		system.planetCode = []int{1, 3, 0, 0}
	case 17:
		system.Stars = append(system.Stars, *NewStar("F"))
		system.Stars[0].description = "Yellow-White Star"
		system.planetCode = []int{2, 1, 4, 6}
	case 18:
		system.Stars = append(system.Stars, *NewStar("F"))
		system.Stars[0].description = "Yellow-White Star"
		system.planetCode = []int{3, 2, 2, 4}
		system.Note = "Large Asteroid Belt"
	case 19:
		system.Stars = append(system.Stars, *NewStar("F"))
		system.Stars[0].description = "Yellow-White Star"
		system.planetCode = []int{3, 2, 1, 3}
	case 20:
		system.Stars = append(system.Stars, *NewStar("A"))
		system.Stars[0].description = "White Star"
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[1].description = "Yellow companion"
		system.planetCode = []int{4, 1, 0, 2}
		system.Note = "Large Asteroid Belt"
	case 21:
		system.Stars = append(system.Stars, *NewStar("F"))
		system.Stars[0].description = "Yellow-White Star"
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[1].description = "Orange companion"
		system.planetCode = []int{3, 0, 2, 3}
		system.Note = "Large Asteroid Belt"
	case 22:
		system.Stars = append(system.Stars, *NewStar("F"))
		system.Stars[0].description = "Yellow-White Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Red companion"
		system.planetCode = []int{2, 1, 1, 2}
		system.Note = "Asteroid Belt"
	case 23:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.planetCode = []int{3, 0, 2, 4}
		system.Note = "One Gas Giant has garden type moon"
	case 24:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.planetCode = []int{2, 2, 1, 3}
	case 25:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.planetCode = []int{4, 1, 0, 5}
	case 26:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.planetCode = []int{3, 0, 2, 6}
		system.Note = "One Gas Giant has garden type moon"
	case 27:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.planetCode = []int{2, 1, 1, 4}
		system.Note = "Large Asteroid Belt"
	case 28:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.planetCode = []int{1, 2, 2, 3}
	case 29:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.planetCode = []int{3, 1, 3, 3}

	default:

	}
	system.registerSystem()
	system.fillOrbits()

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
	if s.Note != "" {
		res = res + "Note: " + s.Note + "\n\n"
	}
	for i := range s.Planets {
		res = res + s.Planets[i].toString() + "\n"
	}
	return res
}
