# Go-gRPC

This project apply the concepts of gRPC with Protocol Buffers communication with the following strategies:

  - Unary client/server
  - Client stream requests
  - Server stream responses
  - Bidirectional client/server streaming

## Usage

  Start the server in on terminal window

  ```sh
  make server
  ```

  Then copy and paste the function that you want to see in another terminal window:

  - Unary client/server:

  ```sh
  make unary-client
  ```
  - Client stream requests:

  ```sh
  make client-stream
  ```
  - Server stream responses

  ```sh
  make stream-server
  ```

  - Bidirectional client/server streaming

  ```sh
  make bidirectional
  ```
