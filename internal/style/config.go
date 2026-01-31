package style

import "os"

var (
	USE_ACCESSIBLE_MODE = os.Getenv("GC_ACCESSIBLE_MODE") != ""
)
