### E2e MAT test framework 


## 如何运行

注意，本框架并未将依赖的包全部导入，需要使用者通过govendor或godep来手动导入。
待依赖包都入库，编译通过之后，就可以通过ginkgo灵活运行了:

 ```
 ginkgo -p -v -r -- "-kubeconfig=PATH_TO_KUBECONFIG"
 ```
 
 如果-kubeconfig参数不提供，会默认使用本地位置~/.kube/config的配置文件来访问k8s集群
 
 ## 编译成独立二进制运行
 
 在e2e目录下执行
 ```
go test -c 
```
默认编译为e2e.test的文件，当然可以使用-o参数来重命名

```apple js
e2e.test -h 
```
e2e.test可以独立运行，当然也可以使用-h变量来查看有哪些参数可以使用