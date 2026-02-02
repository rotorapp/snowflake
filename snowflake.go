package snowflake

import (
	"sync"
	"time"
)

// January 1, 2026 in milliseconds
const epoch = 1767225600000

var (
	id    uint16 = 0
	inc   uint16 = 0
	incMx        = &sync.Mutex{}

	lastSnowflake   int64
	lastSnowflakeMx = &sync.Mutex{}
)

func Init(generatorId uint16) {
	if generatorId > 4096 {
		panic("Snowflake node ID must be between 0 and 4095")
	}

	id = generatorId
}

func new() int64 {
	now := int64(time.Now().UnixMilli()-epoch) << 22
	idPart := int64(id) & 0x0FFF << 12

	incMx.Lock()
	incPart := int64(inc & 0x03FF)
	inc++
	if inc > 1023 {
		inc = 0
	}
	incMx.Unlock()

	id := now | idPart | incPart

	return id
}

func New() int64 {
	id := new()

	lastSnowflakeMx.Lock()
	if id <= lastSnowflake {
		id = lastSnowflake + 1
	}
	lastSnowflake = id
	lastSnowflakeMx.Unlock()

	return id
}
