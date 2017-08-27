## 前言

Kubernetes的成功少不了大量工程师的共同参与，而他们之间如何高效的协作，更是建立在一些列的自动化系统基础上的。最近看了他们的e2e测试和框架，还是挺有启发的。

## 怎样才是好的e2e测试？

不同的人写出的测试用例千差万别，尤其这些大多还是开发人员写的，他们可能并没有经历过大量测试用例的熏陶。而Kubernetes社区就非常聪明，他们抽象出来了一些共性的东西，来希望大家遵守。比如说

1. 拒绝“flaky”测试 - 也就是那些偶尔会失败，但是又非常难定位的问题。
2. 错误输出要详细，尤其是做断言时，相关信息要有。不过也不要打印太多无效信息，尤其是在case并未失败的情况。
3. make case run in anywhere。这一点很重要，因为你的case是提交到社区，可能在各种环境下，各种时间内运行。面对着各种cloud provider，各种系统负载情况。所以你的case要尽可能稳定，比如APICall，能异步的，就不要假设是同步; 多用retry机制等。
4. 测试用例要执行的足够快。超过两分钟，就需要给这种测试打上[SLOW]标签。而有这种标签的测试用例，可以运行的场景的比较有限制了。谁又不希望自己写的用例都被尽可能的执行呢？很有激励性的一条规则。

所以在追求测试覆盖的前提下，

## e2e 验收测试





## Kubernetes e2e framework

说了好的e2e测试，我们来看kubernetes是如何来实现的



## 普适性的Kubernetes e2e framework



```

```



## 知识点总结

上面我们共提到了五种并发模式：

- 简单并发模型
- 规定时间内的持续并发模型
- 基于大数据量的持续并发模型
- 等待异步任务结果模型
- 定时反馈异步任务结果模型

归纳下来其核心就是使用了Go的几个知识点：Goroutine, Channel, Select, Time, Timer/Ticker, WaitGroup. 若是对这些不清楚，可以自行Google之。

另完整的Example 代码可以参考这里：https://github.com/jichangjun/golearn/blob/master/src/carlji.com/experiments/concurrency/main.go

使用方式： go run main.go <场景>

比如 :



## 参考文档

- https://github.com/thtanaka/kubernetes/blob/master/docs/devel/writing-good-e2e-tests.md
- ​

这篇是Google官方推荐学习Go并发的资料，从初学者到进阶，内容非常丰富，且权威。

## Contact me ?

Email: jinsdu@outlook.com 

Blog: <http://www.cnblogs.com/jinsdu/>

Github: <https://github.com/CarlJi>

------

> 童鞋，如果觉得本文还算用心，还算有用，何不点个赞呢(⊙o⊙)？