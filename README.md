# Gocase

[![Build Status](https://travis-ci.org/mtibben/gocase.png?branch=master)](https://travis-ci.org/mtibben/gocase)

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

gocase.ToSnake("May the force be with you")      // may_the_force_be_with_you
gocase.ToUpperSnake("May the force be with you") // MAY_THE_FORCE_BE_WITH_YOU
gocase.ToSpinal("May the force be with you")     // may-the-force-be-with-you
gocase.ToTrain("May the force be with you")      // May-The-Force-Be-With-You
gocase.ToCamel("May the force be with you")      // MayTheForceBeWithYou
gocase.ToLowerCamel("May the force be with you") // mayTheForceBeWithYou
```

See https://pkg.go.dev/github.com/mtibben/gocase for documentation.
