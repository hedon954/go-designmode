# Options 模式



## 1. 痛点

我们要在项目里封装一个通用的发 HTTP 请求的工具函数，它的参数可能会有哪些呢？因为是工具函数，要做到通用就必然需要定义很多能配置HTTP客户端的参数，比如：

```go
func HttpRequest(method string, url string, body []byte, headers map[string]string, timeout time.Duration) ...
```

上面这个工具函数，如果只是做GET请求的话，很多HTTP客户端的设置是不需要设置的，而且超时时间我们一般都会设置一个默认的。如果还按普通定义函数的方法来实现的话，函数逻辑里势必会有不少判断空值的逻辑。

```go
if body != nil {
   // 设置请求体Data
  ......
}

if headers != nil {
  // 设置请求头
  ......
}
```

调用的时候，调用者的代码也不得不传一些零值给不需要自定义的配置参数。

```go
HttpRequest('GET', 'https://www.baidu.com', nil, nil, 2 * time.Second)
```

### 1.1 使用配置对象解决

一种常用的解决方案是，工具函数的签名定义时，不再定义各个可能需要配置的可选参数，转而定义一个配置对象。

```go
type HttpClientConfig struct {
  timeout time.Duration
  headers map[string]string
  body    []byte
}

func HttpRequest(method string, url string, config *HttpClientConfig) ...
```

对调用者来说，比上一种方法看起来简洁了不少，如果全都是默认选项只需要给配置对象这个参数传递一个零值即可。

```go
HttpRequest('GET', 'https://www.baidu.com', nil)
```

但是对于函数的实现方来说，仍然少不了那些选项参数非零值的判断，而且因为配置对象在函数外部可以改变，这就有一定几率配置对象在函数内部未被使用前被外部程序改变，真正发生了相关的`BUG`，排查起来会比较头疼。

### 1.2 可变参数

与配置对象方案类似，如果单纯通过可变参数来解决这个问题，也会有不少问题

```go
func HttpRequest(method string, url string, options ...interface{}) ...
```

虽然参数是可变的，但是实现方需要通过遍历设置 HTTP 客户端的不同选项，这就让可变参数 **固定了传递顺序**，调用方如果想要设置某个可选项还得记住参数顺序，切无法直接通过函数签名就确定参数顺序，貌似还不如咱们最原始的解决方案。



## 2. 理解

Options 模式可以让具有多个可选参数的函数或者方法更整洁和好扩展，当一个函数具有五六个甚至十个以上的可选参数时使用这种模式的优势会体现的很明显。

在 gRPC 的 SDK 中，Options 模式出现的频率特别高，比如它是客户端方法就可以传递不少以  `with` 开头的闭包函数：

```go
client, err := grpc.Dial(
   "127.0.0.1:12305",
   grpc.WithInsecure(),
   grpc.WithUnaryInterceptor(...),
   grpc.WithStreamInterceptor(...),
   grpc.WithAuthority(...)
)
```

这些配置方法返回的都是一个名为 `DialOption` 的 `interface`：

```go
type DialOption interface {
 apply(*dialOptions)
}

func WithInsecure() DialOption {
 ...
}
```



## 3. 解决

现在我们就使用 `Options` 模式对我们的工具函数进行一下改造，首先定义一个契约和配置对象。

```go
// requestOption 针对可选的 HTTP 请求配置项
type requestOption struct {
	timeout time.Duration
	data    string
	headers map[string]string
}

type Option struct {
	apply func(*requestOption)
}

// defaultRequestOptions 默认请求选项
func defaultRequestOptions() *requestOption {
	return &requestOption{
		timeout: 5 * time.Second,
		data:    "",
		headers: nil,
	}
}
```

接下来我们要定义的配置函数，每个都会设置请求配置对象里的某一个配置。

```go
func WithTimeout(timeout time.Duration) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.timeout = timeout
		},
	}
}

func WithData(data string) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.data = data
		},
	}
}

func WithHeaders(headers map[string]string) *Option {
	return &Option{
		apply: func(option *requestOption) {
			option.headers = headers
		},
	}
}
```

那么此时我们的工具函数的签名就应用上上面定义的接口契约，其实现里我们只需要遍历 `options` 这个可变参数，调用每个 `Option` 对象的 `apply` 方法对配置对象进行配置即可，不用在担心可变参数的顺序。

```go
// HttpRequest 省略返回值
func HttpRequest(method string, url string, options ...*Option) {
	reqOpts := defaultRequestOptions() // 默认的请求选项
	for _, opt := range options {      // 在reqOpts上应用通过options设置的选项
		opt.apply(reqOpts)
	}
	// 创建请求对象
	req, err := http.NewRequest(method, url, strings.NewReader(reqOpts.data))
	if err != nil {
		// 异常处理
	}

	// 设置请求头
	for key, value := range reqOpts.headers {
		req.Header.Add(key, value)
	}
	// 发起请求

	return
}
```

效果：

```go
HttpRequest("GET", url)
HttpRequest("POST", url, WithHeaders(headers))
HttpRequest("POST", url, WithTimeout(timeout), WithHeaders(headers), WithData(data))
```



## 参考

- [一些实用的编程模式 | Options模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247489777&idx=1&sn=a9c17cc31cb77f9139a45c484057f7ac&chksm=fa80c966cdf74070c095a8578ae7b17ffc51fc381535b175f562cc4af13db8e76ba600fd6f16&token=1449569934&lang=zh_CN&scene=21#wechat_redirect)