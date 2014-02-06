Gocase
======

Gocase converts strings between different casing schemes.
- Snake: `may_the_force_be_with_you`
- Upper Snake: `MAY_THE_FORCE_BE_WITH_YOU`
- Spinal: `may-the-force-be-with-you`
- Train: `May-The-Force-Be-With-You`
- Camel: `MayTheForceBeWithYou`
- Lower Camel: `mayTheForceBeWithYou`

For example:
```go
import "github.com/mtibben/gocase"

gocase.ToSnake("May the force be with you")

```

See http://godoc.org/github.com/mtibben/gocase for documentation.
