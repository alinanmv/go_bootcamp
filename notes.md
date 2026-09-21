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
- **Zero value ("null value")** — every type has a default value it's automatically set to if a variable is declared without an explicit assignment:
  - `int` → `0`
  - `float64` → `0.0`
  - `string` → `""`
  - `bool` → `false`
  - pointer → `nil` — a special built-in value representing the absence of an address (a pointer pointing at no memory location).

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

## Pointers
A pointer is a variable that stores the **memory address** of another variable, instead of storing a value directly.

- `&variable` — the **address-of** operator. Gives you the address where `variable` is stored (a pointer to it).
- `*Type` — declares a variable's type as "pointer to `Type`".
- `*pointer` — the **dereference** operator. Used on a pointer, it gives you the value stored at that address (read or write through the pointer).

```go
x := 10
p := &x        // p is of type *int, holds the address of x
fmt.Println(p)  // prints an address, e.g. 0xc0000140a0
fmt.Println(*p) // dereference: prints 10 (the value at that address)

*p = 20        // writes through the pointer
fmt.Println(x) // 20 — x itself changed, because p points at it
```

- The zero value of a pointer is `nil` (points at nothing). Dereferencing a `nil` pointer panics at runtime.
  ```go
  var p *int
  fmt.Println(p)  // <nil>
  fmt.Println(*p) // panic: runtime error
  ```

**Why use pointers — passing by reference:**
Go passes arguments **by value** by default — a function gets a *copy* of whatever you pass in, so changes made inside the function don't affect the original variable. Passing a pointer lets a function modify the caller's original variable instead of a copy:

```go
func double(n int) {
    n = n * 2 // only changes the local copy
}

func doubleByPointer(n *int) {
    *n = *n * 2 // changes the original value via its address
}

func main() {
    x := 5
    double(x)
    fmt.Println(x) // 5 — unchanged

    doubleByPointer(&x)
    fmt.Println(x) // 10 — changed
}
```

- Pointers are also used to **avoid copying large values** (structs) on every function call — passing a pointer copies just the address, not the whole struct.
- With a pointer to a struct, Go lets you access fields directly without manual dereferencing — `p.Field` is shorthand for `(*p).Field`.

## Structs
A `struct` is a custom type that groups related fields together.

```go
type User struct {
    firstName string
    lastName  string
    birthDate string
    createdAt time.Time
}
```

- Field names starting with a **lowercase** letter are unexported — only visible/usable inside the package that declares them (here, package `user`). Uppercase field/type names would be exported and usable from other packages.
- A struct literal can set fields by name: `User{firstName: "a", lastName: "b"}` (order doesn't matter, and you can skip fields to leave them at their zero value).

### Constructor functions
Go has no built-in constructors, so the convention is a `New` function that builds and returns the struct:

```go
func New(firstName, lastName, birthday string) (*User, error) {
    if firstName == "" || lastName == "" || birthday == "" {
        return nil, errors.New("all fields are required")
    }

    return &User{
        firstName: firstName,
        lastName:  lastName,
        birthDate: birthday,
        createdAt: time.Now(),
    }, nil
}
```

- Returning `*User` (a pointer) instead of `User` avoids copying the struct and lets callers share/mutate the same instance.
- The `(value, error)` return pattern is idiomatic Go: return a zero value (`nil` for a pointer) alongside a non-nil `error` when construction fails, and check it immediately with `if err != nil`.
- `errors.New("message")` creates a simple error value with a fixed message.

### Methods
A method is a function with a **receiver** — it attaches the function to a type:

```go
func (u *User) GetUserOutput() {
    fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}
```

- `(u *User)` is the receiver; inside the method, `u` refers to the specific struct instance the method was called on (`appUser.GetUserOutput()`).
- A **pointer receiver** (`*User`) is used so the method works uniformly whether it needs to read or later mutate the struct, and so it doesn't copy the whole struct on every call — same reasoning as passing pointers to functions.
- Go automatically takes the address for you when calling a pointer-receiver method on an addressable value, so `appUser.GetUserOutput()` works whether `appUser` is a `User` or a `*User`.

### Struct embedding (composition)
Go doesn't have classical inheritance — instead, one struct can **embed** another by declaring it without a field name:

```go
type Admin struct {
    email    string
    password string
    User
}
```

- `Admin` embeds `User` anonymously. This gives `Admin` direct access to all of `User`'s fields and methods, as if they were its own: `admin.firstName`, `admin.GetUserOutput()`.
- Building an embedded struct requires naming the embedded field explicitly in the literal (the field name is the type name):
  ```go
  func NewAdmin(email, password string) Admin {
      return Admin{
          email:    email,
          password: password,
          User: User{
              firstName: "admin",
              lastName:  "admin",
              birthDate: time.Now().Format("2006-01-02"),
              createdAt: time.Now(),
          },
      }
  }
  ```
- This is Go's way of achieving code reuse between related types ("has-a" via embedding, instead of "is-a" via inheritance).

### Packages
- Splitting code into packages (e.g. moving `User`/`Admin` into a `user` package under `user/user.go`) groups related types and enforces encapsulation — only exported (capitalized) names are visible from `main`, everything else stays private to the package.
- Import a local package by its module path: `import "example/bootcamp/user"`, then reference its exported names as `user.User`, `user.New(...)`, `user.NewAdmin(...)`.