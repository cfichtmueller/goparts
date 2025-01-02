# Goparts

Goparts is a dependency-free component library for Go, designed to make building web apps simpler and more intuitive. Inspired by the ease of use found in libraries like React or Vite, Goparts enables you to write clean, component-based HTML generation without the verbosity and complexity of `html/template`.

## Features

- **No Dependencies**: Pure Go implementation, keeping your projects lightweight.
- **Simple Components**: Build reusable components effortlessly.
- **Intuitive API**: Minimalistic and easy-to-use design for generating HTML.
- **HTMX Support**: Built-in support for HTMX attributes like `hx-get` => `e.HXGet()` for dynamic interactions. 
- **Open Graph Tags**: Built-in support for the most common Open Graph tags, e.g., `og:title` => `e.OGTitle()`.
- **Twitter Tags**: Built-in support for the most common Twitter tags, e.g., `twitter:title` => `e.TwitterTitle()`.

## Motivation

- `html/template` can be cumbersome for building dynamic web applications.
- React and Vite offer great component-based workflows, but they’re not available in Go.
- I wanted something **super easy to use** without unnecessary overhead.
- I didn' want to learn another templating language

## Installation

Install Goparts using `go get`:

```bash
go get github.com/cfichtmueller/goparts
```

## Getting Started

Import Goparts into your Go project:

```go
import "github.com/cfichtmueller/goparts/e"
```

Here’s an example of building a simple HTML page using Goparts:

```go
package main

import (
    "fmt"
    "github.com/cfichtmueller/goparts/e"
)

func main() {
    page := e.Group(
        e.Doctype(),
        e.Html(
            e.Head(
                e.Title(e.Raw("Welcome to Goparts!")),
                e.Meta(e.Charset("UTF-8")),
                e.Meta(e.Name("viewport"), e.Content("width=device-width, initial-scale=1.0")),
                e.OGTitle("Welcome to Goparts!"),
                e.TwitterTitle("Welcome to Goparts!"),
                e.Link(e.RelStylesheet(), e.Href("/css/style.css")),
            ),
            e.Body(
                e.H1(e.Text("Hello, World!")),
                e.P(e.Text("This is a simple example built with Goparts.")),
                e.Div(e.HXGet("/dynamic-content")),
                e.Script(e.Defer(), e.Src("/js/htmx.min.js")),
            ),
        ),
    )

    w := &bytes.Buffer{}
	_ = page.Render(w)
	fmt.Println(w.String())
}
```

### Output:

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Welcome to Goparts!</title>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <meta property="og:title" content="Welcome to Goparts!" />
        <meta property="twitter:title" content="Welcome to Goparts!" />
        <link rel="stylesheet" href="/css/style.css" />
    </head>
    <body>
        <h1>Hello, World!</h1>
        <p>This is a simple example built with Goparts.</p>
        <div hx-get="/dynamic-content"></div>
        <script defer src="/js/htmx.min.js"></script>
    </body>
</html>
```

Note: goparts doesn't pretty print. The output has been prettified to improve readability.

## Design Decisions

### Tag Omission

Some HTML tags must not have an end tag (e.g. `<br />`). Goparts will render those elements without an end tag if they don't contain any non-attribute nodes. Adding any non-attribute nodes to such tags will result in invalid HTML.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

---

Enjoy building intuitive, component-based web applications with Goparts!