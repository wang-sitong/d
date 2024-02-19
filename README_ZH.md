# wang-sitong/d

[En document](./README.md)

### 以人类可用的方式获取日期

```
go get github.com/wang-sitong/d
```

### 例子:
```
import (
	"github.com/wang-sitong/d"
)
func main() {
	// "2024-02-19 14:25:22" 当前时间
	date := d.Date()
	// "2024-02-20 15:38:49" 明天
	date = d.Date("Y-m-d H:i:s", 1)
	// "2024-02-18 15:39:24" 昨天
	date = d.Date("Y-m-d H:i:s", -1)
	// "2024" 年
	date = d.Date("Y") 
	// "02" 月
	date = d.Date("m")
	// "19" 日
	date = d.Date("d")
	// "15" 时
	date = d.Date("H")
	// "42" 分
	date = d.Date("i")
	// "25" 秒
	date = d.Date("s")
}
```
