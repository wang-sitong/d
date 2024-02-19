# wang-sitong/d

[中文文档](./README_ZH.md)

### Golang Simplified Date Component

```
go get github.com/wang-sitong/d
```

### Examples:
```
import (
	"github.com/wang-sitong/d"
)
func main() {
	// "2024-02-19 14:25:22" now
	date := d.Date()
	// "2024-02-20 15:38:49" tomorrow
	date = d.Date("Y-m-d H:i:s", 1)
	// "2024-02-18 15:39:24" yesterday
	date = d.Date("Y-m-d H:i:s", -1)
	// "2024" year
	date = d.Date("Y") 
	// "02" month
	date = d.Date("m")
	// "19" day
	date = d.Date("d")
	// "15" hour
	date = d.Date("H")
	// "42" minute
	date = d.Date("i")
	// "25" second
	date = d.Date("s")
}
```

