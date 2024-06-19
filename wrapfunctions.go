package bootcamp

import (
	"github.com/alem-platform/ap"
)

func PrintString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
	ap.PutRune('\n')
}

func WrapperPrintStr(fn func(str *string)) func(str *string) {
	return func(str *string) {
		PrintString(*str)
		fn(str)
	}
}

func WrapperRot1(fn func(str *string)) func(str *string) {
	return func(str *string) {
		var res string
		for _, c := range *str {
			if c >= 'a' && c <= 'z' {
				res += string(rune('a' + ((c-'a')+1)%26))
			} else if c >= 'A' && c <= 'Z' {
				res += string(rune('A' + ((c-'A')+1)%26))
			} else {
				res += string(c)
			}
		}
		*str = res
		fn(str)
	}
}

func WrapperRot13(fn func(str *string)) func(str *string) {
	return func(str *string) {
		var res string
		for _, c := range *str {
			if c >= 'a' && c <= 'z' {
				res += string(rune('a' + ((c-'a')+13)%26))
			} else if c >= 'A' && c <= 'Z' {
				res += string(rune('A' + ((c-'A')+13)%26))
			} else {
				res += string(c)
			}
		}
		*str = res
		fn(str)
	}
}

func WrapperReverseStr(fn func(str *string)) func(str *string) {
	return func(str *string) {
		var res string
		for i := len(*str); i > 0; i-- {
			res += string((*str)[i-1])
		}
		*str = res
		fn(str)
	}
}

func WrapFunctions(decs []func(fn func(str *string)) func(str *string)) func(str *string) {
	return func(str *string) {
		if len(decs) == 0 {
			return
		}
		fnc := decs[0](func(str *string) {})
		for i := 1; i < len(decs); i++ {
			fnc = (decs[i])(fnc)
		}
		fnc(str)
	}
}

// func main() {
// 	mockFn := func(str *string) {
// 		return
// 	}

// 	fnPrint := WrapperPrintStr(mockFn)

// 	str := "salem?"
// 	fnPrint(&str) // salem?

// 	fnRot1 := WrapperRot1(mockFn)
// 	fnRot1(&str)
// 	fnPrint(&str) // tbmfn?

// 	fnRot13 := WrapperRot13(mockFn)
// 	fnRot13(&str)
// 	fnPrint(&str) // gozsa?

// 	fnReverse := WrapperReverseStr(mockFn)
// 	fnReverse(&str)
// 	fnPrint(&str) // ?aszog

// 	fmt.Println("United Func Results")
// 	wrappedFns := WrapFunctions([]func(fn func(str *string)) func(str *string){
// 		WrapperPrintStr,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperRot1,
// 		WrapperPrintStr,
// 		WrapperReverseStr,
// 		WrapperPrintStr,
// 	})
// 	wrappedFns(&str)
// 	// ?aszog
// 	// gozsa?
// 	// salem?
// }
