package json

type FontItem struct {
	Family string
	Variants []string
	Subsets []string
	Version string
	LastModified string
	Files map[string]string
	Category string
	Kind string
	Menu string
}

type FullResponse struct {
	Kind string
	Items []FontItem
}
