package main

type AlienPresense struct {
	alienCiv          *Visitors
	sizeOfInvolvement string
	spacePort         string
	agenda            string
}

type Visitors struct {
	id         int
	xenoType   string
	techRating TechRating
	goverment  string
}

var VisitorMap map[int]*Visitors

func NewVisitors() *Visitors {
	if VisitorMap == nil {
		VisitorMap = make(map[int]*Visitors)
		nullVis := Visitors{}
		nullVis.xenoType = "TEST CASE"
		VisitorMap[len(VisitorMap)] = &nullVis

	}
	vis := Visitors{}
	xeno := roll1dX(6, 0)
	vis.xenoType = RandomLifeform()
	if xeno < 3 {
		vis.xenoType = "Human"
	}
	vis.SetRandomTechRating()
	vis.SetRandomGoverment()
	vis.id = len(VisitorMap) + 1
	VisitorMap[vis.id] = &vis
	return &vis
}

func VisitorByID(i int) *Visitors {
	if len(VisitorMap) < 1 {
		return nil
	}
	return VisitorMap[i]
}

func (vis *Visitors) SetTechRating(template int) {
	vis.techRating.SetTechRating(template)
}

func (vis *Visitors) SetRandomTechRating() {
	roll := roll1dX(6, 0)
	vis.techRating.SetTechRating(10 + roll)
}

func (vis *Visitors) SetRandomGoverment() {
	template := randInt(1, 6) - randInt(1, 6) + vis.techRating.rating
	vis.goverment = SetGovermetSystem(template)
}

func (vis *Visitors) toString() string {
	str := "Stellar Race Report:\n"
	str = str + "Xenotype: " + vis.xenoType + "\n"
	str = str + "Technology is at " + vis.techRating.Descrition() + "\n"
	str = str + "General goverment type: " + vis.goverment + "\n"
	return str
}
