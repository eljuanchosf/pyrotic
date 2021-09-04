# pyrotic
code generator inspired by https://www.hygen.io/ for golang.



## Install

```
go get -u github.com/code-gorilla-au/pyrotic

```

initial setup creates a `_template` directory at the root of the project to hold the generators

```
pyrotic init
```

## Run

```
pyrotic generate <name of generator> --name <name-to-pass-in>

eg: pyrotic generate cmd --name setup
```

## Built in template functions

ships with some already built in template funcs, some [examples](example/_templates/fakr/farkr_case.tmpl)

- caseSnake (snake_case)
- caseKebab (kebab-case)
- casePascal (PascalCase)
- caseLower (lowercase)
- caseTitle (Titlecase)
- caseCamel (camelCase)

## Dev mode
provides the short file name with logging

```bash

ENV=DEV ./pyrotic -p example/_templates generate fakr

```