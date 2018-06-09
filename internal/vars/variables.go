package vars

var varMap = make(map[string]string)

const (
	// JwtSecret ...
	JwtSecret = "JWT_SECRET"
)

// GetVar ...
func GetVar(name string) string {
	return varMap[name]
}

// SetVar ...
func SetVar(name string, value string) {
	varMap[name] = value
}
