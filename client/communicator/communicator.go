package communicator

import (
	"bytes"
	"encoding/json"
	"github.com/yaronsumel/simpleuser/server/semaphore"
	"github.com/yaronsumel/simpleuser/server/user"
	"log"
	"net/http"
	"time"
)

type CommManager struct {
	addr      string
	semaphore semaphore.Semaphore
}

func NewCommunicator(addr string, maxConnections int) *CommManager {
	return &CommManager{
		addr:      addr,
		semaphore: semaphore.NewSemaphore(maxConnections),
	}
}

var httpClient = http.Client{
	Timeout: time.Second * 60,
}

func (c *CommManager) SendEvent(u *user.Object) error {
	// gain semaphore
	c.semaphore.Wait()
	defer c.semaphore.Release()
	// encode to json
	jUser, err := json.Marshal(u)
	if err != nil {
		return err
	}
	if _, err := httpClient.Post(c.addr, "application/json", bytes.NewReader(jUser)); err != nil {
		log.Println("retrying", err)
	}
	return nil
}
