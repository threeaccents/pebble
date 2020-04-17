# cache

Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/threeaccents/cache/api"
)

func main() {
	ctx := context.Background()

	c, err := api.NewClient(":4200", nil)
	if err != nil {
		panic(err)
	}

	if err := c.Set(ctx, "hello", []byte("world")); err != nil {
		panic(err)
	}

	value, err := c.Get(ctx, "hello")
	if err != nil {
		panic(err)
	}

	fmt.Printf("value: %s\n", string(value))
}
```