package attention

import (
	"unicode/utf8"
)

func length(){
	str:="⬇汉字"
	println(len(str))
}


func length1(){
	str:="⬇汉字"
	println(utf8.RuneCountInString(str))
}





