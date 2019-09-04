# gRPC basic
https://grpc.io/docs/tutorials/basic/go/ を参考にgRPCの基礎的なコードを実装する

## まずはGuideを読みすすめる
- https://grpc.io/docs/guides/
- gRPCはprotocol buffersを利用することができる。デフォルトでそれを利用する。
- protocol buffersはIDL(Interface Definition Language)・基礎となるメッセージ交換フォーマットとして利用できる
- protocは特殊なgRPCのプラグイン、`.proto`ファイルからコードを生成する
- バージョンが現在3（proto3）まで、3になるとsyntaxがシンプルになり、有用な機能が色々追加されている

## gRPC Concepts
- https://grpc.io/docs/guides/concepts/
- RPC: Remote Procedure Call
- Service definitionを行う、それはデフォルトではprotocol buffersで行われている。
    - https://developers.google.com/protocol-buffers/
    - service methodがいくつかある
        - Unary: 単一リクエストに対して単一レスポンスを返す
        - Server streaming: 一連のメッセージをストリームで送る
        - Client streaming: クライアント版、serverは受け付ける
        - Bidirectional streaming: 両方向
- Using the API surface
    - .protoファイルからプラグインによって生成されたコードを利用する
    - クライアントは定義したサービスのAPIをコールし、サーバーは対応したAPIを実装する
- 同期（Synchronous）・非同期（asynchronous）の両方に対応
- RPC life cycle
    - Unary RPC
        - シンプルなタイプ
        - client -> (metadata) -> server -> (metadata) -> client -> (request) -> server -> (response)
- Deadlines/Timeouts
    - どれだけRPC通信の完了に時間をかけるかをクライアントが指定することができる。
- Metadata
    - key(string)-value(string|binary)ペアのリスト形式
