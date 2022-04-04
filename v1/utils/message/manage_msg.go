package message

import (
	"fmt"
	"io"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/meta-liam/websocket-go/v1/services/global"
	"github.com/meta-liam/websocket-go/v1/utils/conv"
	"github.com/meta-liam/websocket-go/v1/utils/enum"
)

func ManageMsg9(w io.Writer, op ws.OpCode, p []byte) error {
	wsutil.WriteServerMessage(w, op, p)
	return nil
}

func ManageMsg(w io.Writer, op ws.OpCode, b []byte) error {
	res := make(map[string]interface{})
	msg := string(b)
	fmt.Printf("[INFO]%+v\n", msg)
	_map := ToMap(b)
	res[enum.Jsonrpc_id] = _map["id"]
	res[enum.Jsonrpc] = enum.JsonrpcVersion
	if _map[enum.Jsonrpc_service] == "service-global" {
		return global.ManageMethod(w, op, _map, res)
	}
	//fmt.Printf("[res]Struct:%+v \n", res)
	p := []byte(conv.StructToJsonString(res))
	fmt.Printf("[res]back:%+v \n", conv.StructToJsonString(res))
	return wsutil.WriteServerMessage(w, op, p)
}
