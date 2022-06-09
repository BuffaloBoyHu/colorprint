package colorprint

import "fmt"

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
	Reset  = "\033[0m"
)

func println(keyColor, key, valueColor string, value interface{}) {
	cKey := keyColor + key                   // 定义key 的颜色
	cVal := Reset + fmt.Sprintf("%v", value) // 定义value的颜色
	fmt.Printf("%s: %s\n", cKey, cVal)
}

func Info(key string, Value interface{}) {
	println(Blue, key, Reset, Value)
}
