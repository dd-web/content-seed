# Content-Seed
go content/text generation tool for rapid application prototyping.

*NOTE* this is very new for me. I made this for myself so I don't need to keep doing it for new projects every time. Feel free to use and I'll do the best I can to update and maintain. Updates coming soon.

## Install

Using go get

```bash
go get github.com/dd-web/content-seed
```

## Usage

By default content-seed is set up to generate a couple random paragraphs. You'll need to modify the config via the config funcs provided - more on that later.

```go
seeder := NewContentSeed()
output := seeder.Generate()
	fmt.Printf("%s = %s" ,output, seeder.Output)
```

