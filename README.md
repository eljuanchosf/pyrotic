# pyrotic
code generator inspired by https://www.hygen.io/ for golang.


## Motivation
Why not use hygen? great question! I would recommend [hygen](https://www.hygen.io/) over this, however [hygen](https://www.hygen.io/) is written in js.
This project is for people who want to use a code generator and not have to install node.



## Install

```
go install github.com/code-gorilla-au/pyrotic@latest

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

default file extension is `.tmpl`

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

### Different shared folder

default shared templates path is `_templates/shared`

```
pyrotic --shared foo/bar generate cmd --name setup
pyrotic -s foo/bar generate cmd --name setup
```

## Formatter properties

Formatter will pick up any of these variables within the `---` block and hydrate the metadata for the template. Any properties matching the signature will be added to the Meta property, for example `foo: bar` will be accessible by `{{ Meta.foo }}`. View more [examples](example/_templates).

| Property | Type | Default | Example |
| -------- | ---- | ------- | ------- |
| to: | string (path) | "" | src/lib/utils/readme.md |
| append: | bool | false | false |
| inject: | bool | false | false |
| before: | string | "" | type config struct |
| after: | string | "" | // commands |


### Using shared templates

In some instances you will want to reuse some templates across multiple generators. This can be done by having a `shared` directory within the `_templates` directory.
Any templates that are declared in the [shared](example/_templates/shared/config.tmpl) directory will be loading along with the generator. [Reference](example/_templates/fakr/shared_config.tmpl) the shared template within your generator directory in order to inject / append / create file.


## Built in template functions

ships with some already built in template funcs, some [examples](example/_templates/fakr/farkr_case.tmpl)

- caseSnake (snake_case)
- caseKebab (kebab-case)
- casePascal (PascalCase)
- caseLower (lowercase)
- caseTitle (TITLECASE)
- caseCamel (camelCase)

we also provide some Inflections using [flect](https://github.com/gobuffalo/flect)

- pluralise
- singularise
- ordinalize
- titleize
- humanize

## Pass in meta via cmd

you can pass in meta data via the `--meta` or `-m` flag, which takes in a comma (`,`) delimited list, and overrides the `{{ .Meta.<your-property> }}` within the template.

```

pyrotic generate fakr --meta foo=bar,bin=baz
pyrotic generate fakr -m foo=bar,bin=baz

```


## Dev mode
provides the short file name with logging

```bash

ENV=DEV ./pyrotic -p example/_templates generate fakr --meta foo=bar,bin=baz

```