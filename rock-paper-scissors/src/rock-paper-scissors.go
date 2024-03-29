package main

import(
	"fmt"
	 "bufio"
	 "os"
	 "strconv"
	 "math/rand"
	 "time"
)

/*
 * グーチョキパーと数値をマッピング
 */
var handMp = map[uint]string{
	1: "グー",
	2: "チョキ",
	4: "パー",
}

/*
 * じゃんけん結果と数値をマッピング
 * handMpの数値の論理和
 */
var rpsResult = map[uint]string{
	1: "あいこ", // 001
	2: "あいこ", // 010
	4: "あいこ", // 100
	3: "グーの勝ち", // 011
	5: "パーの勝ち", // 101
	6: "チョキの勝ち", // 110
	7: "あいこ", // 111
}

/*
 * 入力された数値を元にhandMpのキーに該当する数値を返す
 */
func getHand(num int) (hand uint) {
	if (num < 0 || 2 < num) {
		panic("invalid argument num")
	}
	switch num {
	case 0:
		hand = 1 // グー : 001
	case 1:
		hand = 2 // チョキ : 010
	case 2:
		hand = 4 // パー : 100
	}
	return
}

func main() {
	fmt.Println("start!")
	fmt.Println("最初はグー")
	fmt.Println("じゃんけん...")
	fmt.Println(`
0 グー
1 チョキ
2 パー
	`)

	var userHand uint

	// 標準入力から読み取る
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// 0~2の数値が入力されたら次へ
		if num, err := strconv.Atoi(s.Text()); (err == nil && 0 <= num && num <= 2) {
			userHand = getHand(num)
			goto DONE
		}
		fmt.Println("0~2の数字を入力してください")
	fmt.Println(`
0 グー
1 チョキ
2 パー
	`)
	}
	DONE:

	fmt.Println("you:", handMp[userHand])

	rand.Seed(time.Now().UnixNano())
	cpHand := getHand(rand.Intn(3))
	fmt.Println("cp:", handMp[cpHand])

	// 論理和をとる
	res := userHand | cpHand
	fmt.Println(rpsResult[res])
}