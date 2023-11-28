package contentseed

import (
	"testing"
)

func TestUnmodifiedSeeder(t *testing.T) {
	seeder := NewContentSeed()
	output := seeder.Generate()

	if len(output) == 0 {
		t.Error("Expected output to be non-empty")
	}

	if output != seeder.Output {
		t.Error("Expected output to equal seeder.Output")
	}
}

func BenchmarkUnmodifiedSeeder(b *testing.B) {
	seeder := NewContentSeed()
	_ = seeder.Generate()
}
