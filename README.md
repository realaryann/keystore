<h1> KeyStore: A Redis-Server Implementation </h1>

<h2> Info </h2>
KeyStore is a thread safe key-value store that is capable of managing concurrent clients on a network. It utilizes a custom implementation of the RESP protocol over TCP which works natively with redis-cli. 

<h3>Supported Commands (Jan 2026)</h3>

| Command | Description  |
| :-------- | :------- | 
| `SET` | `SET an individual key to a value` | 
| `GET` | `GET an individual key's value` | 
| `HSET` | `SET a field and value associated with a key` | 
| `HGET` | `GET a value associated with a key's field` | 
| `EXPIRE` | `Add a TTL (seconds) to an existing key` | 
| `DEL` | `DELETE a key from KeyStore` | 
| `INCR` | `INCREMENT a key value` | 
| `DECR` | `DECREMENT a key value` |
| `RPUSH` | `RIGHT PUSH on a list within the cache` | 
| `LPUSH` | `LEFT PUSH on a list within the cache` | 
| `RPOP` | `RIGHT POP on a list within the cache` | 
| `LPOP` | `LEFT POP on a list within the cache` | 
| `PING/HELLO` | `PING the server and get a demo response` | 


<h3>Build Instructions</h3>

```go
    go build -o keystore main.go
    ./keystore
```
<h3>CLI Arguments</h3>

| Options | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `-p` | `string` | Custom TCP port number. Default: 6000 |

<h2> Model: </h2>

<img width="1536" height="1024" alt="img" src="https://github.com/user-attachments/assets/cd581988-3b3f-4828-8625-9007ecd3a057" />

