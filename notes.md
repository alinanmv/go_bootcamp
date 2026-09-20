# Go Notes

## Data types

- `bool`
- `string`
- `int` — whole numbers
- `uint` — positive whole numbers
- `byte`
- `rune` — alias for `int32`
- `float` — decimal numbers
- `complex` — imaginary numbers

## Variables & constants

- Declare with `var name type`, or use `:=` to let Go infer the type.
- Constants (`const`) can be computed at compile time.

## fmt package

- `Print` — prints args as-is; adds a space between two operands only if neither is a string; no trailing newline.
- `Println` — always adds a space between args and a trailing `\n`. Good for quick debug output.
- `Printf` — formats using verbs (`%v`, `%T`, `%d`, `%s`, `%f`, ...); no automatic spacing or newline, you control it via the format string.
- `%v` — placeholder for a variable's default/natural formatting.
- `%T` — prints the **type** of a variable (only works with `Printf`, not `Println`/`Print`).
- `%f` — used for floats; `%.0f` controls how many digits after the decimal point (e.g. `%.2f` = 2 decimal places).
- `Sprintf` — same formatting as `Printf`, but **returns a string** instead of printing it (useful to store, return, or reuse the text, e.g. building error messages or filenames).
- `Fprintf` — writes formatted output to an `io.Writer` (a file, `os.Stderr`, a buffer, etc.) instead of stdout.

## Conditionals

- `if` with initializer — a short form of `if` that declares a variable right before the condition, scoped only to the `if`/`else` block:
  ```go
  if x := getValue(); x > 10 {
      // x visible here (and in else)
  }
  // x not visible here
  ```
  Common idiom for error checks:
  ```go
  if err := doSomething(); err != nil {
      // handle err
  }
  ```
  variables that declared in that function scoped in that function 
- 