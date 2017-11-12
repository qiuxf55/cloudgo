# Cloudgo

- ## 框架使用
在这里使用的是beego框架

Beego是一个快速开发Go应用的http框架，可用于快速开发Api、web及后端服务等各种应用，是一个RESTful框架。

```
go get github.com/astaxie/beego
```

- ## 运行程序


```
go run http.go -port 9090
```

- ### 打开网页
![image](http://i4.bvimg.com/617453/b90328cc5eac6092.png)

- ### 终端截图

![Markdown](http://i4.bvimg.com/617453/f7b311950b3f33b2.png)

- ### 测试


```
curl -v http://localhost:9090/hello/qiuxf
```

![Markdown](http://i4.bvimg.com/617453/3fd9ab8a116c5885.png)


- ## 压力测试

![Markdown](http://i1.bvimg.com/617453/a578bba0031476aa.png)
![Markdown](http://i1.bvimg.com/617453/bf78818b23f82367.png)

完成50%的请求耗时6ms

完成100%的请求耗时24ms


- ## 参数详解
- 

-n 1000 表示总请求数位1000

-c 100表示并发用户数为100

http://localhost:9090/hello/qiuxf 表示这些请求的目标URL。


Server Software
表示被测试的Web服务器软件名称，这里是beegoServer:1.9.0


Server Hostname
表示请求的URL中的主机部分名称，它来自于http请求数据的头信息.

Server Port
表示被测试的Web服务器软件的监听端口.

Document Path
表示请求的URL中根绝对路径，它同样来自于http请求数据的头信息。

Document Length
表示http响应数据的正文长度。

Concurrency Level
表示并发用户数，这是我们设置的参数。

Time taken for tests
表示所有这些请求被处理完成花费的总时间。顺便提一下，某些Apache版本如2.2.4附带的ab，对于这一统计项存在一些计算上的bug，当总请求数较少时，其统计的总时间会无法小于0.1s。

Complete requests
表示总请求数，这是我们设置的相应参数。

Failed requests
表示失败的请求数，这里的失败是指请求的连接服务器、发送数据、接收数据等环节发生异常，以及无响应后超时的情况。对于超时时间的设置可以用ab的-t参数。
而如果接受到的http响应数据的头信息中含有2xx以外的状态码，则会在测试结果显示另一个名为“Non-2xx responses”的统计项，用于统计这部分请求数，这些请求并不算是失败的请求。

Total transferred
表示所有请求的响应数据长度总和，包括每个http响应数据的头信息和正文数据的长度。注意这里不包括http请求数据的长度，所以Total
transferred代表了从Web服务器流向用户PC的应用层数据总长度。通过使用ab的-v参数即可查看详细的http头信息。

HTML transferred
表示所有请求的响应数据中正文数据的总和，也就是减去了Total transferred中http响应数据中头信息的长度。

Requests per second
这便是我们重点关注的吞吐率，它等于：
Complete requests / Time taken for tests

Time per request
这便是前面提到的用户平均请求等待时间，它等于:
Time taken for tests / (Complete requests /Concurrency Level)

Time per request(across all concurrent requests)
这便是前面提到的服务器平均请求处理时间，它等于：
Time taken for tests / Complete requests

这正是吞吐率的倒数。同时，它也等于：
Time per request / Concurrency Level

Transfer rate
表示这些请求在单位时间内从服务器获取的数据长度，它等于：
Total transferred / Time taken for tests
这个统计项可以很好的说明服务器在处理能力达到限制时，其出口带宽的需求量。
