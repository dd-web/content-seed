package contentseed

import (
	"regexp"
	"testing"
)

func TestDefaultPassage(t *testing.T) {
	seeder := NewContentSeed()
	output := seeder.GeneratePassage()

	if len(output) == 0 {
		t.Error("Expected output to be non-empty")
	}

	if output != seeder.Output {
		t.Error("Expected output to equal seeder.Output")
	}
}

func TestDefaultParagraph(t *testing.T) {
	seeder := NewContentSeed()
	output := seeder.GenerateParagraph()

	if len(output) == 0 {
		t.Error("Expected output to be non-empty")
	}

	if output != seeder.Output {
		t.Error("Expected output to equal seeder.Output")
	}

	match, err := regexp.MatchString(`(?m)^\s\s.*[\.?!]$`, output)
	if err != nil || !match {
		t.Error("Expected output to match regexp")
	}
}

func TestDefaultSentence(t *testing.T) {
	seeder := NewContentSeed()
	output := seeder.GenerateSentence()

	if len(output) == 0 {
		t.Error("Expected output to be non-empty")
	}

	if output != seeder.Output {
		t.Error("Expected output to equal seeder.Output")
	}

	match, err := regexp.MatchString(`(?m)^[A-Z].*[\.?!]$`, output)
	if err != nil || !match {
		t.Error("Expected output to match regexp")
	}
}

func TestDefaultWord(t *testing.T) {
	seeder := NewContentSeed()
	output := seeder.GenerateWord()

	if len(output) == 0 {
		t.Error("Expected output to be non-empty")
	}

	if output != seeder.Output {
		t.Error("Expected output to equal seeder.Output")
	}

	if len(output) < 3 || len(output) > 10 {
		t.Error("Expected output to be between 3 and 10 characters")
	}

	match, err := regexp.MatchString(`^[a-zA-Z]+$`, output)
	if err != nil || !match {
		t.Error("Expected output to match regexp")
	}
}

func BenchmarkUnmodifiedSeeder(b *testing.B) {

	for i := 0; i < b.N; i++ {
		seeder := NewContentSeed()
		_ = seeder.GeneratePassage()
	}

}
