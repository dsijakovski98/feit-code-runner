# FEIT Code - Runner

REST API responsible for running the code and tests, provided by Students and Teachers.

## Authentication

The API uses Clerk Authentication as it's middleware to access the endpoints. In order to authenticate to it, a `Bearer Token` must be passed as a Header:

```json
{
  "Authorization": "Bearer <CLERK_TOKEN_HERE>"
}
```

## Usage

### `POST /run`

Runs the code provided to it, returning the output of the executed program.

#### Request body

```jsonc
{
  "name": string, // Name of the task/program being executed (Ex. find_min)
  "code": string, // The actual code that needs to be executed
  "language": string // The programming language that this code is written in
}
```

#### Response

```jsonc
// ⛔️
{
    "error": string // Explanation of an error that occurred
}

// ✅
{
    "output": string // The output of the executed program
}
```

---

### `POST /cleanup`

- Removes any print/debug statements currently added to the provided code file
- Adds a special placeholder comment to the end\* of the code file

> \*end - This position can vary depending on the programming language (ex. for Golang this means the end of the `main` function).

This placeholder comment can then be replaced with a test call command

#### Request body

```jsonc
{
  "taskName": string, // Name of the task/program being executed (Ex. find_min)
  "code": string, // The actual code that needs to be executed
  "language": string // The programming language that this code is written in
}
```

#### Response

```jsonc
// ⛔️
{
    "error": string // Explanation of an error that occurred
}

// ✅
{
    "cleanCode": string // The cleaned up code
}
```

#### Usage (TypeScript)

Code file (BEFORE):

```ts
// someTask.ts

function SomeTask(number) {
  console.log("Debug statement");
  return number * 2;
}

console.log("Debug statement 2");
```

This file will be cleaned up:

- Debug statements will be removed
- Placeholder comment will be added at the end

Code file (AFTER):

```ts
// someTask.ts

function SomeTask(number) {
  return number * 2;
}

// PLACEHOLDER_PLACEHOLDER_PLACEHOLDER_PLACEHOLDER
```

Then, we can add any test call expression in place of the placeholder comment, to test for the correct output:

```diff
// someTask.ts

function SomeTask(number) {
  return number * 2;
}

- // PLACEHOLDER_PLACEHOLDER_PLACEHOLDER_PLACEHOLDER
+ console.log(taskFunction(2)); // Expecting 4
```

#### Usage (Golang)

Code file (BEFORE):

```go
// someTask.go
package main

import (
  "log"
  "fmt"
)

func SomeTask(num int) int {
  log.Println("Log debug")
  fmt.Printf("Fmt debug")

  return number * 2
}

func main() {
  fmt.Println("Debug 2")
}
```

This file will be cleaned up:

- Debug statements will be removed
- Placeholder comment will be added at the end

Code file (AFTER):

```go
// someTask.go
package main

import (
  "log"
  "fmt"
)

func SomeTask(num int) int {
  return number * 2
}

func main() {
  // PLACEHOLDER_PLACEHOLDER_PLACEHOLDER_PLACEHOLDER
}
```

Then, we can add any test call expression in place of the placeholder comment, to test for the correct output:

```diff
// someTask.go

func SomeTask(num int) int {
  return number * 2
}

func main() {
- // PLACEHOLDER_PLACEHOLDER_PLACEHOLDER_PLACEHOLDER
+ fmt.Println(taskFunction(2)) // Expecting 4
}
```

> In the case of `go` tests, an `"fmt"` import is also added, if it's not there already, for the test call print expression to work

This way, we can compare the output of the function being tested, using the `/run` endpoint.
