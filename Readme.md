
### go-kavenegar

A Golang SDK for [kavenegar](https://kavenegar.com) API.

All the REST APIs listed in [kavenegar API document](https://kavenegar.com/rest.html) are implemented.

For best compatibility, please use Go >= 1.22.

Make sure you have read kavenegar API document before continuing.

### Installation

```shell
go get github.com/parparvaz/kavenegar-sdk-golang
```

### REST API

#### Setup

Init client for API services. Get ApiKey from your kavenegar account.

```golang
var apiKey = "your api key"
client := kavenegar.NewClient(apiKey)
```

A service instance stands for a REST API endpoint and is initialized by client.NewXXXService function.

Simply call API in chain style. Call Do() in the end to send HTTP request.

If you have any questions, please refer to the specific reference definitions or usage methods

##### Proxy Client

```
proxyUrl := "http://127.0.0.1:7890" // Please replace it with your exact proxy URL.
client := binance.NewProxyClient(apiKey, proxyUrl)
```


#### Send Lookup SMS

```golang
res, err := client.NewLookupService().
	Receptor("mobile number").
	Token("token").
	Template("template").
	Do(context.Background())
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(res)

```
