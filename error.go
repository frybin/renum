package renum

// Error allows types to conform to a strongly defined interface, as well as act as enriched error builtins.
// The point of this is that as types that satisfy Error pass across package boundry, context and metadata
// is not lost.
type Error interface {
	Enum
	error
	Message() string
}

// YARPCError extends the Error interface to allow a type to additionally self-report a YARPC error code
// in order to enrich the handler's ability to respond with the proper code when an error of this type
// is encountered.
type YARPCError interface {
	Error
	YARPCResponder
}

// HTTPError extends the Error interface to allow a type to additionally self-report an HTTP Response code
// in order to enrich a net/http handler's ability to respond with the proper status code when an error
// of this type is encountered.
type HTTPError interface {
	Error
	HTTPResponder
}

// ProcessError extends the Error interface to allow a type to additionally self-report a specific
// exit code it wishes the handler to exit the process with.
type ProcessError interface {
	Error
	ProcessResponder
}

// ErrorTypeInfo is a type used to hold all the metadata associated with a given renum.Error where
// the fields of this structure are associated directly with the return values from the renum.Error interface.
// This acts as a convenience to help things like structured loggers or HTTP JSON responses to be have
// information extracted into a self contained object.
type ErrorTypeInfo struct {
	Name        string `json:"name,omitempty" mapstructure:"name,omitempty" yaml:"name,omitempty" toml:"name,omitempty"`
	Code        int    `json:"code,omitempty" mapstructure:"code,omitempty" yaml:"code,omitempty" toml:"code,omitempty"`
	Namespace   string `json:"namespace,omitempty" mapstructure:"namespace,omitempty" yaml:"namespace,omitempty" toml:"namespace,omitempty"`
	Path        string `json:"path,omitempty" mapstructure:"path,omitempty" yaml:"path,omitempty" toml:"path,omitempty"`
	Kind        string `json:"kind,omitempty" mapstructure:"kind,omitempty" yaml:"kind,omitempty" toml:"kind,omitempty"`
	Source      string `json:"source,omitempty" mapstructure:"source,omitempty" yaml:"source,omitempty" toml:"source,omitempty"`
	ImportPath  string `json:"import_path,omitempty" mapstructure:"import_path,omitempty" yaml:"import_path,omitempty" toml:"import_path,omitempty"`
	Description string `json:"description,omitempty" mapstructure:"description,omitempty" yaml:"description,omitempty" toml:"description,omitempty"`
	Message     string `json:"message,omitempty" mapstructure:"message,omitempty" yaml:"message,omitempty" toml:"message,omitempty"`
}

// ExtractErrorTypeInfo is used to take a renum.Error type and expand it's details into a more annotated
// structure. The primary purpose of this is to act as a helper to loggers who wish to expand interface methods
// of the renum.Error type into a nested, flat structure.
func ExtractErrorTypeInfo(e Error) ErrorTypeInfo {
	return ErrorTypeInfo{
		Name:        e.String(),
		Code:        e.Code(),
		Namespace:   e.Namespace(),
		Path:        e.Path(),
		Kind:        e.Kind(),
		Source:      e.Source(),
		ImportPath:  e.ImportPath(),
		Description: e.Description(),
		Message:     e.Message(),
	}
}