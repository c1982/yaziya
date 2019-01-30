# Yaziya
Rakamı Yazıya Çevirir

## Ne Yapabilir

desilyon birimine kadar çevirir.

## Örnek

```go
import (		
	"fmt"
	"github.com/c1982/yaziya"	
)

func main() {

	yazi, _ := yaziya.Cevir(1)
	fmt.Println(yazi)

	yazi, _ = yaziya.Cevir(28)
	fmt.Println(yazi)

	yazi, _ = yaziya.Cevir(152)
	fmt.Println(yazi)

	yazi, _ = yaziya.Cevir(1048)
	fmt.Println(yazi)

	yazi, _ = yaziya.Cevir(65870)
	fmt.Println(yazi)
}

//output:
//bir
//yirmisekiz
//yüzelliiki
//bin kırksekiz
//altmışbeş bin sekizyüzyetmiş
```

## Yükleme (Installation)

```
go get -u github.com/c1982/yaziya
```

## Bugs

Öneri veya İstekler için [issue tracker](https://github.com/c1982/Yaziya/issues) kullanın.


## İletişim

Oğuzhan
aspsrc@gmail.com

yeni özellik ekledim.

ingilizce brach'ı için değişiklklşe


release 1.3.0 için readme güncellendi.