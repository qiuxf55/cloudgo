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

![Markdown](http://i4.bvimg.com/617453/e32171b15e580f14.png)


完成50%的请求耗时6ms

完成100%的请求耗时24ms
