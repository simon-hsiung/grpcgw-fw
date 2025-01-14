# gRPC-Gateway general kits


## Why using [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/)?

* To have well-defined interface and payload
    - Generate documents automatically
* To serve both web requests and internal data changing


## Protocol buffer toolchain

Use [`buf`](https://buf.build/docs/) toolchain to build protocol buffers and
leverage [`protovalidate`](https://github.com/bufbuild/protovalidate) to do input validation.

`protovalidate` is the spiritual successor to `protoc-gen-validate`.
It does not require any code generation and supports custom constraints.


## HTTP Middlewares

### Incoming/outgoing header matcher

By default gRPC-Gateway passes only
[few predefined](https://github.com/grpc-ecosystem/grpc-gateway/blob/v1.16.0/runtime/mux.go#L73)
HTTP headers to/from gRPC context.
Need to write custom header matcher to pass specified custom headers if necessary.

### Marshaler

Specify marshaler/unmarshaler according to MIME type.

#### Streaming

Sometimes, we need to transfer large files.
To be compatible with frontend requirements and avoid memory exhaustion,
using the `multipart/form-data` format with chunked transfer is essential.
Our marshaler should be able to accommodate this need.

### Error handler

Custom error handler to generate http response suit for product needs while API fails.

A good handler should be able to
* understand the standard gRPC `status.Status` error
* Write proper HTTP status code and headers according to the error content
* Compose proper body according to error content and requested content-type.

### Other common middlewares

#### Encoding

Ability to decode/encode body data according to the `Content-Encoding`/`Accept-Encoding` headers.

#### Request ID

Let each request be marked with unique ID which is unified between each sub-services.

#### Logging

Writing debug logs for each request.


## gRPC middlewares

### Data validation

Data validation should be executed in gRPC layer because we need the unmarshaled data
that having strict data type and compiled validation rules.
