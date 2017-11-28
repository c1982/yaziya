package yaziya

import(
	"testing"
)

var validations = []struct {
	input    int64
	expected string
}{
	{5, "beş"},	
	{13, "onüç"},
	{25, "yirmibeş"},
	{587, "beşyüzseksenyedi"},
	{1458, "bin dörtyüzellisekiz"},
	{14580, "ondört bin beşyüzseksen"},
	{145800, "yüzkırkbeş bin sekizyüz"},
	{1458001, "bir milyon dörtyüzellisekiz bin bir"},
	{14580012, "ondört milyon beşyüzseksen bin oniki"},
	{145800123, "yüzkırkbeş milyon sekizyüz bin yüzyirmiüç"},	
	{1458001235, "bir milyar dörtyüzellisekiz milyon bir bin ikiyüzotuzbeş"},
	{14580012356, "ondört milyar beşyüzseksen milyon oniki bin üçyüzellialtı"},
	{145800123567, "yüzkırkbeş milyar sekizyüz milyon yüzyirmiüç bin beşyüzaltmışyedi"},
	{1458001235678, "bir trilyon dörtyüzellisekiz milyar bir milyon ikiyüzotuzbeş bin altıyüzyetmişsekiz"},
	{14580012356789, "ondört trilyon beşyüzseksen milyar oniki milyon üçyüzellialtı bin yediyüzseksendokuz"},
	{1458001235678910, "bir katrilyon dörtyüzellisekiz trilyon bir milyar ikiyüzotuzbeş milyon altıyüzyetmişsekiz bin dokuzyüzon"},
}


func Test_Yaziya_Cevir(t *testing.T){

	for _,tt := range validations {
		text, err := Cevir(tt.input)

		if err != nil{			
			t.Errorf("Cevirirken hata meydana geldi: %v %v %v", tt.expected, tt.input, err)
		}

		if text != tt.expected{
			t.Errorf("Beklenmeyen değer. text:%v expect:%v input:%v",text, tt.expected, tt.input)
		}
	}
}