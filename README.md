# Golang Microservice with gRPC & Protocol Buffers

### JSON in HTTP API vs Payload in Protocol Buffers Size

Payload in protocol buffers is much SMALLER compared to JSON. We can save a lot of bandwidth or storage space.
Additionally, passing JSON is actually quite CPU intensive due to its human readable format compared to passing in protocol buffers which is a binary.

=> Payload in protocol buffers: faster, more efficiency

=> Good for mobile or micro-controller which have a weaker spec

### Reference

- [gRPC website](https://grpc.io/)
- [HTTP/2](https://http2.github.io/)
- [HTTP/2 vs HTTP/1](https://imagekit.io/demo/http2-vs-http1)
