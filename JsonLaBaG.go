package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	PMap = make(map[string]*P)
	Game *JsonLaBaG
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

	Game = &JsonLaBaG{
		AllData:     map[string]map[string]int{},
		OneData:     map[string]int{},
		DataIndex:   0,
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

type JsonLaBaGInternal interface {
	// 啦八機內部有的方法
	Reset()
	Logic()
	GameRunning() bool
	NodMod() string
	Random()
	CalculateScore()
	Result()
	JudgeMod()
}

type JsonLaBaG struct {
	AllData   map[string]map[string]int
	OneData   map[string]int
	DataIndex int

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

func (self *JsonLaBaG) Reset() {
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

func (self *JsonLaBaG) Logic() {
	self.Reset()

	self.DataIndex = 0
	// 清空 AllData
	for key := range self.AllData {
		delete(self.AllData, key)
	}

	for self.GameRunning() {
		// 清空 OneData
		for key := range self.OneData {
			delete(self.OneData, key)
		}

		self.Random()
		self.CalculateScore()
		self.Result()
		self.JudgeMod()
	}
}

func (self *JsonLaBaG) GameRunning() bool {
	return self.Played < self.Times
}

func (self *JsonLaBaG) NowMod() string {
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

func (self *JsonLaBaG) Random() {
	// 產生隨機數 依據隨機數更換 Ps 中的 P
	RandNums := [3]int{rand.Intn(99) + 1, rand.Intn(99) + 1, rand.Intn(99) + 1}
	self.SuperNum = rand.Intn(99) + 1
	self.GreenNum = rand.Intn(99) + 1

	for i := 0; i < 3; i++ {
		self.OneData[fmt.Sprintf("RandNums[%d]", i)] = RandNums[i]
	}
	self.OneData["SuperHHH"] = self.SuperNum
	self.OneData["GreenWei"] = self.GreenNum

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
}

func (self *JsonLaBaG) CalculateScore() {
	// 計算分數
	// 根據 typ 的情況增加 MarginScore 的分數
	MarginAdd := func(p *P, typ int) {
		self.MarginScore += p.Scores[typ]
	}

	self.ScoreTime = self.ScoreTimeMap[self.NowMod()]

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

func (self *JsonLaBaG) Result() {
	// 結果
	self.DataIndex++
	self.Played++
	self.Score += self.MarginScore
	self.MarginScore = 0
	self.AllData[fmt.Sprintf("%d", self.DataIndex)] = self.OneData
}

func (self *JsonLaBaG) JudgeMod() {
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
			if self.PiKaChu {
				self.PiKaChu = false
			}

			// 超級阿禾加倍
			if AllP(func(p *P) bool { return p.Code == "B" }) {
				double_score := self.Score * self.ScoreTime / 2
				self.Score += double_score
			}
		}

		// 判斷綠光阿瑋
		AllA := AllP(func(p *P) bool { return p.Code == "A" })
		if self.GreenNum <= self.GreenRate && AllA {
			self.GreenWei = true
			self.GreenTimes += 2
			if self.PiKaChu {
				self.PiKaChu = false
			}
		} else if self.GssNum >= 20 { // 咖波累積數達到20
			self.GreenWei = true
			self.GreenTimes += 2
			self.GssNum = 0
			if self.PiKaChu {
				self.PiKaChu = false
			}
		}
		return
	case "SuperHHH":
		self.SuperTimes -= 1
		if AllP(func(p *P) bool { return p.Code == "B" }) {
			self.SuperTimes += 2
		}
		if self.SuperTimes <= 0 { // 超級阿禾次數用完
			self.SuperHHH = false
			self.JudgeMod() // 判斷是否可再進入特殊模式
		}
		return
	case "GreenWei":
		self.GreenTimes -= 1
		if AllP(func(p *P) bool { return p.Code == "A" }) {
			self.GreenTimes += 1
		}
		if self.GreenTimes <= 0 { // 綠光阿瑋次數用完
			self.GreenWei = false
			self.JudgeMod() // 判斷是否可再進入特殊模式
		}
		return
	}
}

func InputTarget() int {
	for {
		fmt.Println("請輸入目標分數: ")
		var input string
		fmt.Scanln(&input)
		Target, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("請輸入有效的數字:", err)
			continue
		}

		if Target > 0 {
			return Target
		} else {
			fmt.Println("目標分數必須大於 0")
		}
	}
}

func main() {

	Target := InputTarget()

	i := 0
	recent_max, recent_total := 0, 0

	for {
		Game.Logic()

		i++
		recent_total = (recent_total + Game.Score) % 1000000000000000

		if Game.Score > recent_max {
			recent_max = Game.Score
		}
		fmt.Printf("第%d次 分數：%8d【目前最大值：%d】【目前平均值：%.2f】\n", i, Game.Score, recent_max, float64(recent_total)/float64(i))

		if Game.Score >= Target {
			break
		}
	}

	// 確保目錄存在
	output_dir := "C:\\JsonLaBaG\\"
	err := os.MkdirAll(output_dir, os.ModePerm) // os.ModePerm 是一個常數，代表“所有權限”，即 0777，表示目錄擁有讀、寫、執行權限。
	if err != nil {
		fmt.Println("創建目錄失敗:", err)
		return
	}

	timestamp := time.Now().Format("20060102") // YYYYMMDD

	// 建立 JSON 文件
	file, err := os.Create(fmt.Sprintf("%s%d_%s.json", output_dir, Game.Score, timestamp))
	if err != nil {
		fmt.Println("創建文件失敗:", err)
		return
	}
	defer file.Close() // 最後將文件關閉

	// 將 Game.AllData 寫入 JSON 文件
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // 設置 JSON 輸出的縮排格式
	err = encoder.Encode(Game.AllData)
	if err != nil {
		fmt.Println("寫入 JSON 失敗:", err)
		return
	}

	fmt.Println("文件成功寫入")
}
