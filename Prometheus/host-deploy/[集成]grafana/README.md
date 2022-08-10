### **部署Grafana（master）**
##### 1> 获取grafana二进制安装包
```
wget https://dl.grafana.com/enterprise/release/grafana-enterprise_9.0.6_amd64.deb
```
```
dpkg -i grafana-enterprise_9.0.6_amd64.deb
```
##### 2> 开启服务
```
systemctl start grafana-server
systemctl status grafana-server
systemctl enable grafana-server
```
##### 3> 查看端口
```
netstat -npal|grep 3000
```
##### 4> 访问
```
http://152.32.170.211:3000/
```
