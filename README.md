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

### POST /run

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
