# Golang Microservice with gRPC & Protocol Buffers

### JSON in REST API vs Payload in Protocol Buffers Size

Payload in protocol buffers is much SMALLER compared to JSON. We can save a lot of bandwidth or storage space.
Additionally, passing JSON is actually quite CPU intensive due to its human readable format compared to passing in protocol buffers which is a binary.

- Payload in protocol buffers: faster, more efficiency
- Good for mobile or micro-controller which have a weaker spec
- Payload in protocol buffers is more restricting because the schema is fixed and we have types, cannot dynamically add parameters unlike JSON.
- gRPC uses HTTP/2
- gRPC also supports streaming unlike REST that only supports unary
- gRPC is a free design unlike REST (GET/POST/UPDATE/DELETE)

### Types API in gRPC

- Unary: one client's request per one server's response (for traditional way, like REST)
- Server Streaming: one client's request, multiple server's responses (for real time data from server like tax price streaming)
- Client Streaming: multiple client's requests, one server's response (for client uploading or updating data)
- Bi-directional Streaming: multiple client's requests, multiple server's responses (after the first request, the responses could arrive in any order)

### Reference

- [gRPC website](https://grpc.io/)
- [HTTP/2](https://http2.github.io/)
- [HTTP/2 vs HTTP/1](https://imagekit.io/demo/http2-vs-http1)
