# 基于Golang和Prometheus的性能测试套件

## 目标:
- [x] 支持并发, 提供灵活的并发接口，方便使用
      - Run(worker, func(), taskName)
      - RunUtil(worker, func(), taskName, endTime)
- [ ] 并发方法自动支持记录性能数据，并上报到PushGateway, 包括但不限于
    a. TPS
    b. 任务执行时间的95值
    c. 任务执行的成功率
    d. 总的worker数
- [ ] 支持单独上报metric到Prometheus
      PushMetric(metric)
- [ ] 标准prometheus go客户端内置的metric，作为可选项      
- [ ] 支持动态调整worker数
- [ ] 支持动态调整执行时间，任务间隔时间

## 最佳实践
提供在Kubernetes里运行的例子:
- [ ] 编写实际性能测试例子
- [ ] 编写例子的dockerfile
- [ ] 通过k8s的deployment部署该测试程序 
- [ ] 给出 TPS, 成功率，响应时间95 值的PromQL