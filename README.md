# Usage

```go
package main

import (
	"fmt"

	"github.com/dotnetdr/devp"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{1, "a", 18}
	devp.PrintStack("")
	fmt.Println(devp.Stack())
	fmt.Println(devp.GetObjectJson("u", u))
}
```

output

```
[2016-03-25 15:45:56.66551]
goroutine 1 [running]:
github.com/dotnetdr/devp.getStackInfo(0x0, 0x0)
	/go/src/github.com/dotnetdr/devp/devp.go:63 +0x84
github.com/dotnetdr/devp.PrintStack(0x0, 0x0)
	/go/src/github.com/dotnetdr/devp/devp.go:49 +0x215
main.main()
	/go/src/github.com/dotnetdr/devp/demo/main.go:17 +0x74

goroutine 1 [running]:
github.com/dotnetdr/devp.getStackInfo(0x0, 0x0)
	/go/src/github.com/dotnetdr/devp/devp.go:63 +0x84
github.com/dotnetdr/devp.Stack(0x0, 0x0)
	/go/src/github.com/dotnetdr/devp/devp.go:54 +0x31
main.main()
	/go/src/github.com/dotnetdr/devp/demo/main.go:18 +0x79

u {"id":1,"name":"a","age":18}
```
