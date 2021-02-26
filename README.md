# Golist

Golist is an easy to use botblock api wrapper written in golang!

## Installation

```
go get github.com/botblock/golist
```

## Getting started

You can view the basic documentation of the package at [pkg.go.dev](https://pkg.go.dev/github.com/botblock/golist)!

```go
package main

import (
    "fmt"
    "github.com/botblock/golist"
)

func main(){
    client := golist.Client{
        map[string]string{
            "botsfordiscord.com": "token",
            "top.gg": "token",
        }
    }

    res, err := client.PostStats("bot id", golist.Stats{
        ServerCount: 20,
    })

    if err != nil{
        fmt.Println(err.Error())
    } else {
        fmt.Println(res)
    }
}
```

### Creating a client with default properties

```go
client := golist.NewClient() // This method returns the Client struct with default properties!
client.AddToken("top.gg", "token")

res, err := client.PostStats("bot id", golist.Stats{
    ServerCount: 20,
})

if err != nil{
    fmt.Println(err.Error())
} else {
    fmt.Println(res)
}
```

### Post Stats

```go
res, err := client.PostStats("bot id", golist.Stats{
    ServerCount: 20,
    ShardID: 123,
    ShardCount: 1,
    Shards: [1, 2, 3, 4],
}) // Post stats and returns PostResponse struct!
```

### Get Bot

```go
bot, err := client.GetBot("bot id") // Returns the bot information
```

### Get lists

```go
lists, err := client.GetAllLists() // Returns an map of key with botlist id and value with list
list, err := client.GetList("list id"); // Return one paticualr list by id
ids, err := client.GetLegacyIDS(); // Returns the legacy ids of botlists registered on botblock
```

