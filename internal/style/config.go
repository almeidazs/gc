package style

import "os"

var (
	USE_ACCESSIBLE_MODE = os.Getenv("ACCESSIBLE_MODE") != ""
)
