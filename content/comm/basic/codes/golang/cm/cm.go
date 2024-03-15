package cm

import "fmt"

var AllTypes []string = []string{
	"bool", "byte", // 2
	"complex64", "complex128", // 2
	"int8", "int16", "int32", "int64", "int", "uint8", "uint16", "uint32", "uint64", "uint", // 10
	"rune", "uintptr", "string",
	"array", "slice", "map", "struct", "channel", "interface", "pointer", "func",
}

func PrintHeader(str string) {
	fmt.Println("|---", str, "---|")
}

func AddComma(str string) string {
	if len(str) > 0 {
		str += ","
	}
	return str
}

func Print(params ...any) {
	if len(params) <= 0 {
		return
	}

	fmtStr := ""
	for _, param := range params {
		switch param.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			fmtStr += AddComma(fmtStr) + "%d"
		case bool:
			fmtStr += AddComma(fmtStr) + "%T"
		case float32, float64:
			fmtStr += AddComma(fmtStr) + "%.6f(保留6位小数)"
		default:
			fmtStr += AddComma(fmtStr) + "%#v"
		}
	}
	fmt.Printf(fmtStr, params...)
}

func showTypeDetail(t string) {
	switch t {
	case "int8", "int16", "int32", "int64", "int", "uint8", "uint16", "uint32", "uint64", "uint":
		PrintHeader(t)
		fmt.Println("声明：\t var v ", t)
		fmt.Println("使用")

	}
}

func ShowTypeDetail(ts ...string) {
	if len(ts) > 0 {
		for _, t := range ts {
			_ = t
		}
	} else {

	}

}
