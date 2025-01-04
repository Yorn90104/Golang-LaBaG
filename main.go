package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	PMap = make(map[string]*P)
	Game *LaBaG
)

func init() {
	NewP("A", [3]int{625, 350, 150}, map[string]int{
		"Normal":   36,
		"SuperHHH": 19,
		"GreenWei": 36,
		"PiKaChu":  36,
	})
	NewP("B", [3]int{1250, 650, 220}, map[string]int{
		"Normal":   24,
		"SuperHHH": 5,
		"GreenWei": 24,
		"PiKaChu":  24,
	})
	NewP("C", [3]int{2100, 1080, 380}, map[string]int{
		"Normal":   17,
		"SuperHHH": 19,
		"GreenWei": 17,
		"PiKaChu":  17,
	})
	NewP("D", [3]int{2500, 1250, 420}, map[string]int{
		"Normal":   12,
		"SuperHHH": 19,
		"GreenWei": 12,
		"PiKaChu":  12,
	})
	NewP("E", [3]int{10000, 5000, 1250}, map[string]int{
		"Normal":   8,
		"SuperHHH": 19,
		"GreenWei": 8,
		"PiKaChu":  8,
	})
	NewP("F", [3]int{20000, 10000, 2500}, map[string]int{
		"Normal":   3,
		"SuperHHH": 19,
		"GreenWei": 3,
		"PiKaChu":  3,
	})

	Game = &LaBaG{
		Times:       30,
		Played:      0,
		Score:       0,
		MarginScore: 0,
		ScoreTime:   1,
		ScoreTimeMap: map[string]int{
			"Normal":   1,
			"SuperHHH": 1,
			"GreenWei": 3,
			"PiKaChu":  1,
		},
		Ps:         [3]*P{},
		seq:        [6]string{"A", "B", "C", "D", "E", "F"},
		SuperHHH:   false,
		SuperRate:  15,
		SuperNum:   0,
		SuperTimes: 0,
		GreenWei:   false,
		GreenRate:  35,
		GreenNum:   0,
		GreenTimes: 0,
		GssNum:     0,
		PiKaChu:    false,
		KaChuTimes: 0,
	}
	rand.Seed(time.Now().UnixNano())
}

type P struct {
	Code    string
	Scores  [3]int
	RateMap map[string]int
}

func NewP(code string, scores [3]int, ratemap map[string]int) {
	// 創建新的 P 並放進 PMap 裡面
	// 如果 ratemap 為 nil，則初始化為預設值
	if ratemap == nil {
		ratemap = map[string]int{"Normal": 0}
	}

	p := &P{
		Code:    code,
		Scores:  scores,
		RateMap: ratemap,
	}

	PMap[p.Code] = p

}

type LaBaGInternal interface {
	// 啦八機內部有的方法
	Reset()
	Logic()
	GameRunning() bool
	NodMod() string
	Random()
	CalculateScore()
	Result()
	GameOver()
	JudgeMod()
}

type LaBaG struct {
	Times  int
	Played int

	Score        int
	MarginScore  int
	ScoreTime    int
	ScoreTimeMap map[string]int

	Ps  [3]*P
	seq [6]string

	SuperHHH   bool
	SuperRate  int
	SuperNum   int
	SuperTimes int

	GreenWei   bool
	GreenRate  int
	GreenNum   int
	GreenTimes int
	GssNum     int

	PiKaChu    bool
	KaChuTimes int
}

func (self *LaBaG) Reset() {
	self.Played = 0
	self.Score = 0
	self.MarginScore = 0
	self.ScoreTime = 0

	self.SuperHHH = false
	self.SuperNum = 0
	self.SuperTimes = 0

	self.GreenWei = false
	self.GreenNum = 0
	self.GreenTimes = 0
	self.GssNum = 0

	self.PiKaChu = false
	self.KaChuTimes = 0
}

func (self *LaBaG) Logic() {
	Game.Random()
	Game.CalculateScore()
	Game.Result()
	Game.JudgeMod()
}

func (self *LaBaG) GameRunning() bool {
	return self.Played < self.Times
}

func (self *LaBaG) NowMod() string {
	switch true {
	case self.SuperHHH:
		return "SuperHHH"
	case self.GreenWei:
		return "GreenWei"
	case self.PiKaChu:
		return "PiKaChu"
	default:
		return "Normal"
	}
}

func (self *LaBaG) Random() {
	// 產生隨機數 依據隨機數更換 Ps 中的 P
	RandNums := [3]int{rand.Intn(99) + 1, rand.Intn(99) + 1, rand.Intn(99) + 1}
	self.SuperNum = rand.Intn(99) + 1
	self.GreenNum = rand.Intn(99) + 1

	fmt.Println("P 隨機數:", RandNums[0], RandNums[1], RandNums[2])
	fmt.Println("超級阿禾隨機數:", self.SuperNum)
	fmt.Println("綠光阿瑋隨機數:", self.GreenNum)

	// 累積機率
	AccRate := func() []int {
		acc := 0
		res := []int{}
		for _, s := range self.seq {
			acc += PMap[s].RateMap[self.NowMod()]
			res = append(res, acc)
		}
		return res
	}

	RateRange := AccRate()
	fmt.Println("機率區間:", RateRange)

	for i := 0; i < 3; i++ {
		for idx, rate := range RateRange {
			if RandNums[i] <= rate {
				self.Ps[i] = PMap[self.seq[idx]]
				break
			}
		}
	}

	// 增加咖波累積數
	for _, p := range self.Ps {
		if p.Code == "A" && self.GssNum < 20 {
			self.GssNum += 1
		}
	}
	fmt.Println("咖波累積數:", self.GssNum)
}

func (self *LaBaG) CalculateScore() {
	// 計算分數
	// 根據 typ 的情況增加 MarginScore 的分數
	MarginAdd := func(p *P, typ int) {
		self.MarginScore += p.Scores[typ]
	}

	self.ScoreTime = self.ScoreTimeMap[self.NowMod()]
	fmt.Println("加分倍數:", self.ScoreTime)

	switch {
	case self.Ps[0] == self.Ps[1] && self.Ps[1] == self.Ps[2]:
		MarginAdd(self.Ps[0], 0)
	case self.Ps[0] == self.Ps[1] || self.Ps[1] == self.Ps[2] || self.Ps[2] == self.Ps[0]:
		switch {
		case self.Ps[0] == self.Ps[1]:
			MarginAdd(self.Ps[0], 1)
			MarginAdd(self.Ps[2], 2)
		case self.Ps[1] == self.Ps[2]:
			MarginAdd(self.Ps[1], 1)
			MarginAdd(self.Ps[0], 2)
		case self.Ps[2] == self.Ps[0]:
			MarginAdd(self.Ps[2], 1)
			MarginAdd(self.Ps[1], 2)
		}
		self.MarginScore = self.MarginScore * self.ScoreTime * 10 / 13
	case self.Ps[0] != self.Ps[1] && self.Ps[1] != self.Ps[2] && self.Ps[2] != self.Ps[0]:
		MarginAdd(self.Ps[0], 2)
		MarginAdd(self.Ps[1], 2)
		MarginAdd(self.Ps[2], 2)
		self.MarginScore = self.MarginScore * self.ScoreTime / 3
	}
}

func (self *LaBaG) Result() {
	// 印出結果
	self.Played++
	self.Score += self.MarginScore
	fmt.Println()
	fmt.Println(self.Ps[0].Code, "|", self.Ps[1].Code, "|", self.Ps[2].Code)
	fmt.Println("+", self.MarginScore)
	fmt.Println("目前分數：", self.Score)
	fmt.Println("剩餘次數：", self.Times-self.Played)
	self.MarginScore = 0
}

func (self *LaBaG) GameOver() {
	fmt.Println()
	fmt.Println("遊戲已結束，最終分數為：", self.Score)
}

func (self *LaBaG) JudgeMod() {
	// 判斷模式
	AnyP := func(cond func(*P) bool) bool {
		for _, p := range self.Ps {
			if cond(p) {
				return true
			}
		}
		return false
	}

	AllP := func(cond func(*P) bool) bool {
		for _, p := range self.Ps {
			if !cond(p) {
				return false
			}
		}
		return true
	}

	if !self.GameRunning() {
		// 關掉其他模式
		self.SuperHHH = false
		self.GreenWei = false

		// 判斷皮卡丘充電
		if AnyP(func(p *P) bool { return p.Code == "E" }) {
			self.PiKaChu = true
			self.Played -= 5
			self.KaChuTimes += 1
			fmt.Println("皮卡丘為你充電")
			fmt.Printf("已觸發 %d 次皮卡丘充電\n", self.KaChuTimes)
		} else {
			self.PiKaChu = false
		}
		return
	}

	switch self.NowMod() {
	case "Normal", "PiKaChu":
		// 判斷超級阿禾
		AnyB := AnyP(func(p *P) bool { return p.Code == "B" })
		if self.SuperNum <= self.SuperRate && AnyB {
			self.SuperHHH = true
			self.SuperTimes += 6
			fmt.Println("超級阿禾出現")
			if self.PiKaChu {
				self.PiKaChu = false
			}

			// 超級阿禾加倍
			if AllP(func(p *P) bool { return p.Code == "B" }) {
				double_score := self.Score * self.ScoreTime / 2
				self.Score += double_score
				fmt.Println("超級阿禾加倍分:", double_score)
			}
		}

		// 判斷綠光阿瑋
		AllA := AllP(func(p *P) bool { return p.Code == "A" })
		if self.GreenNum <= self.GreenRate && AllA {
			self.GreenWei = true
			self.GreenTimes += 2
			fmt.Println("綠光阿瑋出現")
			if self.PiKaChu {
				self.PiKaChu = false
			}
		} else if self.GssNum >= 20 { // 咖波累積數達到20
			self.GreenWei = true
			self.GreenTimes += 2
			self.GssNum = 0
			fmt.Println("綠光阿瑋出現")
			if self.PiKaChu {
				self.PiKaChu = false
			}
		}
		return
	case "SuperHHH":
		self.SuperTimes -= 1
		if AllP(func(p *P) bool { return p.Code == "B" }) {
			self.SuperTimes += 2
			fmt.Println("全阿禾，次數不消耗且+1！")
		}
		fmt.Printf("超級阿禾剩餘次數: %d 次\n", self.SuperTimes)
		if self.SuperTimes <= 0 { // 超級阿禾次數用完
			self.SuperHHH = false
			self.JudgeMod() // 判斷是否可再進入特殊模式
		}
		return
	case "GreenWei":
		self.GreenTimes -= 1
		if AllP(func(p *P) bool { return p.Code == "A" }) {
			self.GreenTimes += 1
			fmt.Println("全咖波，次數不消耗！")
		}
		fmt.Printf("綠光阿瑋剩餘次數: %d 次\n", self.GreenTimes)
		if self.GreenTimes <= 0 { // 綠光阿瑋次數用完
			self.GreenWei = false
			self.JudgeMod() // 判斷是否可再進入特殊模式
		}
		return
	}
}

func main() {
	for Game.GameRunning() {
		var press string
		fmt.Println("請按ENTER")
		fmt.Scanln(&press)

		if press == "" {
			Game.Logic()
		}
	}
	Game.GameOver()
}
