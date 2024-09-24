# Pagination Go Project

The project revolves around fetching and managing folder structures, particularly for an organization. 

A token is returned, which can be further used to fetch results in chunks, so that the user interface can load data in installments.

## Compile instruction

Requires `Go` >= `1.20`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine
```
  go run main.go
```

## Folder structure

```
| go.mod
| README.md
| sample.json
| main.go
| folders
    | folders.go
    | folders_test.go
    | static.go
```

## What is pagination?
  - Pagination helps break down a large dataset into smaller chunks.
  - The small data chunk will then be served to the client side usually accompanied a token that points to the next chunk.
  - Further explanation and demonstration:
```
  original data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
  
  With pagination implementation, the results look like this:
  request() -> { data: [1, 2], token: "nQsjz" }

  The token can then be used to fetch more result:
  
  request("nQsjz") -> { data : [3, 4], token: "uJsnQ" }

  .
  .
  .

  And more results until there's no data left:
  
  { data: [9, 10], token: "" }
```
