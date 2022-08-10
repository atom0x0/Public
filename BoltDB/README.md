# **BoltDB存储引擎**
```
采用B+Tree，基于LMDB及K/V模型实现的单点文件型存储引擎
```
### **0x01 存储机制**
##### 1.1> 名词解释
```
DB：数据库；即一个文件
Bucket：Bolt中存储单元集合的概念，抽象类似RDB中的table
Key/Value：Bolt中存储的最小单元
Cursor：游标
page：页；磁盘概念
node：块；内存概念
COW: Copy On Write；写时复制
LMDB：Light Memory Map Database；内存关系映射型数据库
mmap：Memeoy Map；一种内存映射文件的方法
```
##### 1.2> 存储顺序
```
读盘顺序：  file -> page -> node -> get
落盘顺序：  set -> node -> page -> file
``` 
##### 1.3> page & node 转换
### **0x02 检索机制**
### **0x03 事务机制**
### **0x04 CRUD**
### **优劣**
```
1、K/V模型的存储引擎，简单快速可靠
2、基于内存映射文件存储数据
3、完整支持ACID
4、基于字节切片做存储读取；节省了sql解析转化和结构化开销
```
```
1、原生不支持高阶特性
2、适用于读多写少场景
3、适用于内存大于数据库场景
4、原生不支持分布式，单节点数据库
```