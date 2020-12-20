# gomrepo : go module report

This small tool adds license information to the `go module` information and outputs it in various formats.

## Usage

```sh
gomrepo -format {your_favorite_format} {your_project_directory_path}

# ex) gomrepo -format markdown ./my-golang-app
```

## Output sample

The following is an example of using this tool to generate a list of licenses for this tool's own dependent modules.

|ID|Version|License|
|:---|:---|:---|
|github.com/PuerkitoBio/goquery|v1.5.1|BSD-3-Clause|
|github.com/andybalholm/cascadia|v1.1.0|BSD-2-Clause|
|github.com/google/go-cmp|v0.5.4|BSD-3-Clause|
|golang.org/x/crypto|v0.0.0-20200622213623-75b288015ac9|BSD-3-Clause|
|golang.org/x/net|v0.0.0-20200822124328-c89045814202|BSD-3-Clause|
|golang.org/x/sys|v0.0.0-20200323222414-85ca7c5b95cd|BSD-3-Clause|
|golang.org/x/text|v0.3.0|BSD-3-Clause|
|golang.org/x/xerrors|v0.0.0-20191204190536-9bdfabe68543|BSD-3-Clause|

## Supported formats

- Markdown
- HTML
- AsciiDoc
- Textile

## Special thanks

This tool is using information from [pkg.go.dev](https://pkg.go.dev/) site.
I would like to take this opportunity to thank all the contributors of [pkg.go.dev](https://pkg.go.dev/) site.
