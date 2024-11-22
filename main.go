package main

import (
	"fmt"
	"math/rand"
	"time"
)

type LaBaG struct {
	//遊玩次數
	Times  int
	Played int

	//分數
	Score       int
	MarginScore int

	//內部判斷變數
	Rams   [3]int
	Prizes [3]string
	//機率與分數字典
	Rates     [6]int
	ScoreDict map[string][6]int
}

func (Game *LaBaG) result() {
	Game.Score += Game.MarginScore
	fmt.Println("第", Game.Played, "次")
	fmt.Println(Game.Prizes[0], "|", Game.Prizes[1], "|", Game.Prizes[2])
	fmt.Println("+", Game.MarginScore)
	fmt.Println("目前分數：", Game.Score)
	Game.MarginScore = 0
}

func (Game *LaBaG) GameRunning() bool {
	if Game.Played < Game.Times {
		return true
	} else {
		return false
	}
}

func (Game *LaBaG) ramdon_value() {
	//隨機數
	rand.Seed(time.Now().UnixNano())
	Game.Rams[0], Game.Rams[1], Game.Rams[2] = rand.Intn(99)+1, rand.Intn(99)+1, rand.Intn(99)+1
}

func acc_rate(Rates [6]int) []int {
	//累加機率成機率區間
	var result_rate []int
	acc := 0
	for _, Rate := range Rates {
		acc += Rate
		result_rate = append(result_rate, acc)
	}
	return result_rate
}

func (Game *LaBaG) classify_prize() {
	//分類
	rate_range := acc_rate(Game.Rates)
	fmt.Println("機率區間：", rate_range)

	for i := 0; i < 3; i++ {
		switch {
		case Game.Rams[i] <= rate_range[0]:
			Game.Prizes[i] = "A"
		case Game.Rams[i] <= rate_range[1]:
			Game.Prizes[i] = "B"
		case Game.Rams[i] <= rate_range[2]:
			Game.Prizes[i] = "C"
		case Game.Rams[i] <= rate_range[3]:
			Game.Prizes[i] = "D"
		case Game.Rams[i] <= rate_range[4]:
			Game.Prizes[i] = "E"
		case Game.Rams[i] <= rate_range[5]:
			Game.Prizes[i] = "F"
		}
	}
}

func get_score(Prize string, MarginScore int, arr [6]int) int {
	// 抓取從Array分數
	switch Prize {
	case "A":
		MarginScore += arr[0]
	case "B":
		MarginScore += arr[1]
	case "C":
		MarginScore += arr[2]
	case "D":
		MarginScore += arr[3]
	case "E":
		MarginScore += arr[4]
	case "F":
		MarginScore += arr[5]
	}
	return MarginScore
}

func (Game *LaBaG) calculate_score() {
	switch {
	case Game.Prizes[0] == Game.Prizes[1] && Game.Prizes[1] == Game.Prizes[2]: //三個相同
		Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same3"])

	case Game.Prizes[0] == Game.Prizes[1] || Game.Prizes[1] == Game.Prizes[2] || Game.Prizes[0] == Game.Prizes[2]: //兩個相同
		switch {
		case Game.Prizes[0] == Game.Prizes[1]:
			Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same2"])
			Game.MarginScore = get_score(Game.Prizes[2], Game.MarginScore, Game.ScoreDict["same1"])
		case Game.Prizes[1] == Game.Prizes[2]:
			Game.MarginScore = get_score(Game.Prizes[1], Game.MarginScore, Game.ScoreDict["same2"])
			Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same1"])
		case Game.Prizes[2] == Game.Prizes[0]:
			Game.MarginScore = get_score(Game.Prizes[2], Game.MarginScore, Game.ScoreDict["same2"])
			Game.MarginScore = get_score(Game.Prizes[1], Game.MarginScore, Game.ScoreDict["same1"])
		}
		Game.MarginScore = Game.MarginScore / 13 * 10

	case Game.Prizes[0] != Game.Prizes[1] && Game.Prizes[1] != Game.Prizes[2] && Game.Prizes[0] != Game.Prizes[2]: //三個皆不同
		Game.MarginScore = get_score(Game.Prizes[0], Game.MarginScore, Game.ScoreDict["same1"])
		Game.MarginScore = get_score(Game.Prizes[1], Game.MarginScore, Game.ScoreDict["same1"])
		Game.MarginScore = get_score(Game.Prizes[2], Game.MarginScore, Game.ScoreDict["same1"])

		Game.MarginScore /= 3
	}
}

func main() {

	Game := LaBaG{
		Times:       30,
		Played:      0,
		Score:       0,
		MarginScore: 0,
		Rams:        [3]int{0, 0, 0},
		Prizes:      [3]string{"", "", ""},
		Rates:       [6]int{36, 24, 17, 12, 8, 3},
		ScoreDict: map[string][6]int{
			"same3": {625, 1250, 2100, 2500, 10000, 20000},
			"same2": {350, 650, 1080, 1250, 5000, 10000},
			"same1": {150, 220, 380, 420, 1250, 2500},
		},
	}

	for Game.GameRunning() {
		var press string
		fmt.Println("請按ENTER")
		fmt.Scanln(&press)

		if press == "" {
			Game.ramdon_value()

			Game.classify_prize()
			Game.calculate_score()

			Game.Played++
			Game.result()
		} else {
			fmt.Println("請按ENTER")
		}

	}

	fmt.Printf("\n遊戲結束！\n最終分數為：%d", Game.Score)
}
