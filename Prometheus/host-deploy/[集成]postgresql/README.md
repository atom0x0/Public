### **部署PostgreSQL**
#### 1> 拉取PostgreSQL镜像
```
docker pull postgres:12.1
```
#### 2> 构建volume存储卷
```
docker volume create dv_pgdata
```
