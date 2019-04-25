package main

import (
	"fmt"
	"strconv"
)

var totalMiners int

type miner struct {
	id         int
	skill      int
	efficiency float64
	mined      float64
	spent      int
}

type timer struct {
	hours int
	days  int
}

func (t *timer) addHour() {
	t.hours++
	if t.hours > 23 {
		t.days++
		t.hours = 0
	}
}

func (t *timer) addDay() {
	t.days++
}

func (t *timer) showTime() string {
	return "Day: " + strconv.Itoa(t.days) + " Time: " + strconv.Itoa(t.hours) + ":00"
}

func NewTimer() *timer {
	t := &timer{}
	t.days = 1
	return t
}

func NewMiner(skill, id int) *miner {
	mnr := &miner{}
	mnr.id = id
	mnr.skill = skill
	mnr.efficiency = 1.0
	return mnr
}

func (mnr *miner) Crisis() bool {
	r := rollXdY(2, 6)
	r = r + (mnr.skill * 2)
	if r > 7 {
		return false
	}
	mnr.efficiency = mnr.efficiency - 0.05
	return true
}

func (mnr *miner) Mine(ast *ateroidData) {
	mnr.efficiency = mnr.efficiency + 0.1
	if mnr.efficiency > 1.0 {
		mnr.efficiency = 1.0
	}
	if ast.asteroidYeld < 0 {
		return
	}
	if mnr.efficiency < 0.001 {
		fmt.Println("Miner", mnr.id+1, "в лазарете и не способен копать", mnr.efficiency)
		return
	}

	r := rollXdY(2, 6) + (mnr.skill)
	if r < 6 {
		mnr.efficiency = mnr.efficiency - float64(rollXdY(1, 4))
		mnr.mined = mnr.mined + 0.0*mnr.efficiency

		fmt.Println("Miner", mnr.id+1, "mined Nothing")
		fmt.Println("Roll Crisis-----------------------------------------------CRISIS!!!")
		return
	}
	if r < 10 {
		mnr.mined = mnr.mined + 1.0*mnr.efficiency

		fmt.Println("Miner", mnr.id+1, "mined 1 Ton")
		fmt.Println("Miner", mnr.id+1, "mined total:", mnr.mined)
		ast.asteroidYeld = ast.asteroidYeld - 1.0*mnr.efficiency
		return
	}
	if r < 14 {
		mnr.mined = mnr.mined + 1.0*mnr.efficiency

		fmt.Println("Miner", mnr.id+1, "mined 1.0 Tons")
		fmt.Println("Miner", mnr.id+1, "mined total:", mnr.mined)
		ast.asteroidYeld = ast.asteroidYeld - 1.0*mnr.efficiency
		return
	}
	mnr.mined = mnr.mined + 1.0*mnr.efficiency
	fmt.Println("Miner", mnr.id+1, "mined Exotic-----------------------------------------------EXOTIC!!!")
	fmt.Println("Miner", mnr.id+1, "mined total:", mnr.mined)
	ast.asteroidYeld = ast.asteroidYeld - 1.0*mnr.efficiency
	return

}

func foundType(index int) string {
	if index < 6 {
		return "Рудный Шлак"
	}
	if index < 8 {
		return "Бедная Руда"
	}
	if index < 10 {
		return "Обычная Руда"
	}
	if index < 12 {
		return "Богатая руда"
	}
	if index < 14 {
		return "Самородная жила"
	}
	return "Редкие Изотопы"
}

func orePrice(ore string) float64 {
	if ore == "Рудный Шлак" {
		return 50
	}
	if ore == "Бедная Руда" {
		return 100
	}
	if ore == "Обычная Руда" {
		return 200
	}
	if ore == "Богатая руда" {
		return 500
	}
	if ore == "Самородная жила" {
		return 1000
	}
	return 2500
}

func lifeSupportStatus(currentCrew, crewMax int, timer *timer) (string, int) {
	totalReserve := crewMax * 60
	reserveSpent := currentCrew * (timer.days - 1)
	reserveLeft := totalReserve - reserveSpent

	return "припасов осталось на " + strconv.Itoa(reserveLeft/currentCrew) + " дней", reserveLeft
}

func main() {
	randomSeed()
	timer := NewTimer()
	mSkillInt, mSkill := TakeOptions("Какое качество команды?", "Необученный сброд (10 кредитов в день)", "Умелые матросы (30 кредитов в день)", "Эксперты (100 кредитов в день)", "Мастера (300 кредитов в день)", "Гении космонавтики (1000 кредитов в день)")
	minerCrew := InputInt("Сколько людей в команде? ")
	crewMax := InputInt("Каков максимум экипажа для этого корабля? ")
	lifeSupportReserve := crewMax * 60
	spikeDrive := InputInt("Какой уровень Спайк Драйва у корабля? ")
	oreLevelInt, oreLevel := TakeOptions("Какое месторождение ищем?", "Любое", "Бедная Руда", "Обычная Руда", "Богатая руда", "Самородная жила", "Редкие Изотопы")
	fmt.Println(spikeDrive, oreLevel, oreLevelInt, mSkill)
	asteroidFound := false
	asteroidType := ""
	asteroidSize := float64(rollXdY(10, 6) * rollXdY(1, 1))
	for !asteroidFound {
		fmt.Println(timer.showTime(), " Ищем месторождение...")
		lsReserveStr, lsRI := lifeSupportStatus(minerCrew, crewMax, timer)
		if lsRI < minerCrew*3 {
			break
		}
		fmt.Println(lsReserveStr)
		r := rollXdY(2, 6)
		searchMod := mSkillInt
		fmt.Println("Бросок: ", r, searchMod)
		searchResult := r + searchMod
		searchTarget := 2 + (oreLevelInt * 2)
		fmt.Println("Найден астеройд. Основной состав:", foundType(searchResult))
		h := 48 / spikeDrive
		for i := 0; i < h; i++ {
			timer.addHour()
		}
		if searchResult > searchTarget-1 {
			asteroidFound = true

		}
		asteroidType = foundType(searchResult)
	}
	fmt.Println(timer.showTime())
	fmt.Println("Начинаем разработку нового астеройда...")
	fmt.Println("Тип:", asteroidType)
	fmt.Println("Запасы:", asteroidSize, "тон")
	ast := NewAsteriod(asteroidSize, asteroidType)
	var crew []miner
	for i := totalMiners; i < minerCrew; i++ {
		crew = append(crew, *NewMiner(mSkillInt, i))
		fmt.Println("Miner Created:", crew[i])

	}

	fmt.Println(timer.showTime(), lifeSupportReserve)
	yeld := ast.asteroidYeld

	for yeld > 0.0 {
		ldR, lsI := lifeSupportStatus(minerCrew, crewMax, timer)
		if lsI < minerCrew*3 {
			fmt.Println(timer.showTime(), ldR, lsI)
			fmt.Println("Корабль возвращается...")
			break
		}
		for i := range crew {
			crew[i].Mine(ast)
			yeld = ast.asteroidYeld
			//fmt.Println(ast)
		}
		timer.addDay()

		fmt.Println(timer.showTime(), ldR, lsI)
	}
	fmt.Println("ИТОГО:")
	minedTotal := 0.0
	for i := range crew {
		minedTotal = minedTotal + crew[i].mined
	}
	cost := minedTotal * orePrice(asteroidType)
	fmt.Println("накопано на:", int(cost), "всего", int(minedTotal), "тон")
	zp := timer.days * minerCrew * payGrade(mSkillInt)
	fmt.Println("затраты на зарплату:", zp)
	ls := timer.days * minerCrew * 20
	fmt.Println("затраты на жизнеобеспечение:", ls)
	fmt.Println("---------------")
	fmt.Println("Возможная прибыль:", int(cost)-zp-ls)

}

func payGrade(i int) int {
	switch i {
	case 0:
		return 10
	case 1:
		return 30
	case 2:
		return 100
	case 3:
		return 300
	case 4:
		return 1000
	}
	return 3000
}

type ateroidData struct {
	asteroidYeld float64
	asteroidType string
}

func NewAsteriod(size float64, asType string) *ateroidData {
	ast := &ateroidData{}
	ast.asteroidType = asType
	ast.asteroidYeld = size
	return ast
}
