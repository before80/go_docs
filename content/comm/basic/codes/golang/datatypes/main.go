package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"unsafe"
)

func main() {
	dataType := map[string][]string{}
	var b bool
	dataType["bool"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(b)), //占用大小
		fmt.Sprintf("%v", b),                //默认值
	}
	var bt byte
	dataType["byte"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(bt)), //占用大小
		fmt.Sprintf("%q", bt),                //默认值
	}

	var r rune
	dataType["rune"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(r)), //占用大小
		fmt.Sprintf("%q", r),                //默认值
	}

	var c64 complex64
	dataType["complex64"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(c64)), //占用大小
		fmt.Sprintf("%v", c64),                //默认值
	}

	var c128 complex128
	dataType["complex128"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(c128)), //占用大小
		fmt.Sprintf("%v", c128),                //默认值
	}

	var f32 float32
	dataType["float32"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(f32)), //占用大小
		fmt.Sprintf("%v", f32),                //默认值
		fmt.Sprintf("%.39e %s", math.SmallestNonzeroFloat32, "最小正非零值（保留39位小数）"), // 最小正非零值
		fmt.Sprintf("%.38e %s", math.MaxFloat32, "（保留38位小数）"),                   // 最大值
	}
	var f64 float64
	dataType["float64"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(f64)), //占用大小
		fmt.Sprintf("%v", f64),                //默认值
		fmt.Sprintf("%.44e %s", math.SmallestNonzeroFloat64, "最小正非零值（保留44位小数）"), // 最小正非零值
		fmt.Sprintf("%.42e %s", math.MaxFloat64, "（保留42位小数）"),                   // 最大值
	}

	var i8 int8
	dataType["int8"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(i8)), //占用大小
		fmt.Sprintf("%v", i8),                //默认值
		fmt.Sprintf("%v", math.MinInt8),      // 最小值
		fmt.Sprintf("%v", math.MaxInt8),      // 最大值
	}

	var i16 int16
	dataType["int16"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(i16)), //占用大小
		fmt.Sprintf("%v", i16),                //默认值
		fmt.Sprintf("%v", math.MinInt16),      // 最小值
		fmt.Sprintf("%v", math.MaxInt16),      // 最大值
	}

	var i32 int32
	dataType["int32"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(i32)), //占用大小
		fmt.Sprintf("%v", i32),                //默认值
		fmt.Sprintf("%v", math.MinInt32),      // 最小值
		fmt.Sprintf("%v", math.MaxInt32),      // 最大值
	}
	var i64 int64
	dataType["int64"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(i64)), //占用大小
		fmt.Sprintf("%v", i64),                //默认值
		fmt.Sprintf("%v", math.MinInt64),      // 最小值
		fmt.Sprintf("%v", math.MaxInt64),      // 最大值
	}

	var i int
	dataType["int"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(i)), //占用大小
		fmt.Sprintf("%v", i),                //默认值
		fmt.Sprintf("%v", math.MinInt),      // 最小值
		fmt.Sprintf("%v", math.MaxInt),      // 最大值
	}

	var ui8 uint8
	dataType["uint8"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(ui8)), //占用大小
		fmt.Sprintf("%v", ui8),                //默认值
		fmt.Sprintf("%v", 0),                  // 最小值
		fmt.Sprintf("%v", math.MaxUint8),      // 最大值
	}

	var ui16 uint16
	dataType["uint16"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(ui16)), //占用大小
		fmt.Sprintf("%v", ui16),                //默认值
		fmt.Sprintf("%v", 0),                   // 最小值
		fmt.Sprintf("%v", math.MaxUint16),      // 最大值
	}

	var ui32 uint32
	dataType["uint32"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(ui32)),    //占用大小
		fmt.Sprintf("%v", ui32),                   //默认值
		fmt.Sprintf("%v", 0),                      // 最小值
		fmt.Sprintf("%v", uint64(math.MaxUint32)), // 最大值
	}
	var ui64 uint64
	dataType["uint64"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(ui64)),    //占用大小
		fmt.Sprintf("%v", ui64),                   //默认值
		fmt.Sprintf("%v", 0),                      // 最小值
		fmt.Sprintf("%v", uint64(math.MaxUint64)), // 最大值
	}

	var ui uint
	dataType["uint"] = []string{
		fmt.Sprintf("%d", unsafe.Sizeof(ui)),    //占用大小
		fmt.Sprintf("%v", ui),                   //默认值
		fmt.Sprintf("%v", 0),                    // 最小值
		fmt.Sprintf("%v", uint64(math.MaxUint)), // 最大值
	}

	var typeNames []string
	for k, _ := range dataType {
		typeNames = append(typeNames, k)
	}

	slices.Sort(typeNames)

	fmt.Println("| 数据类型    | 占用字节数   | 默认值 | 数据最小值       | 数据最大值 |")
	fmt.Println("| ---------- | ---------- | ----  | ----------------------------  | ---------------------------- |")
	for _, tn := range typeNames {
		v := []string{tn}
		v = append(v, dataType[tn]...)
		for len(v) < 5 {
			v = append(v, "")
		}
		fmt.Println("| ", strings.Join(v, " | "), " |")
	}
}
