# pyrotic
code generator inspired by https://www.hygen.io/ for golang.



## Install

```
go install -u github.com/code-gorilla-au/pyrotic

```

initial setup creates a `_templates` directory at the root of the project to hold the generators

```
pyrotic init
```

create your first generator

```
pyrotic new cmd
```

## Run

default template path is `_templates` and default file extension is `.tmpl`

```
pyrotic generate <name of generator> --name <name-to-pass-in>

eg: pyrotic generate cmd --name setup
```

### Use different directory

```
pyrotic --path example/_templates generate cmd --name setup
pyrotic -p example/_templates generate cmd --name setup
```

### Use different file extension

```
pyrotic --extension ".template" generate cmd --name setup
pyrotic -x ".template" generate cmd --name setup
```

### Dry run mode

dry run will log to console rather than write to file

```
pyrotic -d generate cmd --name setup
pyrotic --dry-run generate cmd --name setup
```

## Pass in meta via cmd

you can pass in meta data via the `--meta` or `-m` flag, which takes in a comma (`,`) delimited list, and overrides the metadata within the template.

```

pyrotic generate fakr --meta foo=bar,bin=baz
pyrotic generate fakr --m foo=bar,bin=baz

```



## Formatter properties

formatter will pick up any of these variables within the `---` block and hydrate the metadata for the template. Any properties matching the signature will be added to the Meta property, for example `foo: bar` will be accessible by `{{ Meta.foo }}`.

| Property | Type | Default | Example |
| -------- | ---- | ------- | ------- |
| to: | string (path) | "" | src/lib/utils/readme.md |
| append: | bool | false | false |
| inject: | bool | false | false |
| before: | string | "" | type config struct |
| after: | string | "" | // commands |


## Built in template functions

ships with some already built in template funcs, some [examples](example/_templates/fakr/farkr_case.tmpl)

- caseSnake (snake_case)
- caseKebab (kebab-case)
- casePascal (PascalCase)
- caseLower (lowercase)
- caseTitle (Titlecase)
- caseCamel (camelCase)

we also provide some Inflections using [flect](https://github.com/gobuffalo/flect)

- pluralise
- singularise
- ordinalize
- titleize
- humanize

## Dev mode
provides the short file name with logging

```bash

ENV=DEV ./pyrotic -p example/_templates generate fakr --meta foo=bar,bin=baz

```