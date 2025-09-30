# Templates

This directory contains [Templ](https://templ.guide/) templates for generating HTML.

## Usage

1. Create a `.templ` file in this directory
2. Run `make generate` to generate Go code from templates
3. Import and use the generated components in your handlers

## Example

See `example/hello.templ` for a simple example.

```go
// In your handler
import "github.com/tagpro/auth-api/templates/example"

func handler(w http.ResponseWriter, r *http.Request) {
    component := example.HelloWorld("World")
    component.Render(r.Context(), w)
}
```

## Generating Templates

```bash
make generate
```

This runs `templ generate` to convert `.templ` files to Go code.
