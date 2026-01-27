<h1> KeyStore: A Redis-Server Clone </h1>

<h2> Info </h2>
KeyStore is a mock-up redis server that is capable of managing concurrent clients on a network. It utilizes a custom implementation of the RESP protocol over TCP which works natively with redis-cli. 

Supported Commands (Jan 2026)

| Command | Description  |
| :-------- | :------- | 
| `SET` | `SET an individual key to a value` | 
| `GET` | `GET an individual key's value` | 
| `HSET` | `SET a field and value associated with a key` | 
| `HGET` | `GET a value associated with a key's field` | 
| `EXPIRE` | `Add a TTL (seconds) to an existing key` | 
| `DEL` | `DELETE a key from KeyStore` | 
| `PING/HELLO` | `PING the server and get a demo response` | 

```go
    go build -o keystore main.go
    ./keystore
```
CLI Arguments

| Options | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `-p` | `string` | Custom TCP port number. Default: 6000 |

<h2> Model: </h2>

<img width="1536" height="1024" alt="img" src="https://github.com/user-attachments/assets/cd581988-3b3f-4828-8625-9007ecd3a057" />

