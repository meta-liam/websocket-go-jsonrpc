package global

import (
	"fmt"
	"io"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/meta-liam/websocket-go/v1/utils/conv"
	"github.com/meta-liam/websocket-go/v1/utils/enum"
)

/*
{"jsonrpc":"2.0","method":"getVersion","params":"","id":199,"service":"service-global"}
{"jsonrpc":"2.0","method":"init","params":"liam","id":199,"service":"service-global"}

{"jsonrpc":"2.0","method":"askAndAnswer","params":["liam",{"type":"JSON_RPC_CALLBACK","id":188}],"id":399,"service":"service-global"}
{"jsonrpc":"2.0","method":"listen","params":[{"type":"JSON_RPC_CALLBACK","id":288},2,500],"id":499,"service":"service-global"}

*/
var server = new(ServiceGlobal)

func ManageMethod(w io.Writer, op ws.OpCode, m map[string]interface{}, res map[string]interface{}) error {
	fmt.Printf("[INFO]%+v\n", m)
	switch m[enum.Jsonrpc_method] {
	case "getVersion":
		r := server.GetVersion()
		res[enum.Jsonrpc_result] = r
		return wsutil.WriteServerMessage(w, op, []byte(conv.StructToJsonString(res)))
	case "init":
		r := server.Init(conv.Interface2String(m[enum.Jsonrpc_params]))
		res[enum.Jsonrpc_result] = r
		return wsutil.WriteServerMessage(w, op, []byte(conv.StructToJsonString(res)))
	case "askAndAnswer":
		return server.AskAndAnswer(w, op, m, res)
	case "listen":
		return server.Listen(w, op, m, res)
	case "close":
		res[enum.Jsonrpc_result] = server.Close()
		return wsutil.WriteServerMessage(w, op, []byte(conv.StructToJsonString(res)))
	}
	//other
	res[enum.Jsonrpc_error] = "no this method"
	p := []byte(conv.StructToJsonString(res))
	return wsutil.WriteServerMessage(w, op, p)
}
