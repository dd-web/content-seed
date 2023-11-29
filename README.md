# Content-Seed
go content/text generation tool for rapid application prototyping.

## Install

Using go get

```bash
go get github.com/dd-web/content-seed
```

## Usage

First you need a reference to the content seeder. Config functions must be passed on instantiation, so if you need different outputs that depend on separate configurations, use two different content seeder instances.

Once we have a reference we can call generation methods to get our output.

```go
seeder := NewContentSeed()
output := seeder.GeneratePassage()
	fmt.Printf("%s = %s" ,output, seeder.Output)
```

There are four main generation methods, each represents a marginal increase/decrease in byte output.

 - GenerateWord() for a single word
 - GenerateSentence() for a single sentence
 - GenerateParagraph() for a single paragraph
 - GeneratePassage() for multiple paragraphs

Each time they are called, the Output field is overridden by the newly generated value.

## Config

Some configuration options aren't used for all of these, for instance sentence punctuation does nothing when generating a single word. For cases like that you can just leave them at their default.

Pass configuration functions directly when instantiating an instance.

```go
seeder := NewContentSeed(MinWordLength(1), MaxParagraphLength(5))
output := seeder.GenerateParagraph()
	fmt.Printf("%s = %s", output, seeder.Output)
```

### Word Config

 - `MinWordLength(i int)` minimum character count each generated word must have
 - `MaxWordLength(i int)` maximum character count each generated word can have

### Sentence Config

 - `MinSentenceLength(i int)` minimum word count each sentence must have
 - `MaxSentenceLength(i int)` maximum word count each sentence can have
 - `CapitalizeSentences(b bool)` if true will capitalize first character of each new sentence
 - `PunctuateSentences(b bool)` if true will end each sentence with a punctuation character (more on that later)

### Paragraph Config

 - `MinParagraphLength(i int)` minimum sentence count each generated paragraph must have
 - `MaxParagraphLength(i int)` maximum sentence count each generated paragraph can have
 - `MinParagraphCount(i int)` minimum number of generated paragraphs when calling `GeneratePassage()`
 - `MaxParagraphCount(i int)` maximum number of generated paragraphs when calling `GeneratePassage()`
 - `IndentParagraphs(b bool)` if true each new paragraph will begin with an indentation (two spaces)


