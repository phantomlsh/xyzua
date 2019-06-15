package utils

import (
	"time"
)

// obtain the TimeStamp after lag second from now
func TimeStamp(lag uint) uint {
	return uint(time.Now().Unix()) + lag
}
