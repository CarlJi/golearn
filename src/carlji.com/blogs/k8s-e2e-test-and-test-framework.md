## 前言

Kubernetes的成功少不了大量工程师的共同参与，而他们之间如何高效的协作的，非常值得我们探究。最近看了他们的e2e测试和框架，还是挺有启发的。

## 怎样才是好的e2e测试？

不同的人写出的测试用例千差万别，尤其这些大多还是开发人员写的，他们可能并没有经历过大量测试用例场景的熏陶。而Kubernetes社区就非常聪明，他们抽象出来了一些共性的东西，来希望大家遵守。比如说

1. 拒绝“flaky”测试 - 也就是那些偶尔会失败，但是又非常难定位的问题。
2. 错误输出要详细，尤其是做断言时，相关信息要有。不过也不要打印太多无效信息，尤其是在case并未失败的情况。
3. make case run in anywhere。这一点很重要，因为你的case是提交到社区，可能在各种环境下，各种时间段内运行。面对着各种cloud provider，各种系统负载情况。所以你的case要尽可能稳定，比如APICall，能异步的，就不要假设是同步; 比如多用retry机制等。
4. 测试用例要执行的足够快。超过两分钟，就需要给这种测试打上[SLOW]标签。而有这种标签的测试用例，可以运行的场景就比较有限制了。谁又不希望自己写的用例都被尽可能的执行呢？很有激励性的一条规则。

规则已经定下了，但是如何让大家方便的遵守呢，好在Kubernetes社区开发一些列的基础设施，以及类库来帮助开发者使用。这些我们接下来会细讲。

## e2e 验收测试

搞过测试的应该都知道，在面对复杂系统测试时，我们通常有多套测试环境，但是测试代码通常只有一份。所以为了能更好的区分测试用例，我们通常采取打标签的方式来给用例分类。而在Kubernetes的e2e里，这也不例外。

Kubernetes默认将测试用例分为下面几类，需要开发者在实际开发用例时，打上合适的标签。

- 没标签的，默认测试用例是稳定的，支持并发，且运行足够快的
- [Slow] 执行比较慢的用例.(对于具体的时间阈值，Kubernetes不同的文档表示不一致，此处需要修复)
- [Serial] 不支持并发的测试用例，比如占用太多资源，还比如需要重启Nodes的
- [Disruptive] 会导致其他测试用例失败或者被破坏的测试用例
- [Flaky] 不稳定的用例，且很难修复。使用它要非常慎重，因为常规CI job并不会运行这些测试用例
- [Feature:.+] 围绕特定非默认Kubernetes集群功能或者非核心功能的测试用例

有了这些标签，那么我们在实际使用时就可以灵活的选择我们要跑的测试用例了。

比如针对任意Kubernetes集群,做Conformance验收测试，就可以通过下面的简单几步来执行：

```
# under kubernetes folder, compile test cases and ginkgo tool
make WHAT=test/e2e/e2e.test && make ginkgo

# setup for conformance tests
export KUBECONFIG=/path/to/kubeconfig
export KUBERNETES_CONFORMANCE_TEST=y
export KUBERNETES_PROVIDER=skeleton

# run all conformance tests
go run hack/e2e.go -v --test --test_args="--ginkgo.focus=\[Conformance\]"
```

## Kubernetes e2e test framework

上面说Kubernetes社区为了让开发者写的出的e2e能尽可能的遵循现有的一些规则，精心维护了一套测试框架，以及一些好用的Util类。



## 普适性的Kubernetes e2e framework



```

```



## 知识点总结




## 参考文档

- https://github.com/thtanaka/kubernetes/blob/master/docs/devel/writing-good-e2e-tests.md
- https://github.com/thtanaka/kubernetes/blob/master/docs/devel/e2e-tests.md

## Contact me ?

Email: jinsdu@outlook.com

Blog: <http://www.cnblogs.com/jinsdu/>

Github: <https://github.com/CarlJi>

------

> 童鞋，如果觉得本文还算用心，还算有用，何不点个赞呢(⊙o⊙)？