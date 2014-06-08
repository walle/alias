# Alias

[API Reference](http://godoc.org/github.com/walle/alias)

## Description

package `alias` adds alias/redirection to Martini apps using regular expressions to strings. Like mod_rewrite.

## Usage

```go
package main

import (
    "github.com/walle/alias"
    "github.com/go-martini/martini"
)

func main() {
    m := martini.Classic()

    alias.Add("^/my/path/to/rewrite", "/my/path", alias.SERVE)
    alias.Add("^/my/path/to/redirect", "/my/path", alias.REDIRECT)
    m.Use(alias.Handler())

    m.Get("/**", func(request *http.Request) string {
        return "Hello " + request.URL.Path
    })

    m.Run()
}
```

## Authors

* [Fredrik Wallgren](https://github.com/walle)