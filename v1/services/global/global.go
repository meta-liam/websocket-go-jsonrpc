package global

import (
	"fmt"
	"io"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/meta-liam/websocket-go/v1/utils/conv"
	"github.com/meta-liam/websocket-go/v1/utils/enum"
)

type ServiceGlobal struct {
	Version string
}

const version = "0.0.1"

//{"jsonrpc":"2.0","method":"getVersion","params":[""],"id":199,"service":"service-global"}
func (s *ServiceGlobal) GetVersion() string {
	return version
}

//{"jsonrpc":"2.0","method":"init","params":["{\"version\":\"0.0.0\"}"],"id":199,"service":"service-global"}
func (s *ServiceGlobal) Init(config string) string {
	if config == "" {
		return "ok"
	}
	return config
}

//{"jsonrpc":"2.0","method":"askAndAnswer","params":["hi",{"type":"JSON_RPC_CALLBACK","id":188}],"id":199,"service":"service-global"}
func (s *ServiceGlobal) AskAndAnswer(w io.Writer, op ws.OpCode, m map[string]interface{}, res map[string]interface{}) error {
	res[enum.Jsonrpc_result] = "ok"
	wsutil.WriteServerMessage(w, op, []byte(conv.StructToJsonString(res)))

	resCb := make(map[string]interface{})
	resCb[enum.Jsonrpc] = enum.JsonrpcVersion
	params := m[enum.Jsonrpc_params]
	//resCb[enum.Jsonrpc_id] = m[enum.Jsonrpc_params].()[1]["id"]
	conv.Print_interfaces(params)
	fmt.Printf("params: %+v \n", params)
	_para := conv.Interface2Array(params)
	_m := conv.Interface2map(_para[1])
	resCb[enum.Jsonrpc_id] = _m[enum.Jsonrpc_id]
	resCb[enum.Jsonrpc_type] = _m[enum.Jsonrpc_type]
	resCb[enum.Jsonrpc_params] = []string{"ok"}
	wsutil.WriteServerMessage(w, op, []byte(conv.StructToJsonString(resCb)))
	return nil
}

//{"jsonrpc":"2.0","method":"listen","params":[{"type":"JSON_RPC_CALLBACK","id":188},3,10],"id":199,"service":"service-global"}
func (s *ServiceGlobal) Listen(w io.Writer, op ws.OpCode, m map[string]interface{}, res map[string]interface{}) error {
	res[enum.Jsonrpc_result] = "ok"
	wsutil.WriteServerMessage(w, op, []byte(conv.StructToJsonString(res)))

	resCb := make(map[string]interface{})
	resCb[enum.Jsonrpc] = enum.JsonrpcVersion
	params := m[enum.Jsonrpc_params]
	//resCb[enum.Jsonrpc_id] = m[enum.Jsonrpc_params].()[1]["id"]
	conv.Print_interfaces(params)
	fmt.Printf("params: %+v \n", params)
	_para := conv.Interface2Array(params)
	_m := conv.Interface2map(_para[0])
	resCb[enum.Jsonrpc_id] = _m[enum.Jsonrpc_id]
	resCb[enum.Jsonrpc_type] = _m[enum.Jsonrpc_type]

	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
		st := fmt.Sprintf("hello %d", sum)
		resCb[enum.Jsonrpc_params] = []string{st}
		wsutil.WriteServerMessage(w, op, []byte(conv.StructToJsonString(resCb)))
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *ServiceGlobal) Close() string {
	return "ok"
}
