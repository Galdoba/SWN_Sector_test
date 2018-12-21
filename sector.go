package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Planet struct {
	Name        string
	Gravity     float64
	Moons       string
	Pressure    string
	Climate     string
	Hydrosphere int
	Ecosystem   string
	EModRating  int
	nextStep    string
}

type Star struct {
	starColor   string
	class       string
	temperature float64
	mass        float64
	description string
}

type StarSystem struct {
	Name    string
	Stars   []Star
	Planets []Planet
}

func NewStar(clss string) *Star {
	star := Star{}
	star.class = clss
	switch clss {
	case "O":
		star.starColor = "Bright Blue"
		star.temperature = randFloat(33.0, 99.0, 3)
		star.mass = randFloat(16.0, 50.0, 2)
	case "B":
		star.starColor = "Blue-White"
		star.temperature = randFloat(10.0, 33.0, 3)
		star.mass = randFloat(2.1, 16.0, 2)
	case "A":
		star.starColor = "White"
		star.temperature = randFloat(7.5, 10.0, 3)
		star.mass = randFloat(1.4, 2.1, 2)
	case "F":
		star.starColor = "Yellowish-White"
		star.temperature = randFloat(6, 7.5, 3)
		star.mass = randFloat(1.04, 1.4, 2)
	case "G":
		star.starColor = "Yellow"
		star.temperature = randFloat(5.2, 6, 3)
		star.mass = randFloat(0.8, 1.04, 2)
	case "K":
		star.starColor = "Orange"
		star.temperature = randFloat(3.7, 5.2, 3)
		star.mass = randFloat(0.45, 0.8, 2)
	case "M":
		star.starColor = "Red"
		star.temperature = randFloat(3, 3.7, 3)
		star.mass = randFloat(0.35, 0.45, 2)
	case "C":
		star.starColor = "Dim Red"
		star.temperature = randFloat(2.7, 3, 3)
		star.mass = randFloat(0.26, 0.35, 2)
		star.description = "Red Giant (Swollen Dying Star)"
	case "T":
		star.starColor = "Faint Brown"
		star.temperature = randFloat(2.3, 2.7, 3)
		star.mass = randFloat(0.17, 0.26, 2)
		star.description = "Brown Dwarf (Tiny Proto Star)"
	case "D":
		star.starColor = "Faint White"
		star.temperature = randFloat(1.6, 2.3, 3)
		star.mass = randFloat(0.14, 0.17, 2)
		star.description = "White Dwarf (Star Remnant)"
	case "Black Hole":
		star.starColor = "Black"
		star.temperature = randFloat(0, 0, 3)
		star.mass = randFloat(0, 0, 2)
		r := randInt(1, 2)
		if r == 1 {
			star.description = "Black Hole with companion orange star it is being drawn into its gravity well. Powerful tidal forces make this system extremly Hazardous"
		} else {
			star.description = "The system contains the broken remnants of planets."
		}

	default:
		star.starColor = ""
		star.temperature = randFloat(0, 0, 3)
		star.mass = randFloat(0, 0, 2)
		star.description = "UNKNOWN TYPE"
	}
	return &star
}

func (s *Star) toString() string {
	str := ""
	str = str + "Star Class : " + s.class + "\n"
	str = str + "Star Color : " + s.starColor + "\n"
	str = str + "Temperature: " + FloatToString(s.temperature, 3) + " k\n"
	str = str + "Mass       : " + FloatToString(s.mass, 2) + "  M\n"
	str = str + "Description: " + s.description + "\n"
	return str
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

func (p *Planet) toString() string {
	str := ""
	str = str + "Planet Name        : " + p.Name + "\n"
	str = str + "Gravity            : " + FloatToString(p.Gravity, 1) + "g\n"
	str = str + "Atmosphere Pressure: " + p.Pressure + "\n"
	str = str + "Climate            : " + p.Climate + "\n"
	str = str + "Hydrosphere        : " + strconv.Itoa(p.Hydrosphere) + " %\n"
	return str
}

func NewStarSystem(template int) *StarSystem {
	system := StarSystem{}
	switch template {
	case 3:
		system.Stars = append(system.Stars, *NewStar("C"))
		system.Stars[0].description = "Old Red Giant Star"
		//planets :
	default:

	}

	return &system
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(randFloat(48, 16, 1))
	fmt.Println("random =", roll1dX(100, 0))
	fmt.Println("--------------------")
	fmt.Println(NewStar("Black Hole").toString())
	fmt.Println(NewStar("A").toString())
	fmt.Println(NewStar("F").toString())
	fmt.Println(NewStar("G").toString())

	syst := NewStarSystem(3)
	fmt.Println(syst.Stars[0].toString())
}
