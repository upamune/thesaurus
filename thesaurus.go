package thesaurus

// Thesaurus is interface
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
