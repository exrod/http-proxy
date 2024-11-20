

### A Small http forward Proxy Server in Go Lang

Idea Taken from : https://gist.github.com/yowu/f7dc34bd4736a65ff28d

![license](https://img.shields.io/badge/license-MIT-red)
[![discord](https://img.shields.io/discord/987749138743582811?color=7289da&logo=discord&logoColor=white)](https://discord.gg/ghouls)

### How to Run

```bash
git clone https://github.com/exrod/http-proxy
cd http-proxy
```
```bash
go mod tidy
go run main.go
``` 

### Usage
- Start the proxy: `go run main.go`
- Open your web browser to some website



### Note:

Go's HTTP server implementation is truly remarkable, showcasing a clean and efficient design that I’ve taken the time to thoroughly study. Its reliability and performance stand out, making it a solid choice for building web applications or services. That said, one feature I find noticeably absent is the ability to handle multiple pipelined HTTP requests in parallel. This addition could further elevate its utility by improving performance in scenarios where such pipelining is beneficial, especially for applications requiring high throughput or working with clients that rely heavily on HTTP pipelining.

On the client side, Go’s HTTP library has impressed me with its simplicity and ease of use. Even with my relatively limited experience using it in the context of this proxy project, it has proven to be highly intuitive and well-suited for common use cases. However, I haven’t yet had the opportunity to dive into the underlying implementation of the HTTP client. Understanding its design and internal mechanics would likely provide valuable insights into how it achieves its balance of simplicity and functionality. Overall, my experience with Go’s HTTP tools has been positive, with only minor gaps that, if addressed, could make them even more powerful and versatile

## Star History

<a href="https://star-history.com/#exrod/http-proxy&Date">
 <picture>
   <source media="(prefers-color-scheme: dark)" srcset="https://api.star-history.com/svg?repos=exrod/http-proxy&type=Date&theme=dark" />
   <source media="(prefers-color-scheme: light)" srcset="https://api.star-history.com/svg?repos=exrod/http-proxy&type=Date" />
   <img alt="Star History Chart" src="https://api.star-history.com/svg?repos=exrod/http-proxy&type=Date" />
 </picture>
</a>
