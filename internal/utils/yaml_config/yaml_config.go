package yaml_config

import (
	"apier/internal/container"
	"time"
)

var lastChangeTime time.Time
var containerFactory = container.CreateContainersFactory()

func init() {
	lastChangeTime = time.Now()
}
