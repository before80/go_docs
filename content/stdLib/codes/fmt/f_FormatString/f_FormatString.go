package main

import (
	"fmt"
)

// MyType 是一个自定义类型
type MyType struct {
	Value int
}

// Format 实现 fmt.Formatter 接口
func (mt MyType) Format(s fmt.State, verb rune) {
	formatString := fmt.FormatString(s, verb)
	fmt.Println("formatString=", formatString)
	// 可以根据 formatString 中的宽度等进行更复杂的处理
	width, _ := s.Width()
	fmt.Println("width=", width)
	precision, _ := s.Precision()
	fmt.Println("precision=", precision)
	fmt.Println("s.Flag(int('+'))=", s.Flag(int('+')))

	switch verb {
	case 'X': // 自定义大写十六进制输出
		if formatString == "%X" { // 如果没有指定宽度等
			_, _ = fmt.Fprintf(s, "%X", mt.Value)
		} else if formatString == "%#X" { // 如果指定了 # 标志
			_, _ = fmt.Fprintf(s, "0x%X", mt.Value)
		} else {

			formatted := fmt.Sprintf("%X", mt.Value)
			padding := ""
			if width > len(formatted) {
				for i := 0; i < width-len(formatted); i++ {
					padding += " "
				}
			}
			_, _ = fmt.Fprintf(s, "%s%s", padding, formatted)
		}
	case 'd':
		if s.Flag(int('+')) { // 将 '+' 转换为 int 类型
			_, _ = fmt.Fprintf(s, "+%d", mt.Value)
		} else {
			_, _ = fmt.Fprintf(s, "%d", mt.Value)
		}
	default:
		_, _ = fmt.Fprintf(s, "%%!%c(MyType=%d)", verb, mt.Value)
	}
}

func main() {
	my := MyType{Value: 255}
	fmt.Printf("%X\n", my)   // 输出: FF
	fmt.Printf("%#X\n", my)  // 输出: 0xFF
	fmt.Printf("%10X\n", my) // 输出:         FF
	fmt.Printf("%q\n", my)   // 输出:%!q(MyType=255)
	fmt.Printf("%d\n", my)
	fmt.Printf("%+d\n", my)
}
