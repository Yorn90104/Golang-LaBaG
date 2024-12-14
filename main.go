package main

import (
	"fmt"
	"math/rand"
	"time"
)

type P struct {
	code  string
	score [3]int
}

type LaBaG struct {
	Times       int
	Played      int
	Score       int
	MarginScore int
	Ps          [6]P
	all_P       [3]P
}

var Game LaBaG

func init() {
	Game = LaBaG{
		Times:       30,
		Played:      0,
		Score:       0,
		MarginScore: 0,
		all_P:       [3]P{},
		Ps: [6]P{
			{code: "A", score: [3]int{625, 350, 150}},
			{code: "B", score: [3]int{1250, 650, 220}},
			{code: "C", score: [3]int{2100, 1080, 380}},
			{code: "D", score: [3]int{2500, 1250, 420}},
			{code: "E", score: [3]int{10000, 5000, 1250}},
			{code: "F", score: [3]int{20000, 10000, 2500}},
		},
	}
	rand.Seed(time.Now().UnixNano())
}

func (Game *LaBaG) GameRunning() bool {
	return Game.Played < Game.Times
}

func (Game *LaBaG) random_p() {
	//分類
	Rates := [6]int{36, 24, 17, 12, 8, 3} //機率
	acc_rate := func(Rates [6]int) []int {
		var result_rate []int
		acc := 0
		for _, Rate := range Rates {
			acc += Rate
			result_rate = append(result_rate, acc)
		}
		return result_rate
	}
	rate_range := acc_rate(Rates)
	fmt.Println("機率區間：", rate_range)

	Rams := [3]int{rand.Intn(99) + 1, rand.Intn(99) + 1, rand.Intn(99) + 1}

	for i := 0; i < 3; i++ {
		switch {
		case Rams[i] <= rate_range[0]:
			Game.all_P[i] = Game.Ps[0]
		case Rams[i] <= rate_range[1]:
			Game.all_P[i] = Game.Ps[1]
		case Rams[i] <= rate_range[2]:
			Game.all_P[i] = Game.Ps[2]
		case Rams[i] <= rate_range[3]:
			Game.all_P[i] = Game.Ps[3]
		case Rams[i] <= rate_range[4]:
			Game.all_P[i] = Game.Ps[4]
		case Rams[i] <= rate_range[5]:
			Game.all_P[i] = Game.Ps[5]
		}
	}
}
func (Game *LaBaG) get_score(p P, typ int) {
	//typ 得分類型
	switch p.code {
	case "A":
		Game.MarginScore += Game.Ps[0].score[typ]
	case "B":
		Game.MarginScore += Game.Ps[1].score[typ]
	case "C":
		Game.MarginScore += Game.Ps[2].score[typ]
	case "D":
		Game.MarginScore += Game.Ps[3].score[typ]
	case "E":
		Game.MarginScore += Game.Ps[4].score[typ]
	case "F":
		Game.MarginScore += Game.Ps[5].score[typ]
	}
}

func (Game *LaBaG) calculate_score() {

	switch {
	case Game.all_P[0] == Game.all_P[1] && Game.all_P[1] == Game.all_P[2]: //三個相同
		Game.get_score(Game.all_P[0], 0)
	case Game.all_P[0] == Game.all_P[1] || Game.all_P[1] == Game.all_P[2] || Game.all_P[0] == Game.all_P[2]: //兩個相同
		switch {
		case Game.all_P[0] == Game.all_P[1]:
			Game.get_score(Game.all_P[0], 1)
			Game.get_score(Game.all_P[2], 2)
		case Game.all_P[1] == Game.all_P[2]:
			Game.get_score(Game.all_P[1], 1)
			Game.get_score(Game.all_P[0], 2)
		case Game.all_P[2] == Game.all_P[0]:
			Game.get_score(Game.all_P[2], 1)
			Game.get_score(Game.all_P[1], 2)
		}
		Game.MarginScore = Game.MarginScore * 10 / 13
	case Game.all_P[0] != Game.all_P[1] && Game.all_P[1] != Game.all_P[2] && Game.all_P[0] != Game.all_P[2]: //三個皆不同
		Game.get_score(Game.all_P[0], 2)
		Game.get_score(Game.all_P[1], 2)
		Game.get_score(Game.all_P[2], 2)
		Game.MarginScore /= 3
	}
}

func (Game *LaBaG) result() {
	Game.Played++
	Game.Score += Game.MarginScore
	fmt.Println("第", Game.Played, "次")
	fmt.Println(Game.all_P[0].code, "|", Game.all_P[1].code, "|", Game.all_P[2].code)
	fmt.Println("+", Game.MarginScore)
	fmt.Println("目前分數：", Game.Score)
	Game.MarginScore = 0
}

func main() {
	for Game.GameRunning() {
		var press string
		fmt.Println("請按ENTER")
		fmt.Scanln(&press)

		if press == "" {

			Game.random_p()
			Game.calculate_score()

			Game.result()
		}
	}
}
