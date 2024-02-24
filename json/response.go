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

type Metadata struct {
	Service string
}

type ErrorDetail struct {
	Type string
	Reason string
	Domain string
	Metadata Metadata
}

type ErrorItem struct {
	Message string
	Detail string
	Reason string
}

type ErrorObject struct {
	Code int32
	Message string
	Errors []string
	Status string
	Details []ErrorDetail
}

type ErrorResponse struct {
	Error ErrorObject
}
