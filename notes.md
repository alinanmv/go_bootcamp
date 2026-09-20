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
- `%T` — prints the **type** of a variable (only works with `Printf`, not `Println`/`Print`).
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
## for loop
Go has only one looping keyword, `for`, but it covers all the usual loop forms:
- Standard three-part form — init; condition; post:
  ```go
  for i := 0; i < 10; i++ {
      // runs while i < 10
  }
  ```
- Condition-only form (like a `while` loop) — just the condition, no init/post:
  ```go
  for i < 10 {
      i++
  }
  ```
- No condition at all — an **infinite loop**; it runs forever unless stopped from inside:
  ```go
  for {
      // loops forever
  }
  ```
- `break` — exits the loop immediately, skipping everything after it in the loop body and any remaining iterations (similar to how `return` immediately exits a function).
- `continue` — skips the rest of the current iteration and jumps straight to the next one (re-checks the condition / runs the post statement), without exiting the loop entirely. 