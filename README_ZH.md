# wang-sitong/p

### 以人类可看的方式打印golang数据

```
go get github.com/wang-sitong/p
```

### 例子:
```
import (
	"github.com/wang-sitong/p"
)
func main() {
	data := map[string]interface{}{
		"name":  "John Doe",
		"age":   30,
		"email": "john.doe@example.com",
	}
	p.R(data)
	p.R(1)
	p.D(2)
	p.R(3)
	
}
```

![img_2.png](img_2.png)
