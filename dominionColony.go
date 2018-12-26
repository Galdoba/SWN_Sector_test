package main

type TechRating struct {
	rating      int
	description string
}

func NewTechRating() *TechRating {
	tr := TechRating{}
	return &tr
}

func (tr *TechRating) Rating() int {
	return tr.rating
}

func (tr *TechRating) Descrition() string {
	return tr.description

}

func (tr *TechRating) SetTechRating(r int) {
	tr.rating = r
	switch tr.rating {
	case 0:
		tr.description = "TR 0: Animalistic/Feral State"
	case 1:
		tr.description = "TR 1: Stone Age (Fire)"
	case 2:
		tr.description = "TR 2: Bronze Age (Agriculture)"
	case 3:
		tr.description = "TR 3: Iron Age (Metalworking)"
	case 4:
		tr.description = "TR 4: Medieval Age (Architecture)"
	case 5:
		tr.description = "TR 5: Age of Sail (Gunpowder)"
	case 6:
		tr.description = "TR 6: Industrial Age (Steam Engines)"
	case 7:
		tr.description = "TR 7: Electrical Age (Radio)"
	case 8:
		tr.description = "TR 8: Nuclear Age (Atomic Power)"
	case 9:
		tr.description = "TR 9: Digital Age (Electronics)"
	case 10:
		tr.description = "TR 10: Proto-Stellar Age (Gravity Manipulation)"
	case 11:
		tr.description = "TR 11: Interstellar Colonies - 1 Parsec Hyper Drive"
	case 12:
		tr.description = "TR 12: Artificial Intelligence - 2 Parsec Hyper Drive"
	case 13:
		tr.description = "TR 13: Terraforming - 3 Parsec Hyper Drive"
	case 14:
		tr.description = "TR 14: Artificial Life, Clones - 4 Parsec Hyper Drive"
	case 15:
		tr.description = "TR 15: Matter Creation - 5 Parsec Hyper Drive"
	case 16:
		tr.description = "TR 16: Immortality - Interstellar Teleportation"
	default:
		tr.description = "%$#@!Error: This should not happen!"
	}
}

func (col *Colony) SetRandomGoverment() {
	template := randInt(1, 6) - randInt(1, 6) + col.TRating.rating
	col.goverment = SetGovermetSystem(template)
}

func SetGovermetSystem(template int) string {
	switch template {
	case 1:
		return "City-States/Balkanization"
	case 2:
		return "Colony/Puppet State"
	case 3:
		return "Self-perpetuating Oligarchy"
	case 4:
		return "Theocracy"
	case 5:
		return "Feudal"
	case 6:
		return "Imperial"
	case 7:
		return "Ruling Caste"
	case 8:
		return "Representative Democracy"
	case 9:
		return "Civil Service Bureaucracy"
	case 10:
		return "Military Dictatorship"
	case 11:
		return "Socialistic"
	case 12:
		return "Participating Democracy"
	default:
		if template <= 0 {
			return "Tribal"
		}
		return "Techno-Rationalism"
	}
}

type Colony struct {
	wasFounded        bool
	TRating           *TechRating
	goverment         string
	origins           string
	dominantEthnicity string
}

func (p *Planet) NewColony() *Colony {
	col := Colony{}
	trRoll := rollXdY(2, 6) - 2
	col.TRating = NewTechRating()
	col.TRating.SetTechRating(trRoll)
	col.SetRandomGoverment()
	col.SetRandomOrigin()
	col.SetRandomEthnicity()

	return &col

}

func (col *Colony) SetRandomOrigin() {
	temp := rollXdY(2, 6)
	col.SetOrigin(temp)
}

func (col *Colony) SetOrigin(template int) {
	switch template {
	case 2:
		col.origins = "Top Secret Research Facility"
	case 3:
		col.origins = "Xeno-Archeological Dig"
	case 4:
		col.origins = "Prison Colony"
	case 5:
		col.origins = "Terraforming Operation"
	case 6:
		col.origins = "Agricultural"
	case 7:
		col.origins = "General Colonization"
	case 8:
		col.origins = "Mining"
	case 9:
		col.origins = "Military Base"
	case 10:
		col.origins = "Research Colony"
	case 11:
		col.origins = "Religious Separatists"
	case 12:
		col.origins = "Utopian Social Experiment"
	default:
		col.origins = "Unknown"
	}

}

func (col *Colony) Origins() string {
	return col.origins
}

func (col *Colony) SetRandomEthnicity() {
	temp := rollXdY(2, 6)
	col.SetEthnicity(temp)
}

func (col *Colony) SetEthnicity(template int) {
	switch template {
	case 8:
		col.dominantEthnicity = "East Asian"
	case 9:
		col.dominantEthnicity = "South Asian"
	case 10:
		col.dominantEthnicity = "African"
	case 11:
		sl := []string{
			"Caucasian",
			"Hispanic",
			"Middle Eastern",
		}
		col.dominantEthnicity = getRandomFromSliceStr(sl)
	case 12:
		sl := []string{
			"Native American",
			"Pacific Islander",
		}
		col.dominantEthnicity = getRandomFromSliceStr(sl)
	default:
		col.dominantEthnicity = "None. There is a diverse population"
	}

}

func (col *Colony) toString() string {
	str := ""
	str = str + "Attention! Human population is present.\n"
	str = str + "Overal Technological Rating is " + col.TRating.Descrition() + ". "
	str = str + "Most common goverment system seems to be " + col.goverment + ". "
	str = str + "Colony origins is " + col.origins + ". "
	str = str + "Dominant Ethnicity " + col.dominantEthnicity + ".\n"
	return str
}
