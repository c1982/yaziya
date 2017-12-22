package yaziya

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func Test_Yaziya_Cevir_Birler(t *testing.T) {
	var price float64 = 5

	text, err := Cevir(price)
	require.Nil(t, err)

	assert.Equal(t, "beş", text)
}

func Test_Yaziya_Cevir_Onlar(t *testing.T) {
	var price float64
	price = 25

	text, err := Cevir(price)

	if err != nil {
		t.Error(err)
	}

	if text != "yirmibeş" {
		t.Error("Text:", text)
	}
}

func Test_Yaziya_Cevir_Yuzler(t *testing.T) {
	var price float64
	price = 587

	text, err := Cevir(price)

	if err != nil {
		t.Error(err)
	}

	if text != "beşyüzseksenyedi" {
		t.Error("Text:", text)
	}
}

func Test_Yaziya_Cevir_Binler(t *testing.T) {
	var price float64
	price = 1458

	text, err := Cevir(price)

	if err != nil {
		t.Error(err)
	}

	if text != "bin dörtyüzellisekiz" {
		t.Error("Text:", text)
	}
}
