# fetch

Fetch in golang

## Installation

```cmd
go get -u github.com/udfordria/go-fetch
```

## Usage

### `GET` example with `timemout` and `params`

```go
res, err := fetch.Fetch(fetch.FetchArgs{
    Url: "https://jsonplaceholder.typicode.com/posts",
    Params: map[string]string{
        "userId": "1",
    },
    Timeout: time.Second * 5,
})

if err != nil {
    panic(err)
}

if res.StatusCode == http.StatusOK {
    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    bodyString := string(bodyBytes)
    t.Log(bodyString)
}
```

### `PUT` example with `timeout`, `headers` and `body`

```go
type RequestPayload struct {
    Id     int    `json:"id"`
    Title  string `json:"title"`
    Body   string `json:"body"`
    UserId int    `json:"userId"`
}

payloadObj := RequestPayload{
    Id:     12,
    Title:  "fetch_test",
    Body:   "my_body",
    UserId: 33,
}

payloadBytes, err := json.Marshal(payloadObj)

if err != nil {
    panic(err)
}

reader := bytes.NewReader(payloadBytes)

res, err := fetch.Fetch(fetch.FetchArgs{
    Method: "PUT",
    Params: map[string]string{
        "Content-type": "application/json; charset=UTF-8",
    },
    Url:     "https://jsonplaceholder.typicode.com/posts/3",
    Body:    reader,
    Timeout: time.Second * 10,
})

if err != nil {
    panic(err)
}

if res.StatusCode == http.StatusOK {
    bodyBytes, err := io.ReadAll(res.Body)
    if err != nil {
        log.Fatal(err)
    }
    bodyString := string(bodyBytes)
    t.Log(bodyString)
}
```