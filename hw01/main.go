package main

type spellItem struct {
	Name string
	Spell func(x, y int) bool
}

var spells = []spellItem{
	spellItem{"spell 01 [x<y]", func(x, y int) bool {
		return x < y
	}},
	spellItem{"spell 02 [x=y]", func(x, y int) bool {
		return x == y
	}},
	spellItem{"spell 03 [x+y == 24]", func(x, y int) bool {
		return x+y == 24
	}},
	spellItem{"spell 04 [30 - y > x]", func(x, y int) bool {
		return 30 - y > x
	}},
	spellItem{"spell 05 [(x*2) == y || (x+1)*2  == y+1]", func(x, y int) bool {
		return (x*2) == y || (x+1)*2  == y+1
	}},
	spellItem{"spell 06 [10 - x >= 1 || 10 - y >= 1]", func(x, y int) bool {
		return 10 - x >= 1 || 10 - y >= 1
	}},
	spellItem{"spell 07 [x >= 16 &&  y >= 16]", func(x, y int) bool {
		return x >= 16 &&  y >= 16
	}},
	spellItem{"spell 08 [x * y == 0]", func(x, y int) bool {
		return x * y == 0
	}},
	spellItem{"spell 09 [x + 10 < y || y + 10 < x]", func(x, y int) bool {
		return x + 10 < y || y + 10 < x
	}},
	spellItem{"spell 10 [x < y && x * 2 >= y -1]", func(x, y int) bool {
		return x < y && x * 2 >= y -1
	}},
	spellItem{"spell 11 [x == 1 || y == 1 || x == 24-1 || y == 24-1]", func(x, y int) bool {
		return x == 1 || y == 1 || x == 24-1 || y == 24-1
	}},
	spellItem{"spell 12 [x*x + y*y <= 400]", func(x, y int) bool {
		return  x*x + y*y <= 400
	}},
	spellItem{"spell 13 [x+y >= 20 && x+y <= 28]", func(x, y int) bool {
		return x+y >= 20 && x+y <= 28
	}},
	spellItem{"spell 15 [x + 10 <= y &&  x + 20 >= y || y + 10 <= x &&  y + 20 >= x]", func(x, y int) bool {
		return x + 10 <= y &&  x + 20 >= y || y + 10 <= x &&  y + 20 >= x
	}},
	spellItem{"spell 16 [x - y >= -9 && y - x >= -9 && x + y >= 15 && x + y <= 33]", func(x, y int) bool {
		return x - y >= -9 && y - x >= -9 && x + y >= 15 && x + y <= 33
	}},
	spellItem{"spell 18 [x > 0 && y < 2 || y > 0 && x < 2]", func(x, y int) bool {
		return  x > 0 && y < 2 || y > 0 && x < 2
	}},
	spellItem{"spell 19 [x > 0 && y < 2 || y > 0 && x < 2]", func(x, y int) bool {
		return  x <= 0 || y <= 0 || x >= 24 || y >= 24
	}},
	spellItem{"spell 20 [(x+y)%2 == 0]", func(x, y int) bool {
		return (x+y)%2 == 0
	}},
	spellItem{"spell 22 [(x+y)%3 == 0]", func(x, y int) bool {
		return (x+y)%3 == 0
	}},
	spellItem{"spell 23 [x%3 == 0 && y%2 == 0]", func(x, y int) bool {
		return x%3 == 0 && y%2 == 0
	}},
	spellItem{"spell 24 [x == y || x+y == 24]", func(x, y int) bool {
		return x == y || x+y == 24
	}},
	spellItem{"spell 25 [x%6 == 0 || y%6 == 0]", func(x, y int) bool {
		return x%6 == 0 || y%6 == 0
	}},
}


func main() {
	for _, spellItem := range spells {
		print("~~~~ Ready to cast " + spellItem.Name + " ~~~~\n")
		for x := 0; x < 25; x++ {
			for y := 0; y < 25; y++ {
				if spellItem.Spell(x, y) {
					print("# ")
				} else {
					print(". ")
				}
			}
			print("\n")
		}
	}
}