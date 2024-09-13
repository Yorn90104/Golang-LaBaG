package main

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	//歸屬
	p1 = ""
	p2 = ""
	p3 = ""
	//隨機數
	ram1 = 0
	ram2 = 0
	ram3 = 0
	//分數
	score = 0
	add   = 0
	//次數
	times = 30
	ed    = 0
	press string
)

func test() {
	fmt.Println("第", ed, "次")
	fmt.Println(p1, "|", p2, "|", p3)
	fmt.Println("+", add)
	fmt.Println("目前分數：", score)
}

func P(ram int) (p string) {
	//判斷歸屬
	switch {
	case ram <= 36:
		p = "A"
	case ram <= 60:
		p = "B"
	case ram <= 77:
		p = "C"
	case ram <= 89:
		p = "D"
	case ram <= 97:
		p = "E"
	case ram <= 100:
		p = "F"
	}
	return p
}

func IF() (cond string, typ int) {
	//判斷情況
	switch {
	//三個相同
	case p1 == p2 && p1 == p3:
		cond = "3"
		typ = 0
	//兩個相同
	case p1 == p2 || p2 == p3 || p1 == p3:
		switch {
		case p1 == p2:
			typ = 3
		case p2 == p3:
			typ = 1
		case p1 == p3:
			typ = 2
		}
		cond = "2"
	//三個不同
	case p1 != p2 && p2 != p3 && p1 != p3:
		cond = "1"
		typ = 0
	}
	return cond, typ
}

func ADD(cond string, typ int) (a int) {
	//計算分數
	switch cond {
	case "3":
		switch p1 {
		case "A":
			a = a + 200
		case "B":
			a = a + 600
		case "C":
			a = a + 1600
		case "D":
			a = a + 1800
		case "E":
			a = a + 10000
		case "F":
			a = a + 20000
		}
	case "2":
		switch typ {
		case 3: // 1 & 2
			switch p1 {
			case "A":
				a = a + 100
			case "B":
				a = a + 170
			case "C":
				a = a + 780
			case "D":
				a = a + 870
			case "E":
				a = a + 5000
			case "F":
				a = a + 10000
			}
			switch p3 {
			case "A":
				a = a + 30
			case "B":
				a = a + 50
			case "C":
				a = a + 250
			case "D":
				a = a + 290
			case "E":
				a = a + 1200
			case "F":
				a = a + 2500
			}

			a = a / 13 * 10

		case 1: // 2 & 3
			switch p2 {
			case "A":
				a = a + 100
			case "B":
				a = a + 170
			case "C":
				a = a + 780
			case "D":
				a = a + 870
			case "E":
				a = a + 5000
			case "F":
				a = a + 10000
			}
			switch p1 {
			case "A":
				a = a + 30
			case "B":
				a = a + 50
			case "C":
				a = a + 250
			case "D":
				a = a + 290
			case "E":
				a = a + 1200
			case "F":
				a = a + 2500
			}

			a = a / 13 * 10

		case 2: // 1 & 3
			switch p1 {
			case "A":
				a = a + 100
			case "B":
				a = a + 170
			case "C":
				a = a + 780
			case "D":
				a = a + 870
			case "E":
				a = a + 5000
			case "F":
				a = a + 10000
			}
			switch p2 {
			case "A":
				a = a + 30
			case "B":
				a = a + 50
			case "C":
				a = a + 250
			case "D":
				a = a + 290
			case "E":
				a = a + 1200
			case "F":
				a = a + 2500
			}

			a = a / 13 * 10
		}
	case "1":
		//1
		switch p1 {
		case "A":
			a = a + 30
		case "B":
			a = a + 50
		case "C":
			a = a + 250
		case "D":
			a = a + 290
		case "E":
			a = a + 1200
		case "F":
			a = a + 2500
		}

		//2
		switch p2 {
		case "A":
			a = a + 30
		case "B":
			a = a + 50
		case "C":
			a = a + 250
		case "D":
			a = a + 290
		case "E":
			a = a + 1200
		case "F":
			a = a + 2500
		}

		//3
		switch p3 {
		case "A":
			a = a + 30
		case "B":
			a = a + 50
		case "C":
			a = a + 250
		case "D":
			a = a + 290
		case "E":
			a = a + 1200
		case "F":
			a = a + 2500
		}

		a = a / 3
	}
	a = int(math.Round(float64(a)))
	return a
}

func Score(a int, s int) (scr int) {
	scr = s + a
	a = 0
	return scr
}

func END(s int) {
	fmt.Println("\n遊戲結束，總分為：", s)
}

func main() {
	for ed < times {

		press = ""
		fmt.Println("請按ENTER")
		fmt.Scanln(&press)

		if press == "" {
			ram1 = rand.Intn(99) + 1
			ram2 = rand.Intn(99) + 1
			ram3 = rand.Intn(99) + 1

			p1 = P(ram1)
			p2 = P(ram2)
			p3 = P(ram3)

			cond, typ := IF()

			add = ADD(cond, typ)

			score = Score(add, score)
			ed = ed + 1

			test()

		} else {
			fmt.Println("請按ENTER")
		}
	}

	END(score)
}
