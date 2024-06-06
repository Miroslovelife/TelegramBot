package yandex_gpt

import (
	"sync"
	"time"
)

type MessageStore struct {
	messages []Message
	mutex    sync.Mutex
	duration time.Time
}
