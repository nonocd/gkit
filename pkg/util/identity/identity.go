package identity

import (
	"gkit/pkg/util"
	"log"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	uuid "github.com/satori/go.uuid"
)

var (
	once     sync.Once
	idWorker *snowflake.Node
	nodeMax  = 1024
)

func init() {
	once.Do(func() {
		var err error
		snowflake.Epoch = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Unix() * 1000
		id := getNodeID()
		if idWorker, err = snowflake.NewNode(int64(id)); err != nil {
			log.Fatalf("初始化 id worker 失败，%v", err)
		}
	})
}

// ID 获取主键
func ID() int64 {
	return idWorker.Generate().Int64()
}

// String 获取字符串格式主键
func String() string {
	return idWorker.Generate().String()
}

// UUID get uuid
func UUID() string {
	return uuid.NewV4().String()
}

func getNodeID() int {
	if ip, ok := util.GetIPv4(); ok {
		return (int(ip[2])<<8 + int(ip[3])) % nodeMax
	}
	return 0
}
