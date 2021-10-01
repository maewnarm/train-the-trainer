package classroom

import (
	"test-golang/classRoom/class3"
	"test-golang/classRoom/class4"
	"test-golang/classRoom/class5"
	"test-golang/classRoom/class6"
)

func Mainclass(c int) {
	switch c {
	case 3:
		class3.Mainclass3()
	case 4:
		class4.Mainclass4()
	case 5:
		class5.Mainclass5()
	case 6:
		class6.Mainclass6()
	}
}
