# message

```js
  //send:{"jsonrpc":"2.0","method":"getVersion","params":[""],"id":199,"service":"service-global"}
  //back:{"type":"result","data":{"jsonrpc":"2.0","id":199,"result":"0.0.1"}}

  //send:{"jsonrpc":"2.0","method":"init","params":["{\"version\":\"0.0.0\"}"],"id":199,"service":"service-global"}
  //back:{"type":"result","data":{"jsonrpc":"2.0","id":199,"result":"{\"version\":\"0.0.0\"}"}}

  //send:{"jsonrpc":"2.0","method":"askAndAnswer","params":["hi",{"type":"JSON_RPC_CALLBACK","id":188}],"id":199,"service":"service-global"}
  //back:{"type":"result","data":{"jsonrpc":"2.0","id":199,"result":null}}
  //back:{"type":"callback","data":{"jsonrpc":"2.0","type":"JSON_RPC_CALLBACK","id":188,"params":["Ask: hi\nAnswer: Sorry, I don't know."]}}

  //send:{"jsonrpc":"2.0","method":"listen","params":[{"type":"JSON_RPC_CALLBACK","id":188},3,10],"id":199,"service":"service-global"}
  //back: {"type":"result","data":{"jsonrpc":"2.0","id":199,"result":null}}
  //back:{"type":"callback","data":{"jsonrpc":"2.0","type":"JSON_RPC_CALLBACK","id":188,"params":["Hello world 1"]}}
  //back:{"type":"callback","data":{"jsonrpc":"2.0","type":"JSON_RPC_CALLBACK","id":188,"params":["Hello world 2"]}}
  //back:{"type":"callback","data":{"jsonrpc":"2.0","type":"JSON_RPC_CALLBACK","id":188,"params":["Hello world 3"]}}


  //send:{"jsonrpc":"2.0","method":"close","params":[""],"id":199,"service":"service-global"}
  //back:{"type":"result","data":{"jsonrpc":"2.0","id":199,"result":null}}

```
