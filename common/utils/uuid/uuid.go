package uuidUtil

import (
	"github.com/google/uuid"
	"github.com/mooncake9527/orange/common/config"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Gen() string {
	return strings.Replace(uuid.NewString(), "-", "", -1)
}

func GenOrderNo() string {
	rand.Seed(time.Now().UnixNano())
	return config.Ext.PayConfig.Prefix + time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(100000))
}
