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

func (mnr *miner) Mine() {
	r := rollXdY(2, 6)
	if r == 2 {
		mnr.mined = mnr.mined + 500.0*mnr.efficiency
		fmt.Println("Miner", mnr.id+1, "mined Precious Metals")
		fmt.Println("Miner", mnr.id+1, "mined total:", mnr.mined)
		return
	}
	if r < 7 {
		mnr.mined = mnr.mined + 100.0*mnr.efficiency
		fmt.Println("Miner", mnr.id+1, "mined Common Ore")
		fmt.Println("Miner", mnr.id+1, "mined total:", mnr.mined)
		return
	}
	if r < 12 {
		mnr.mined = mnr.mined + 200.0*mnr.efficiency
		fmt.Println("Miner", mnr.id+1, "mined Uncommon Ore")
		fmt.Println("Miner", mnr.id+1, "mined total:", mnr.mined)
		return
	}
	if r == 12 {
		mnr.mined = mnr.mined + 1000.0*mnr.efficiency
		fmt.Println("Miner", mnr.id+1, "mined Exotic")
		fmt.Println("Miner", mnr.id+1, "mined total:", mnr.mined)
		return
	}
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
	asteroidSize := float64(rollXdY(10, 6) * 10)
	for !asteroidFound {
		fmt.Println(timer.showTime(), " Ищем месторождение...")
		lsReserveStr, _ := lifeSupportStatus(minerCrew, crewMax, timer)
		fmt.Println(lsReserveStr)
		r := rollXdY(2, 6)
		searchMod := (mSkillInt - 1) * 2
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
			asteroidType = foundType(searchResult)
		}
	}
	fmt.Println(timer.showTime())
	fmt.Println("Начинаем разработку нового астеройда...")
	fmt.Println("Тип:", asteroidType)
	fmt.Println("Запасы:", asteroidSize, "тон")

	var crew []miner
	for i := totalMiners; i < minerCrew; i++ {
		crew = append(crew, *NewMiner(mSkillInt, i))
		fmt.Println("Miner Created:", crew[i])

	}

	fmt.Println(timer.showTime(), lifeSupportReserve)

	for i := range crew {
		crew[i].Mine()
	}
}
