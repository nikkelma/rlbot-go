# rlbot-go Design - v1

rlbot-go is intended to provide access to Rocket League's RLBot API in Go as well as providing
standardized building blocks for creating a custom Rocket League bot in Go.

This document intends to detail the design of `rlbot-go`, including the package's responsibilities,
any deviations from existing RLBot libraries, and any context on how the package is intended to be
used by consumers.

## Design Goals

Overall, `rlbot-go` intends to fulfill the following design goals:
- Standardize names to match underlying purposes
- Allow dependency flexibility when possible
- Adhere to effective Go best practices
- Provide performant, robust implementations

## Deviations

Other libraries accomplishing the same purpose for different languages exist - Python, Rust, C++,
Java, C#, Java, and others. Attempts will be made to follow the conventions of both the underlying
shared library and the current implementations for other languages. Even so, a couple complications 
arise:
1. Go as a language deviates in non-trivial ways from these languages in various ways, including
concurrency model and lack of strict inheritance.
2. The underlying shared library's interface has undergone some refactors, and naming conventions do
not fully reflect the intention of the interface.



