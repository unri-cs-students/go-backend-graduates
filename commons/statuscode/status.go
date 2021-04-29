package statuscode

type StatusCode int

const (
	UnknownError StatusCode = iota
	UncompatibleJSON
	OK
	EmptyParam
	UnknownUUID
	NotFound
	NoAccess
	UnknownType
)

func (s StatusCode) String() string {
	return [...]string{"unknown_error", "uncompatible_json", "ok", "empty_param", "unknown_uuid", "not_found", "no_access", "unkn"}[s]
}

