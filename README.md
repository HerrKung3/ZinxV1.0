# ZinxV1.0
基于Golang开发的TCP轻量级并发服务器
基于依赖倒转，单一职责，开闭原则等设计模式，设计了该处理 TCP 长连接的服务器框架，该框架 可以实现应用层的 API 注册，因此具有扩展性强，适用面广的特点；实现了消息队列以及工作池机制， 避免了大量客户端连接后开启 Goroutine 过多导致的性能下降和稳定性下降的问题；实现了简单的负载 均衡方法，避免了工作池中 worker 工作量不均衡的问题；实现了全局配置统一管理以及抽象层、实现 层、应用层的分离，使得框架的维护和扩展更加方便。
git clone到本地后，先后运行ZinxV1.0Test中server和client即可
