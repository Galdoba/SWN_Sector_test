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
		system.Stars[1].description = "Yellow companion"
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
	case 30:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[1].description = "Orange Companion"
		system.planetCode = []int{4, 0, 2, 4}
		system.Note = "Large Asteroid Belt"
	case 31:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Red Companion"
		system.planetCode = []int{4, 1, 3, 4}
		system.Note = "Asteroid Belt"
	case 32:
		system.Stars = append(system.Stars, *NewStar("G"))
		system.Stars[0].description = "Yellow Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf Companion"
		system.planetCode = []int{4, 2, 0, 5}
		system.Note = "One Gas Giant has garden type moon"
	case 33:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{2, 1, 3, 3}
	case 34:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{2, 2, 3, 4}
		system.Note = "Asteroid Belt"
	case 35:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{1, 1, 2, 3}
	case 36:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{1, 1, 3, 2}
	case 37:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{0, 2, 3, 3}
	case 38:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{0, 1, 3, 3}
	case 39:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{2, 1, 4, 4}
	case 40:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{2, 2, 3, 3}
	case 41:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{1, 2, 3, 2}
	case 42:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{1, 1, 4, 3}
		system.Note = "One Gas Giant has garden type moon"
	case 43:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{2, 0, 2, 3}
		system.Note = "Asteroid Belt"
	case 44:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{2, 2, 3, 4}
	case 45:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{1, 2, 2, 3}
	case 46:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.planetCode = []int{2, 1, 3, 3}
		system.Note = "Asteroid Belt"
	case 47:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[1].description = "Orange Star"
		system.planetCode = []int{2, 1, 1, 2}
		system.Note = "Asteroid Belt"
	case 48:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[1].description = "Orange Star"
		system.planetCode = []int{2, 0, 2, 3}
		system.Note = "Asteroid Belt"
	case 49:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Red Companion"
		system.planetCode = []int{2, 1, 3, 3}
		system.Note = "Asteroid Belt"
	case 50:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Red Companion"
		system.planetCode = []int{2, 1, 2, 3}
		system.Note = "Asteroid Belt"
	case 51:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Red Companion"
		system.planetCode = []int{3, 1, 3, 2}
		system.Note = "Asteroid Belt"
	case 52:
		system.Stars = append(system.Stars, *NewStar("K"))
		system.Stars[0].description = "Orange Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf Companion"
		system.planetCode = []int{2, 0, 2, 3}
	case 53:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 1, 3, 2}
	case 54:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 1, 2, 2}
	case 55:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 0, 4, 1}
		system.Note = "Asteroid Belt"
	case 56:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 0, 3, 3}
	case 57:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 1, 3, 2}
	case 58:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 1, 2, 4}
		system.Note = "One Gas Giant has Garden type moon"
	case 59:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 1, 2, 2}
	case 60:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 1, 3, 2}
	case 61:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 1, 2, 2}
	case 62:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 0, 4, 1}
		system.Note = "Asteroid Belt"
	case 63:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{2, 0, 4, 2}
	case 64:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 0, 3, 4}
	case 65:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 1, 2, 3}
	case 66:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 1, 3, 2}
	case 67:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 0, 3, 3}
	case 68:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 0, 3, 2}
	case 69:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 1, 3, 2}
	case 70:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 1, 2, 4}
	case 71:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 2, 3, 2}
	case 72:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 0, 2, 2}
	case 73:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 0, 4, 3}
	case 74:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 1, 4, 2}
	case 75:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{1, 0, 3, 2}
	case 76:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 1, 3, 3}
	case 77:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.planetCode = []int{0, 1, 2, 2}
	case 78:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{1, 2, 3, 1}
		system.Note = "Asteroid Belt"
	case 79:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{2, 0, 1, 2}
		system.Note = "Asteroid Belt"
	case 80:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{1, 1, 3, 2}
		system.Note = "Asteroid Belt"
	case 81:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{1, 1, 3, 3}
		system.Note = "Asteroid Belt"
	case 82:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{2, 0, 2, 2}
		system.Note = "Asteroid Belt"
	case 83:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{2, 1, 2, 2}
		system.Note = "Asteroid Belt"
	case 84:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{1, 1, 3, 3}
		system.Note = "Asteroid Belt"
	case 85:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[1].description = "Small Red Star"
		system.planetCode = []int{1, 0, 3, 2}
		system.Note = "Asteroid Belt"
	case 86:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf Companion"
		system.planetCode = []int{0, 0, 5, 1}
	case 87:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf Companion"
		system.planetCode = []int{1, 1, 3, 2}
	case 88:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf Companion"
		system.planetCode = []int{1, 1, 1, 3}
	case 89:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf Companion"
		system.planetCode = []int{1, 2, 4, 1}
	case 90:
		system.Stars = append(system.Stars, *NewStar("M"))
		system.Stars[0].description = "Small Red Star"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "White Dwarf Companion"
		system.planetCode = []int{1, 1, 5, 2}
	case 91:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{0, 0, 1, 0}
	case 92:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{0, 0, 1, 1}
	case 93:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{0, 1, 3, 1}
	case 94:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{0, 0, 4, 1}
	case 95:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{0, 0, 3, 0}
	case 96:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{0, 1, 4, 0}
	case 97:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{1, 0, 3, 1}
		system.Note = "Asteroid Belt"
	case 98:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{1, 1, 2, 0}
	case 99:
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[0].description = "Tiny Dim White Dwarf"
		system.Stars = append(system.Stars, *NewStar("D"))
		system.Stars[1].description = "Tiny Dim White Dwarf"
		system.planetCode = []int{0, 1, 2, 0}

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
