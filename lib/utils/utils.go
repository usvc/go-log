package utils

func ParseVariadicString(strings []string, defaultValue string) string {
	if len(strings) > 0 {
		return strings[0]
	}
	return defaultValue
}
