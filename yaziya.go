package yaziya

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	birler, onlar, yuzler, binler []string
)

func init() {
	birler = []string{"", "bir", "iki", "üç", "dört", "beş", "altı", "yedi", "sekiz", "dokuz"}
	onlar = []string{"", "on", "yirmi", "otuz", "kırk", "elli", "altmış", "yetmiş", "seksen", "doksan"}
	yuzler = []string{"", "yüz", "ikiyüz", "üçyüz", "dörtyüz", "beşyüz", "altıyüz", "yediyüz", "sekizyüz", "dokuzyüz"}
	binler = []string{"", "bin", "milyon", "milyar", "trilyon", "katrilyon", "kentilyon", "seksilyon", "septilyon", "oktilyon", "nobilyon", "desilyon"}
}

//Cevir rakam olarak aldığı parametreyi sözlü bir şekilde geri döndürür.
func Cevir(number int64) (text string, err error) {
	var b = 0

	numberStr := fmt.Sprint(number)
	numberStr = padding3(numberStr)

	for i := len(numberStr); i > 0; i -= 3 {
		sliceText := admiral(numberStr[i-3 : i])
		if sliceText != "" {
			text = fmt.Sprintf("%s %s %s", sliceText, binler[b], text)
		}
		b++
	}

	text = strings.TrimSpace(text)

	if strings.HasPrefix(text, "bir bin") {
		text = strings.TrimPrefix(text, "bir ")
	}

	return text, err
}

func admiral(str string) string {

	bir, _ := strconv.ParseInt(string(str[0]), 10, 64)
	on, _ := strconv.ParseInt(string(str[1]), 10, 64)
	yuz, _ := strconv.ParseInt(string(str[2]), 10, 64)

	return fmt.Sprint(yuzler[bir], onlar[on], birler[yuz])
}

func padding3(str string) string {

	for {
		m := len(str) % 3
		if m > 0 {
			str = fmt.Sprint("0", str)
		} else if m == 0 {
			break
		}
	}

	return str
}
