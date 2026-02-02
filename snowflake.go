package snowflake

import (
	"fmt"
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
	if generatorId > 4095 {
		panic("Snowflake node ID must be between 0 and 4095")
	}

	id = generatorId
	inc = 0
}

func new() int64 {
	now := int64(time.Now().UnixMilli()-epoch) << 22
	idPart := int64(id) & 0x0FFF << 10

	incMx.Lock()
	incPart := int64(inc & 0x03FF)
	inc++
	if inc > 1023 {
		inc = 0
	}
	incMx.Unlock()

	return now | idPart | incPart
}

type Snowflake int64

func (s Snowflake) String() string {
	return fmt.Sprintf("%d", int64(s))
}

func (s Snowflake) Time() time.Time {
	timestamp := (int64(s) >> 22) + epoch
	return time.UnixMilli(timestamp)
}

func (s Snowflake) GeneratorID() uint16 {
	return uint16((int64(s) >> 10) & 0x0FFF)
}

func (s Snowflake) Increment() uint16 {
	return uint16(int64(s) & 0x03FF)
}

func New() Snowflake {
	sf := new()

	lastSnowflakeMx.Lock()
	if sf <= lastSnowflake {
		sf = lastSnowflake + 1
	}
	lastSnowflake = sf
	lastSnowflakeMx.Unlock()

	return Snowflake(sf)
}
