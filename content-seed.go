package contentseed

import "math/rand"

var (
	// default values for config settings.
	// these can be overridden using the exposed config functions.
	def_min_word_char_count          int = 3
	def_max_word_char_count          int = 10
	def_min_sentence_word_count      int = 8
	def_max_sentence_word_count      int = 18
	def_min_paragraph_sentence_count int = 3
	def_max_paragraph_sentence_count int = 7
	def_min_paragraph_count          int = 1
	def_max_paragraph_count          int = 6

	uppercase_offset        rune = 0x20
	lowercase_charset_start rune = 0x61
	lowercase_charset_end   rune = 0x7a
)

type ContentSeed struct {
	Output string
	config *Config
}

type Config struct {
	minWordCharCount int
	maxWordCharCount int

	minSentenceWordCount int
	maxSentenceWordCount int
	capitalizeSentences  bool
	punctuateSentences   bool
	punctuationWeights   map[string]int

	minParagraphSentenceCount int
	maxParagraphSentenceCount int
	minParagraphCount         int
	maxParagraphCount         int
	indentParagraphs          bool
	paragraphDelimiter        string
}

func defaultConfig() *Config {
	return &Config{
		minWordCharCount:          def_min_word_char_count,
		maxWordCharCount:          def_max_word_char_count,
		minSentenceWordCount:      def_min_sentence_word_count,
		maxSentenceWordCount:      def_max_sentence_word_count,
		minParagraphSentenceCount: def_min_paragraph_sentence_count,
		maxParagraphSentenceCount: def_max_paragraph_sentence_count,
		minParagraphCount:         def_min_paragraph_count,
		maxParagraphCount:         def_max_paragraph_count,
		capitalizeSentences:       true,
		punctuateSentences:        true,
		punctuationWeights:        map[string]int{".": 20, "!": 1, "?": 1},
		indentParagraphs:          true,
		paragraphDelimiter:        string(rune(0x0a)) + string(rune(0x0a)),
	}
}

type ConfigurFunc func(*Config) *Config

/* Config Fns */
/**************/

func MinWordLength(i int) ConfigurFunc {
	return func(c *Config) *Config {
		c.minWordCharCount = i
		return c
	}
}

func MaxWordLength(i int) ConfigurFunc {
	return func(c *Config) *Config {
		c.maxWordCharCount = i
		return c
	}
}

func MinSentenceLength(i int) ConfigurFunc {
	return func(c *Config) *Config {
		c.minSentenceWordCount = i
		return c
	}
}

func MaxSentenceLength(i int) ConfigurFunc {
	return func(c *Config) *Config {
		c.maxSentenceWordCount = i
		return c
	}
}

func MinParagraphCount(i int) ConfigurFunc {
	return func(c *Config) *Config {
		c.minParagraphCount = i
		return c
	}
}

func MaxParagraphCount(i int) ConfigurFunc {
	return func(c *Config) *Config {
		c.maxParagraphCount = i
		return c
	}
}

func CapitalizeSentences(b bool) ConfigurFunc {
	return func(c *Config) *Config {
		c.capitalizeSentences = b
		return c
	}
}

func PunctuateSentences(b bool) ConfigurFunc {
	return func(c *Config) *Config {
		c.punctuateSentences = b
		return c
	}
}

// Punctuations can be any string and are not limited to single characters.
// They are used to terminate a sentence and are chosen randomly by a cumulative weight
// of all punctuations weighted sums.
//
// default characters and their weights are as follows:
//
//	period . = 20
//	exclamation ! = 1
//	question ? = 1
//
// Periods are 20 times more likely than the others to be chosen. To modify the default weights
// use this function passing in the string you want to modify and pass the new weight. Example:
//
//	AddPunctiationItem(".", 10)
//
// Periods are now only 10 times more likely than the others to be chosen.
//
// You can add as many punctuation items as you want and the weights will be calculated for you.
// Weights do not need to add up to any specific sum but together must be greater than 0.
//
// There aren't any checks or panics so don't be stupid.
func AddPunctuationItem(s string, i int) ConfigurFunc {
	return func(c *Config) *Config {
		c.punctuationWeights[s] = i
		return c
	}
}

/* Entry Point */
/***************/

func NewContentSeed(cfg ...ConfigurFunc) *ContentSeed {
	config := defaultConfig()
	for _, fn := range cfg {
		config = fn(config)
	}
	return &ContentSeed{
		config: config,
		Output: "",
	}
}

func (l *ContentSeed) word() string {
	word := ""
	wordLen := randomBetween[int](l.config.minWordCharCount, l.config.maxWordCharCount)
	for i := 0; i < wordLen; i++ {
		word += string(randomBetween[rune](lowercase_charset_start, lowercase_charset_end))
	}
	return word
}

func (l *ContentSeed) sentence() string {
	sentence := ""
	wordCount := randomBetween[int](l.config.minSentenceWordCount, l.config.maxSentenceWordCount)
	for i := 0; i < wordCount; i++ {
		if i > 0 {
			sentence += " "
		}
		sentence += l.word()
	}
	if l.config.capitalizeSentences && len(sentence) > 0 {
		sentence = string(rune(sentence[0])-uppercase_offset) + sentence[1:]
	}
	if l.config.punctuateSentences && len(sentence) > 0 {
		sentence += randomWeightedFromMap[string](l.config.punctuationWeights)
	}
	return sentence
}

func (l *ContentSeed) paragraph() string {
	paragraph := ""
	if l.config.indentParagraphs {
		paragraph += "  "
	}
	sentenceCount := randomBetween[int](l.config.minParagraphSentenceCount, l.config.maxParagraphSentenceCount)
	for i := 0; i < sentenceCount; i++ {
		if i > 0 {
			paragraph += " "
		}
		paragraph += l.sentence()
	}
	return paragraph + l.config.paragraphDelimiter
}

func (l *ContentSeed) Generate() string {
	paragraphs := ""
	paragraphCount := randomBetween[int](l.config.minParagraphCount, l.config.maxParagraphCount)
	for i := 0; i < paragraphCount; i++ {
		paragraphs += l.paragraph()
	}

	l.Output = paragraphs
	return l.Output
}

/* Helper Fns */
/**************/

// random rune or int between supplied min and max values
func randomBetween[T int | rune](min, max T) T {
	return T(T(rand.Intn(int(max)-int(min))) + T(min))
}

// random weighted selection from map
func randomWeightedFromMap[T comparable](weights map[T]int) T {
	var cumulativeWeights []int
	var list []T
	cumulative := 0

	for item, weight := range weights {
		cumulative += weight
		cumulativeWeights = append(cumulativeWeights, cumulative)
		list = append(list, item)
	}

	r := rand.Intn(cumulativeWeights[len(cumulativeWeights)-1])

	for i, weight := range cumulativeWeights {
		if r < weight {
			return list[i]
		}
	}

	return list[0]
}
