# Recover

[![Build Status](https://travis-ci.org/ivpusic/grpool.svg?branch=master)](https://github.com/infinitasx/easi-go-aws)

> Middleware by error recover

## ⚙ Installation

```text
go get -u github.com/eininst/fiber-middleware-recover
```

## ⚡ Quickstart

```go
package main

import (
    "errors"
    recovers "github.com/eininst/fiber-middleware-recover"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    app.Use(recovers.New())

    app.Get("/", func(ctx *fiber.Ctx) error {
        panic(errors.New("my is error"))
        return ctx.Next()
    })
    _ = app.Listen(":8080")
}

```
> http response: 500 (Internal Server Error)

```text
Stack info: 

[ERROR] [88182] 2022/09/08 14:45:31 panic: my is error
goroutine 20 [running]:
github.com/eininst/fiber-middleware-recover.stackTraceHandler({0x1028cf1e0?, 0x1400011b4b0}, 0x500)
        /Users/wangziqing/go/fiber-middleware-recover/recover.go:36 +0x4c
github.com/eininst/fiber-middleware-recover.New.func1.1()
        /Users/wangziqing/go/fiber-middleware-recover/recover.go:56 +0x50
panic({0x1028cf1e0, 0x1400011b4b0})
        /usr/local/go/src/runtime/panic.go:884 +0x204
main.main.func1(0x1028d34a0?)
        /Users/wangziqing/go/fiber-middleware-recover/examples/main.go:14 +0x54
github.com/gofiber/fiber/v2.(*App).next(0x14000154f00, 0x14000176840)
        /Users/wangziqing/go/pkg/mod/github.com/gofiber/fiber/v2@v2.37.1/router.go:132 +0x184
github.com/gofiber/fiber/v2.(*Ctx).Next(0x14000054b18?)
        /Users/wangziqing/go/pkg/mod/github.com/gofiber/fiber/v2@v2.37.1/ctx.go:892 +0x5c
github.com/eininst/fiber-middleware-recover.New.func1(0x1028d34a0?)
        /Users/wangziqing/go/fiber-middleware-recover/recover.go:65 +0x70
```
> See [examples](/examples)

## License

*MIT*