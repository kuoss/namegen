# namegen

```go
package main

import (
	"fmt"

	"github.com/kuoss/namegen/en"
)

func main() {
	fmt.Println(en.Get("-")) // optimistic-wilson
}
```

```go
package main

import (
	"fmt"

	"github.com/kuoss/namegen/ko"
)

func main() {
	fmt.Println(ko.Get("_")) // 산만한_모래시계
}
```