package context

import (
	"fmt"
	"testing"
	"time"

	"github.com/kubeedge/kubeedge/beehive/pkg/core/model"
)

func TestSendSync(t *testing.T) {
	coreContext := GetContext(MsgCtxTypeChannel)
	coreContext.AddModule("test_src")
	messsage := model.NewMessage("")
	messsage.Content = "hello"

	go func() {
		resp, err := coreContext.SendSync("test_dest", *messsage, 5*time.Second)
		fmt.Printf("resp: %v, error: %v\n", resp, err)
	}()

	msg := coreContext.Receive("test_dest")
	fmt.Printf("receive msg: %v\n", msg)
	resp := msg.NewRespByMessage(&msg, "how are you")
	coreContext.SendResp(*resp)

	time.Sleep(5 * time.Second)
}
