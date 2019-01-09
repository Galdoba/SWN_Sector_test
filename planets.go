package main

import (
	"strconv"
)

const (
	climateFrigid    = "Frigid (< -50c)"
	climateCold      = "Cold (-50c to 0c)"
	climateTemperate = "Temperate (0c to 30c)"
	climateHot       = "Hot (30c to 80c)"
	climateRoasting  = "Roasting (80c+)"
)

type Planet struct {
	Name                       string
	Gravity                    float64
	Moons                      string
	Pressure                   string
	Climate                    string
	Hydrosphere                int
	Ecosystem                  string
	EModRating                 int
	Description                string
	DominantIndigenousLifeForm string
	IndigenousIntelligence     string
	colony                     *Colony
	alienCiv                   *Visitors
	alienAgenda                string
	nextStep                   string
	visitorContact             int
}

func NewGardenPlanet(template int) *Planet {
	p := Planet{}
	p.Moons = "DEBUG: planet template № " + strconv.Itoa(template)
	switch template {
	case 1:
		p.Gravity = 0.1
		p.Moons = "None"
		p.Climate = climateFrigid
		p.Hydrosphere = 0
		p.EModRating = 0
		p.nextStep = "H"
	case 2:
		p.Gravity = 0.2
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 22
		p.EModRating = 0
		p.nextStep = "H"
	case 3:
		p.Gravity = 0.3
		p.Moons = "None"
		p.Climate = climateRoasting
		p.Hydrosphere = 0
		p.EModRating = 0
		p.nextStep = "H"
	case 4:
		p.Gravity = 0.4
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 8
		p.EModRating = 0
		p.nextStep = "H"
	case 5:
		p.Gravity = 0.4
		p.Moons = "None"
		p.Climate = climateFrigid
		p.Hydrosphere = 0
		p.EModRating = 0
		p.nextStep = "H"
	case 6:
		p.Gravity = 0.5
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 31
		p.EModRating = 1
		p.nextStep = "G"
	case 7:
		p.Gravity = 0.5
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 51
		p.EModRating = 1
		p.nextStep = "G"
	case 8:
		p.Gravity = 0.5
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 40
		p.EModRating = 3
		p.nextStep = "F"
	case 9:
		p.Gravity = 0.5
		p.Moons = "None"
		p.Climate = climateFrigid
		p.Hydrosphere = 29
		p.EModRating = 1
		p.nextStep = "G"
	case 10:
		p.Gravity = 0.5
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 28
		p.EModRating = 3
		p.nextStep = "F"
	case 11:
		p.Gravity = 0.5
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 4
		p.EModRating = 1
		p.nextStep = "H"
	case 12:
		p.Gravity = 0.5
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 35
		p.EModRating = 2
		p.nextStep = "G"
	case 13:
		p.Gravity = 0.6
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 46
		p.EModRating = 1
		p.nextStep = "H"
	case 14:
		p.Gravity = 0.6
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 14
		p.EModRating = 2
		p.nextStep = "G"
	case 15:
		p.Gravity = 0.6
		p.Moons = "None"
		p.Climate = climateHot
		p.Hydrosphere = 48
		p.EModRating = 2
		p.nextStep = "G"
	case 16:
		p.Gravity = 0.6
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 53
		p.EModRating = 3
		p.nextStep = "F"
	case 17:
		p.Gravity = 0.6
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 24
		p.EModRating = 1
		p.nextStep = "H"
	case 18:
		p.Gravity = 0.6
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 62
		p.EModRating = 3
		p.nextStep = "F"
	case 19:
		p.Gravity = 0.7
		p.Moons = "None"
		p.Climate = climateHot
		p.Hydrosphere = 30
		p.EModRating = 2
		p.nextStep = "G"
	case 20:
		p.Gravity = 0.7
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 25
		p.EModRating = 3
		p.nextStep = "F"
	case 21:
		p.Gravity = 0.7
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 15
		p.EModRating = 2
		p.nextStep = "G"
	case 22:
		p.Gravity = 0.7
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 42
		p.EModRating = 3
		p.nextStep = "F"
	case 23:
		p.Gravity = 0.7
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 34
		p.EModRating = 1
		p.nextStep = "H"
	case 24:
		p.Gravity = 0.7
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 13
		p.EModRating = 2
		p.nextStep = "G"
	case 25:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 68
		p.EModRating = 3
		p.nextStep = "F"
	case 26:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 99
		p.EModRating = 4
		p.nextStep = "F"
	case 27:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 55
		p.EModRating = 4
		p.nextStep = "F"
	case 28:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateHot
		p.Hydrosphere = 22
		p.EModRating = 1
		p.nextStep = "H"
	case 29:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 66
		p.EModRating = 4
		p.nextStep = "F"
	case 30:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 41
		p.EModRating = 3
		p.nextStep = "F"
	case 31:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 65
		p.EModRating = 4
		p.nextStep = "F"
	case 32:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateCold
		p.Hydrosphere = 45
		p.EModRating = 3
		p.nextStep = "F"
	case 33:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateHot
		p.Hydrosphere = 31
		p.EModRating = 3
		p.nextStep = "F"
	case 34:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 66
		p.EModRating = 4
		p.nextStep = "F"
	case 35:
		p.Gravity = 0.8
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 73
		p.EModRating = 3
		p.nextStep = "F"
	case 36:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateHot
		p.Hydrosphere = 44
		p.EModRating = 3
		p.nextStep = "F"
	case 37:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 44
		p.EModRating = 4
		p.nextStep = "F"
	case 38:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 93
		p.EModRating = 3
		p.nextStep = "F"
	case 39:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 56
		p.EModRating = 4
		p.nextStep = "F"
	case 40:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 46
		p.EModRating = 3
		p.nextStep = "F"
	case 41:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 51
		p.EModRating = 4
		p.nextStep = "F"
	case 42:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 50
		p.EModRating = 4
		p.nextStep = "F"
	case 43:
		p.Gravity = 0.9
		p.Moons = "None"
		p.Climate = climateTemperate
		p.Hydrosphere = 56
		p.EModRating = 3
		p.nextStep = "F"
	case 44:
		p.Gravity = 0.9
		p.Moons = "1 small moon"
		p.Climate = climateCold
		p.Hydrosphere = 33
		p.EModRating = 4
		p.nextStep = "F"
	case 45:
		p.Gravity = 1.0
		p.Moons = "1 small moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 61
		p.EModRating = 4
		p.nextStep = "F"
	case 46:
		p.Gravity = 1.0
		p.Moons = "1 small moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 51
		p.EModRating = 5
		p.nextStep = "F"
	case 47:
		p.Gravity = 1.0
		p.Moons = "1 small moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 88
		p.EModRating = 5
		p.nextStep = "F"
	case 48:
		p.Gravity = 1.0
		p.Moons = "1 small moon"
		p.Climate = climateHot
		p.Hydrosphere = 39
		p.EModRating = 3
		p.nextStep = "F"
	case 49:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 63
		p.EModRating = 5
		p.nextStep = "F"
	case 50:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 50
		p.EModRating = 4
		p.nextStep = "F"
	case 51:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 55
		p.EModRating = 5
		p.nextStep = "F"
	case 52:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 81
		p.EModRating = 5
		p.nextStep = "F"
	case 53:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 89
		p.EModRating = 4
		p.nextStep = "F"
	case 54:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 83
		p.EModRating = 5
		p.nextStep = "F"
	case 55:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 75
		p.EModRating = 4
		p.nextStep = "F"
	case 56:
		p.Gravity = 1.0
		p.Moons = "1 moon"
		p.Climate = climateCold
		p.Hydrosphere = 38
		p.EModRating = 4
		p.nextStep = "F"
	case 57:
		p.Gravity = 1.1
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 97
		p.EModRating = 5
		p.nextStep = "F"
	case 58:
		p.Gravity = 1.1
		p.Moons = "1 moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 54
		p.EModRating = 4
		p.nextStep = "F"
	case 59:
		p.Gravity = 1.1
		p.Moons = "1 large moon"
		p.Climate = climateTemperate
		p.Hydrosphere = 53
		p.EModRating = 5
		p.nextStep = "F"
	case 60:
		p.Gravity = 1.1
		p.Moons = "1 large moon"
		p.Climate = climateHot
		p.Hydrosphere = 50
		p.EModRating = 2
		p.nextStep = "G"
	case 61:
		p.Gravity = 1.1
		p.Moons = "2 small moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 78
		p.EModRating = 5
		p.nextStep = "F"
	case 62:
		p.Gravity = 1.1
		p.Moons = "2 small moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 60
		p.EModRating = 5
		p.nextStep = "F"
	case 63:
		p.Gravity = 1.1
		p.Moons = "2 small moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 73
		p.EModRating = 4
		p.nextStep = "F"
	case 64:
		p.Gravity = 1.1
		p.Moons = "2 small moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 68
		p.EModRating = 5
		p.nextStep = "F"
	case 65:
		p.Gravity = 1.1
		p.Moons = "2 moons"
		p.Climate = climateHot
		p.Hydrosphere = 37
		p.EModRating = 3
		p.nextStep = "F"
	case 66:
		p.Gravity = 1.1
		p.Moons = "2 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 72
		p.EModRating = 5
		p.nextStep = "F"
	case 67:
		p.Gravity = 1.2
		p.Moons = "2 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 65
		p.EModRating = 5
		p.nextStep = "F"
	case 68:
		p.Gravity = 1.2
		p.Moons = "2 moons"
		p.Climate = climateCold
		p.Hydrosphere = 45
		p.EModRating = 3
		p.nextStep = "F"
	case 69:
		p.Gravity = 1.2
		p.Moons = "2 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 89
		p.EModRating = 5
		p.nextStep = "F"
	case 70:
		p.Gravity = 1.2
		p.Moons = "2 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 70
		p.EModRating = 4
		p.nextStep = "F"
	case 71:
		p.Gravity = 1.2
		p.Moons = "1 small, 1 large moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 55
		p.EModRating = 4
		p.nextStep = "F"
	case 72:
		p.Gravity = 1.2
		p.Moons = "1 small, 1 large moons"
		p.Climate = climateHot
		p.Hydrosphere = 50
		p.EModRating = 4
		p.nextStep = "F"
	case 73:
		p.Gravity = 1.2
		p.Moons = "1 small, 1 large moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 44
		p.EModRating = 3
		p.nextStep = "F"
	case 74:
		p.Gravity = 1.2
		p.Moons = "1 small, 1 large moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 86
		p.EModRating = 4
		p.nextStep = "F"
	case 75:
		p.Gravity = 1.2
		p.Moons = "3 small moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 65
		p.EModRating = 3
		p.nextStep = "F"
	case 76:
		p.Gravity = 1.2
		p.Moons = "3 small moons"
		p.Climate = climateHot
		p.Hydrosphere = 5
		p.EModRating = 2
		p.nextStep = "G"
	case 77:
		p.Gravity = 1.3
		p.Moons = "3 small moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 50
		p.EModRating = 3
		p.nextStep = "F"
	case 78:
		p.Gravity = 1.3
		p.Moons = "3 small moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 66
		p.EModRating = 3
		p.nextStep = "F"
	case 79:
		p.Gravity = 1.3
		p.Moons = "3 small moons"
		p.Climate = climateCold
		p.Hydrosphere = 38
		p.EModRating = 2
		p.nextStep = "G"
	case 80:
		p.Gravity = 1.3
		p.Moons = "3 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 20
		p.EModRating = 1
		p.nextStep = "H"
	case 81:
		p.Gravity = 1.3
		p.Moons = "3 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 66
		p.EModRating = 3
		p.nextStep = "F"
	case 82:
		p.Gravity = 1.3
		p.Moons = "3 moons"
		p.Climate = climateHot
		p.Hydrosphere = 24
		p.EModRating = 1
		p.nextStep = "H"
	case 83:
		p.Gravity = 1.4
		p.Moons = "3 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 34
		p.EModRating = 2
		p.nextStep = "G"
	case 84:
		p.Gravity = 1.4
		p.Moons = "3 moons"
		p.Climate = climateTemperate
		p.Hydrosphere = 36
		p.EModRating = 3
		p.nextStep = "F"
	case 85:
		p.Gravity = 1.4
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 41
		p.EModRating = 2
		p.nextStep = "G"
	case 86:
		p.Gravity = 1.4
		p.Moons = "Rings"
		p.Climate = climateHot
		p.Hydrosphere = 27
		p.EModRating = 2
		p.nextStep = "G"
	case 87:
		p.Gravity = 1.4
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 37
		p.EModRating = 3
		p.nextStep = "F"
	case 88:
		p.Gravity = 1.4
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 52
		p.EModRating = 3
		p.nextStep = "F"
	case 89:
		p.Gravity = 1.5
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 52
		p.EModRating = 3
		p.nextStep = "F"
	case 90:
		p.Gravity = 1.5
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 39
		p.EModRating = 2
		p.nextStep = "G"
	case 91:
		p.Gravity = 1.5
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 39
		p.EModRating = 3
		p.nextStep = "F"
	case 92:
		p.Gravity = 1.5
		p.Moons = "Rings"
		p.Climate = climateHot
		p.Hydrosphere = 47
		p.EModRating = 2
		p.nextStep = "G"
	case 93:
		p.Gravity = 1.5
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 49
		p.EModRating = 2
		p.nextStep = "G"
	case 94:
		p.Gravity = 1.5
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 55
		p.EModRating = 3
		p.nextStep = "F"
	case 95:
		p.Gravity = 1.5
		p.Moons = "Rings"
		p.Climate = climateRoasting
		p.Hydrosphere = 0
		p.EModRating = 0
		p.nextStep = "H"
	case 96:
		p.Gravity = 1.6
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 33
		p.EModRating = 3
		p.nextStep = "F"
	case 97:
		p.Gravity = 1.6
		p.Moons = "Rings"
		p.Climate = climateHot
		p.Hydrosphere = 36
		p.EModRating = 2
		p.nextStep = "G"
	case 98:
		p.Gravity = 1.7
		p.Moons = "Rings"
		p.Climate = climateRoasting
		p.Hydrosphere = 0
		p.EModRating = 0
		p.nextStep = "F"
	case 99:
		p.Gravity = 1.8
		p.Moons = "Rings"
		p.Climate = climateTemperate
		p.Hydrosphere = 39
		p.EModRating = 3
		p.nextStep = "F"
	case 100:
		p.Gravity = 2.0
		p.Moons = "Rings"
		p.Climate = climateHot
		p.Hydrosphere = 37
		p.EModRating = 1
		p.nextStep = "H"

	default:
	}
	p.Ecosystem = p.createEcoSystem(p.EModRating)
	p.Pressure = p.createAtmosphere(p.Gravity)
	if p.nextStep == "F" {
		p.DominantIndigenousLifeForm = RandomLifeform()
		p.getIndigenousIntelligence()
		p.nextStep = "G"
	}
	if p.nextStep == "G" {
		if roll1dX(12, p.EModRating) > 10 {
			p.colony = p.NewColony()
		}
		p.nextStep = "H"
	}
	if p.nextStep == "H" {
		var visitor *Visitors
		roll := roll1dX(12, 0)
		visID := -1
		p.visitorContact = -1
		if roll < 10 {
			p.visitorContact = -1
		}
		if roll > 9 && roll < 12 {
			if len(VisitorMap) < 1 {
				NewVisitors()
			}
			visID = roll1dX(len(VisitorMap), 0)
			for visID < 2 {
				visID = roll1dX(len(VisitorMap), 0)
			}
			visitor = VisitorByID(visID)
			p.visitorContact = visID // visitor.id
		}
		if roll > 11 {
			visitor = NewVisitors()
			p.visitorContact = visitor.id
		}

		// var vis *Visitors
		// do := 1
		// if roll > 9 {
		// 	do = 2
		// 	if len(VisitorMap) == 0 {
		// 		do = 3
		// 	}
		// }
		// if roll > 11 {
		// 	do = 3
		// }

		// switch do {
		// case 2:
		// 	for _, val := range VisitorMap {
		// 		vis = val
		// 		break
		// 	}
		// case 3:
		// 	vis = NewVisitors()
		// default:
		// }
		// vis.xenoType = "NONO"

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

func (p *Planet) createEcoSystem(emod int) string {
	switch emod {
	case 0:
		return "Lifeless"
	case 1:
		return "Sparce Life"
	case 2:
		return "Partial Ecosystem"
	case 3:
		return "Full Ecosystem"
	case 4:
		return "Extensive Ecosystem"
	case 5:
		return "Very Lush Ecosystem"
	default:
	}
	return "%Error: Unknown"
}

func (p *Planet) createAtmosphere(grav float64) string {
	if grav < 0.2 {
		return "Trace"
	}
	if grav < 0.5 {
		return "Very Thin"
	}
	if grav < 0.8 {
		return "Thin"
	}
	if grav < 1.3 {
		return "Standard"
	}
	if grav < 1.6 {
		return "Dense"
	}
	if grav < 2.1 {
		return "Very Dense"
	}
	return "%Error: Unknown"
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
	str = str + "Gravity    : " + FloatToString(p.Gravity, 1) + "g\n"
	str = str + "Moons      : " + p.Moons + "\n"
	if p.Climate != "" {
		str = str + "Atmosphere : " + p.Pressure + "\n"
		str = str + "Hydrosphere: " + strconv.Itoa(p.Hydrosphere) + " %\n"
		str = str + "Biosphere  : " + p.Ecosystem + "\n"
		str = str + "Climate    : " + p.Climate + "\n"

	}
	if p.DominantIndigenousLifeForm != "" {
		str = str + "Survey data indicates that indigenous dominant lifeform type is " + p.DominantIndigenousLifeForm + ". Ingenious intelligence is " + p.IndigenousIntelligence + "\n"
	}
	if p.colony != nil {
		str = str + p.colony.toString()
	}
	// fmt.Println(p.Name, "has visitor", p.visitorContact)
	if p.visitorContact != -1 && p.visitorContact != 0 {
		// fmt.Println(VisitorMap)
		// fmt.Println(VisitorByID(p.visitorContact))
		// fmt.Println(p.visitorContact)
		str = str + VisitorByID(p.visitorContact).toString()
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

func SetLifeform(template int) string {
	switch template {
	case 1:
		sl := []string{
			"Amoeba",
			"Jellyfish",
			"Worm",
		}
		return getRandomFromSliceStr(sl)
	case 2:
		sl := []string{
			"Insect",
			"Crustacean",
			"Arachnid",
		}
		return getRandomFromSliceStr(sl)
	case 3:
		sl := []string{
			"Fish",
			"Cephalopod",
			"Amphbian",
		}
		return getRandomFromSliceStr(sl)
	case 4:
		return "Mammal"
	case 5:
		return "Mammal"
	case 6:
		return "Mammal"
	case 7:
		sl := []string{
			"Reptile",
			"Bird",
		}
		return getRandomFromSliceStr(sl)
	case 8:
		sl := []string{
			"Reptile",
			"Bird",
		}
		return getRandomFromSliceStr(sl)
	case 9:
		sl := []string{
			"Plant",
			"Fungus",
		}
		return getRandomFromSliceStr(sl)
	default:
	}
	return "Unknown"
}

func RandomLifeform() string {
	r := roll1dX(10, 0)
	return SetLifeform(r)
}

func (p *Planet) getIndigenousIntelligence() {
	r := roll1dX(12, 0)
	if r > 9 {
		p.getSentientLifeformTechRating()
		return
	}

	p.IndigenousIntelligence = "not observable"
}

func (p *Planet) getSentientLifeformTechRating() {
	r := roll1dX(8, -1)
	if r < 0 {
		r = 0
	}
	switch r {
	case 0:
		p.IndigenousIntelligence = "TR 0: Animalistic/Feral State"
	case 1:
		p.IndigenousIntelligence = "TR 1: Stone Age (Fire)"
	case 2:
		p.IndigenousIntelligence = "TR 2: Bronze Age (Agriculture)"
	case 3:
		p.IndigenousIntelligence = "TR 3: Iron Age (Metalworking)"
	case 4:
		p.IndigenousIntelligence = "TR 4: Medieval Age (Architecture)"
	case 5:
		p.IndigenousIntelligence = "TR 5: Age of Sail (Gunpowder)"
	case 6:
		p.IndigenousIntelligence = "TR 6: Industrial Age (Steam Engines)"
	case 7:
		p.IndigenousIntelligence = "TR 7: Electrical Age (Radio)"
	default:
	}

}
