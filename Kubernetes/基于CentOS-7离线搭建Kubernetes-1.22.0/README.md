# 《基于CentOS-7离线部署Kubernetes-1.22.0》

###### 声明：工具引用自：https://www.sealyun.com/

### 1、解压安装包

```
tar xzvf setup-kube.tar.gz -C /root/
```

### 2、配置sealos

```
chmod a+x /root/setup-kube/sealos \
&& cp /root/setup-kube/sealos /usr/bin
```

### 3、三节点配置

```text
| Host           | Role   | /etc/hostname | /etc/hosts            |
| ------ | ------ | ------ | ------ |
| 192.168.32.129 | Master | master        | 192.168.32.129 master
										192.168.32.130 slave-1
										192.168.32.131 slave-2
| 192.168.32.130 | Slave  | slave-1       |                       |
| 192.168.32.131 | Slave  | slave-2       |                       |
```

### 3、部署Kubernetes集群

```
sealos init --passwd "root" \
--master 192.168.32.129 \
--node 192.168.32.130 \
--node 192.168.32.131 \
--pkg-url /root/setup-kube/kube1.22.0.tar.gz \
--version 1.22.0
```

