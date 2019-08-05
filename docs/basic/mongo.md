# MongoDB 操作

## 1.MongoDB 概述

MongoDB 是一款跨平台, 面向文档的数据库。特点：高性能，高可用性，可扩展。两个主要概念：集合和文档

一个 MongoDB 服务器有多个数据库，数据库是集合的实际容器，每个数据库都在文件系统有自己的一组文件

集合是一组 MongoDB 文档。相当于关系型数据库 (RDBMS) 中表的概念。集合位于单独的一个数据库中。集合中的多个文档可以有多个不同的字段。

文档是一组键值对。文档有动态的模式，文档不需要同样的字段和结构。 相当于关系型数据库中的行。字段相当于列。

MongoDB 适用领域：大数据，内容管理及交付，移动和社会化基础设施，用户数据管理，数据中心


---

## 2.MongoDB 安装


---


## 3.MongoDB 主要操作

### 3.1 创建数据库

use + 数据库名称 创建数据库，use 命令创建一个新的数据库，如果数据库存在，则返回该数据库 `use DATABASE_NAME`


```
use mydb      # 创建数据库或返回数据库
db            # 检查查看当前数据库
show dbs      # 检查数据库列表（空表没有数据），mongodb 默认数据库时test
```


`db.dropDatabase()` 删除选定数据库，没有选定会将默认的test 数据库删除

```
use mydb              # 删除数据库前要先返回该数据库
db.dropDatabase()
```

### 3.2 创建集合

`db.createCollection( name, options)` name 为所要创建的集合名称，options 可选，指定集合配置的文档。

```
use test

db.createCollection("myCollection")

show collletioS           # 查看创建集合

---------
db.createCollection("mycol",{capped:true,autoIndexID:true,size:879,max:1000})

# 集合配置参数意义  max 固定集合中包含文档的最大数量 size 为固定集合指定最大值
# capped 创建固定集合，达到最大值覆盖 autoIndexID 自动在 _id 字段创建索引

---------
db.newCollectionName.insert({"name":"jack"}) # 插入文档时，集合自动创建
db.collectionname.drop()   # 删除集合，成功删除返回true

```

### 3.3 MongoDB 数据类型

数据类型 | 类型 | 数据类型 | 类型 |数据类型 | 类型
---|---|---|---|---|---
String | 字符串(utf8) | Integer | 整型数值(32)(64) |Boolean | 布尔值(真假)
Double | 双精度浮点型 | Min/Max keys | 对比元素最小最高值 | Arrays | 数组列表存储为一个键
Timestamp | 记录文档修改时间 | Objext | 用于内嵌文档那个 | Null | 创建空值
Date | 日期时间 | ObjextID | 创建文档ID | Code | 在文档中存储js代码

### 3.3 MongDB 增删改查

#### 插入( 增加)

`db.collection_name.insert(document)  db.collection_name.copy(document)`
文档id(ObjextID) 每个文档都一个唯一ID 如果没有指定ID mongodb 会自动为文档创建一个唯一的ID ，save 方法会覆盖指定id的全部数据

#### 查询

`db.collection_name.find() # 非结构化方式显示所有文档那个 findOne() 只返回一个文档`

`db.collection_name.find().pretty() # 以结构化方式显示结构 `

```
条件查询『字段：条件』 条件：{ 判断 : value }
db.mycol.find({"by":123}).pretty() # ： 字段 等于；类似 where by = 123
db.mycol.find({"likes":{$lt:123}}).pretty()   # $lt 小于 ；类似 where likes < 123
db.mycol.find({"likes":{$lte:123}}).pretty()  # $lte 小于或等于 ；类似 where likes <= 123
db.mycol.find({"likes":{$gt:123}}).pretty()   # $gt 大于 ； 类似 where > 123;
db.mycol.find({"likes":{$gte:123}}).pretty()  # $gte 大于或等于；类似 where >= 123;
db.mycol.find({"likes":{$ne:123}}).pretty()   # $ne 不等于 ；类似 where ！= 123
db.mycol.find({key1:value1,key2:value2}).pretty()  # and 语句 where　key1=value1 and key2=value2
db.mycol.find($or[key1:value1,key2:value2]) # or 语句 
db.mycol.find({key1:value1,$or:[{}{}]})  # and or 结合
```

#### 更新(改)

update() 更新已有文档中的值

```
db.mycol.update({"title":"jacke"},{$set:{"title":"newjacke"}})   # 
db.mycol.update({"title":"jacke"},{$set:{"title":"newjacke"},{multi:true})  # 更新多个文档

db.mycol.save({"_id":value,"key1":newValue})  # 替换原有文档，覆盖全部数据

```

#### 删除

remove() 方法 清除集合中的文档

```
db.collection_name.remove(DELLETION_CRITTERIA,justOne) 
# justOne:设置只删除一个文档 true or 1

db.mycol.remove({"key":"value"})
# 将集合中符合标准所有文档删除
```

#### 映射

MongoDB 映射： 指的是只选择文档中的必要数据

` db.COLLECTION_NAME.find({},{key:1}) # key=字段 1，显示；0，隐藏字段;在查询语句后面添加`

#### 限制记录

`db.mycol.find({},{}).limit(2)   # 限制查询显示文档个数`

`db.mycol.find({},{}).limit(1).skip(1)  # skip 跳过的数量条数`

`db.mycol.find({},{}).sort("key":-1)  # sort 1表示升序，-1表示降序`

## 4.索引

索引是一种特殊的数据结构，能够实现高效查询。

` db.COLLECTION_NAME.ensureIndex({KEY:1})  # 1 代表升序；-1 代表降序 backgroud:true,false 在后台建立索引`

## 5. 聚合操作

聚合能够处理数据记录并返回结果

```
db.COLLECTION_NAME.aggregate(aggregate_operation)  # 聚合操作= count(*) group by
db.mycol.aggregate([$group:{_id:"$by_user",num_tutorial:{$sum:1}}])
==> select by_user,count(*) from mycol group by by_user
```

表达式 | 描述 | 表达式 | 描述 | 表达式 | 描述 | 表达式 | 描述
---|---|---|---|---|---|---|---
$sum | 计算总和 | $min | 最小值 | $last | 根据排序获取最后一个 | $push | 在结果文档中插入值到一个数组
$avg | 计算平均值| $max | 最大值 | $first | 第一个 | $addToSet | 在结果文档中插入值到一个数组，不创建副本





















