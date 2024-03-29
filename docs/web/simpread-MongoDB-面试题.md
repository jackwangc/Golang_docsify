> 本文由 [简悦 SimpRead](http://ksria.com/simpread/) 转码， 原文地址 https://www.cnblogs.com/angle6-liu/p/10791875.html

### 1. 什么是 MongoDB

```
MongoDB是一个文档数据库，提供好的性能，领先的非关系型数据库。采用BSON存储文档数据。
BSON（）是一种类json的一种二进制形式的存储格式，简称Binary JSON.
相对于json多了date类型和二进制数组。
```

### 2.MongoDB 的优势有哪些

*   面向文档的存储：以 JSON 格式的文档保存数据。
*   任何属性都可以建立索引。
*   复制以及高可扩展性。
*   自动分片。
*   丰富的查询功能。
*   快速的即时更新。

### 3 什么是数据库

```
数据库可以看成是一个电子化的文件柜,用户可以对文件中的数据运行新增、检索、更新、删除等操作。数据库是一个
所有集合的容器，在文件系统中每一个数据库都有一个相关的物理文件。
```

### 4. 什么是集合 (表)

```
集合就是一组 MongoDB 文档。它相当于关系型数据库（RDBMS）中的表这种概念。集合位于单独的一个数据库中。
一个集合内的多个文档可以有多个不同的字段。一般来说，集合中的文档都有着相同或相关的目的。
```

### 5 什么是文档 (记录)

```
文档由一组key value组成。文档是动态模式,这意味着同一集合里的文档不需要有相同的字段和结构。在关系型
数据库中table中的每一条记录相当于MongoDB中的一个文
```

### 6 MongoDB 和关系型数据库术语对比图

![](https://img2018.cnblogs.com/blog/1521877/201904/1521877-20190429170250020-1693717595.png)

### 7. 什么是非关系型数据库

```
非关系型数据库的显著特点是不使用SQL作为查询语言，数据存储不需要特定的表格模式。
```

### 8 为什么用 MOngoDB？

*   架构简单
*   没有复杂的连接
*   深度查询能力, MongoDB 支持动态查询。
*   容易调试
*   容易扩展
*   不需要转化 / 映射应用对象到数据库对象
*   使用内部内存作为存储工作区, 以便更快的存取数据。

### 9 在哪些场景使用 MongoDB

*   大数据
*   内容管理系统
*   移动端 Apps
*   数据管理

###  10 MongoDB 中的命名空间是什么意思?

[![](http://common.cnblogs.com/images/copycode.gif)](javascript:void(0); "复制代码")

```
mongodb存储bson对象在丛集(collection)中.数据库名字和丛集名字以句点连结起来叫做名字空间(namespace).

　　一个集合命名空间又有多个数据域(extent)，集合命名空间里存储着集合的元数据，比如集合名称，集合的
第一个数据域和最后一个数据域的位置等等。而一个数据域由若干条文档(document)组成，每个数据域都有一个
头部，记录着第一条文档和最后一条文档的为知，以及该数据域的一些元数据。extent之间，document之间通过
双向链表连接。

索引的存储数据结构是B树，索引命名空间存储着对B树的根节点的指针。
```

[![](http://common.cnblogs.com/images/copycode.gif)](javascript:void(0); "复制代码")

### 11 monogodb 中的分片什么意思

```
分片是将数据水平切分到不同的物理节点。当应用数据越来越大的时候，数据量也会越来越大。当数据量增长
时，单台机器有可能无法存储数据或可接受的读取写入吞吐量。利用分片技术可以添加更多的机器来应对数据量增加
以及读写操作的要求。
```

### 12 为什么要在 MongoDB 中使用分析器

```
mongodb中包括了一个可以显示数据库中每个操作性能特点的数据库分析器.通过这个分析器你可以找到比预期慢
的查询(或写操作);利用这一信息,比如,可以确定是否需要添加索引.
```

### 13 .MongoDB 支持主键外键关系吗

```
默认MongoDB不支持主键和外键关系。 用Mongodb本身的API需要硬编码才能实现外键关联，不够直观且难度
较大
```

###  14 MongoDB 支持哪些数据类型

*   String
*   Integer
*   Double
*   Boolean
*   Object
*   Object ID
*   Arrays
*   Min/Max Keys
*   Datetime
*   Code
*   Regular Expression 等

### 15 为什么要在 MongoDB 中用 "Code" 数据类型

```
"Code"类型用于在文档中存储 JavaScript 代码。
```

### 16 为什么要在 MongoDB 中用 "Regular Expression" 数据类型

```
"Regular Expression"类型用于在文档中存储正则表达式
```

### 17 为什么在 MongoDB 中使用 "Object ID" 数据类型

```
"ObjectID"数据类型用于存储文档id
```

### 18"ObjectID" 有哪些部分组成

```
一共有四部分组成:时间戳、客户端ID、客户进程ID、三个字节的增量计数器
```

### 19 在 MongoDb 中什么是索引

```
索引用于高效的执行查询,没有索引的MongoDB将扫描整个集合中的所有文档,这种扫描效率很低,需要处理大量
的数据.
    索引是一种特殊的数据结构,将一小块数据集合保存为容易遍历的形式.索引能够存储某种特殊字段或字段集的
值,并按照索引指定的方式将字段值进行排序.
```

### 20 如何添加索引

```
使用db.collection.createIndex()在集合中创建一个索引
```

### 21. 如何查询集合中的文档

```
db.collectionName.find({key:value})
```

### 22 用什么方法可以格式化输出结果

```
db.collectionName.find().pretty()
```

### 23 如何使用 "AND" 或 "OR" 条件循环查询集合中的文档

[![](http://common.cnblogs.com/images/copycode.gif)](javascript:void(0); "复制代码")

```
db.mycol.find(
   {
      $or: [
         {key1: value1}, {key2:value2}
      ]
   }
).pretty()
```

[![](http://common.cnblogs.com/images/copycode.gif)](javascript:void(0); "复制代码")

### 24 更新数据

```
db.collectionName.update({key:value},{$set:{newkey:newValue}})
```

### 25 如何删除文档

```
db.collectionName.remove({key:value})
```

### 26 在 MongoDB 中如何排序

并使用 1 和 -1 来指定排序方式，其中 1 表示升序，而 -1 表示降序。

```
db.connectionName.find({key:value}).sort({columnName:1})
```

### 27 什么是聚合

　　聚合操作能够处理数据记录并返回计算结果。聚合操作能将多个文档中的值组合起来，对成组数据执行各种操作，返回单一的结果。它相当于 SQL 中的 count(*) 组合 group by。对于 MongoDB 中的聚合操作，应该使用`aggregate()`方法。

```
db.COLLECTION_NAME.aggregate(AGGREGATE_OPERATION)
```

### 28 在 MongoDB 中什么是副本集（避免单点故障）

```
在MongoDB中副本集由一组MongoDB实例组成，包括一个主节点多个次节点，MongoDB客户端的所有数据都
写入主节点(Primary),副节点从主节点同步写入数据，以保持所有复制集内存储相同的数据，提高数据可用性。
```

### 29 什么是 NoSQL 数据库？NoSQL 和 RDBMS 有什么区别？在哪些情况下使用和不使用 NoSQL 数据库？

[![](http://common.cnblogs.com/images/copycode.gif)](javascript:void(0); "复制代码")

```
NoSQL是非关系型数据库，NoSQL = Not Only SQL。
  关系型数据库采用的结构化的数据，NoSQL采用的是键值对的方式存储数据。
  在处理非结构化/半结构化的大数据时；在水平方向上进行扩展时；随时应对动态增加的数据项时可以优先考虑
使用NoSQL数据库。

  在考虑数据库的成熟度；支持；分析和商业智能；管理及专业性等问题时，应优先考虑关系型数据库。
```

[![](http://common.cnblogs.com/images/copycode.gif)](javascript:void(0); "复制代码")

### 30 MongoDB 支持存储过程吗？如果支持的话，怎么用？

```
MongoDB支持存储过程，它是javascript写的，保存在db.system.js表中。
```

### 31 如何理解 MongoDB 中的 GridFS 机制，MongoDB 为何使用 GridFS 来存储文件？

　　GridFS 是一种将大型文件存储在 MongoDB 中的文件规范。使用 GridFS 可以将大文件分隔成多个小文档存放，这样我们能够有效的保存大文档，而且解决了 BSON 对象有限制的问题。

### 32 为什么 MongoDB 的数据文件很大？

　　MongoDB 采用的预分配空间的方式来防止文件碎片。

### 33 当更新一个正在被迁移的块（Chunk）上的文档时会发生什么？

　　更新操作会立即发生在旧的块（Chunk）上，然后更改才会在所有权转移前复制到新的分片上。

### 34 MongoDB 在 A:{B,C} 上建立索引，查询 A:{B,C} 和 A:{C,B} 都会使用索引吗？

　　不会，只会在 A:{B,C} 上使用索引。

### 35 mongodb 成为最好 nosql 数据库的原因是什么?

　　面向文件的 高性能 高可用性 易扩展性 丰富的查询语言

### 36 如果用户移除对象的属性, 该属性是否从存储层中删除?

　　是的, 用户移除属性然后对象会重新保存 (re-save()).

### 37 允许空值 null 吗?

　　对于对象成员而言, 是的. 然而用户不能够添加空值 (null) 到数据库丛集 (collection) 因为空值不是对象. 然而用户能够添加空对象{}.

### 38 更新操作立刻 fsync 到磁盘?

　　不会, 磁盘写操作默认是延迟执行的. 写操作可能在两三秒 (默认在 60 秒内) 后到达磁盘. 例如, 如果一秒内数据库收到一千个对一个对象递增的操作, 仅刷新磁盘一次.

### 39 如何执行事务 / 加锁?

　　mongodb 没有使用传统的锁或者复杂的带回滚的事务, 因为它设计的宗旨是轻量, 快速以及可预计的高性能. 可以把它类比成 mysql mylsam 的自动提交模式. 通过精简对事务的支持, 性能得到了提升, 特别是在一个可能会穿过多个服务器的系统里.

### 40  启用备份故障恢复需要多久?

　　从备份数据库声明主数据库宕机到选出一个备份数据库作为新的主数据库将花费 10 到 30 秒时间. 这期间在主数据库上的操作将会失败–包括写入和强一致性读取 (strong consistent read) 操作. 然而, 你还能在第二数据库上执行最终一致性查询(eventually consistent query)(在 slaveok 模式下), 即使在这段时间里.

### 41  什么是 master 或 primary?

　　它是当前备份集群 (replica set) 中负责处理所有写入操作的主要节点 / 成员. 在一个备份集群中, 当失效备援 (failover) 事件发生时, 一个另外的成员会变成 primary.

### 42  我应该启动一个集群分片 (sharded) 还是一个非集群分片的 mongodb 环境?

　　(数据量大用集群分片, 数据量小用非集群)

　　为开发便捷起见, 我们建议以非集群分片 (unsharded) 方式开始一个 mongodb 环境, 除非一台服务器不足以存放你的初始数据集. 从非集群分片升级到集群分片 (sharding) 是无缝的, 所以在你的数据集还不是很大的时候没必要考虑集群分片(sharding).

### 43 分片 (sharding) 和复制 (replication) 是怎样工作的?

　　每一个分片 (shard) 是一个分区数据的逻辑集合. 分片可能由单一服务器或者集群组成, 我们推荐为每一个分片 (shard) 使用集群.

### 44 数据在什么时候才会扩展到多个分片 (shard) 里?

　　mongodb 分片是基于区域 (range) 的. 所以一个集合 (collection) 中的所有的对象都被存放到一个块 (chunk) 中. 只有当存在多余一个块的时候, 才会有多个分片获取数据的选项. 现在, 每个默认块的大小是 64mb, 所以你需要至少 64 mb 空间才可以实施一个迁移.

### 45 当我试图更新一个正在被迁移的块 (chunk) 上的文档时会发生什么?

　　更新操作会立即发生在旧的分片 (shard) 上, 然后更改才会在所有权转移 (ownership transfers) 前复制到新的分片上.

### 46 如果在一个分片 (shard) 停止或者很慢的时候, 我发起一个查询会怎样?

　　如果一个分片 (shard) 停止了, 除非查询设置了 “partial” 选项, 否则查询会返回一个错误. 如果一个分片 (shard) 响应很慢, mongodb 则会等待它的响应.

### 47 可以把 movechunk 目录里的旧文件删除吗?

　　没问题, 这些文件是在分片 (shard) 进行均衡操作 (balancing) 的时候产生的临时文件. 一旦这些操作已经完成, 相关的临时文件也应该被删除掉. 但目前清理工作是需要手动的, 所以请小心地考虑再释放这些文件的空间.

### 48 如果块移动操作 (movechunk) 失败了, 我需要手动清除部分转移的文档吗?

 　　不需要, 移动操作是一致 (consistent) 并且是确定性的(deterministic); 一次失败后, 移动操作会不断重试; 当完成后, 数据只会出现在新的分片里(shard).

### 49 mongodb 是否支持事务

　　MongoDB 4.0 的新特性——事务（Transactions）：MongoDB 是不支持事务的，因此开发者在需要用到事务的时候，不得不借用其他工具，在业务代码层面去弥补数据库的不足。

　　事务和会话 (Sessions) 关联，一个会话同一时刻只能开启一个事务操作，当一个会话断开，这个会话中的事务也会结束。