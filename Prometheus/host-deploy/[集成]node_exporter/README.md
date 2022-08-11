### **部署Node-exporter（slave）**
##### 1> 获取node-exporter二进制安装包
```
wget https://github.com/prometheus/node_exporter/releases/download/v1.4.0-rc.0/node_exporter-1.4.0-rc.0.linux-amd64.tar.gz
```
```
tar xzvf node_exporter-1.4.0-rc.0.linux-amd64.tar.gz \
&& mv node_exporter-1.4.0-rc.0.linux-amd64 /usr/local/ \
&& cd /usr/local/node_exporter-1.4.0-rc.0.linux-amd64
```
##### 2> 运行
```
./node_exporter
```
##### 3> 转为系统服务
```
vim /etc/systemd/system/node_exporter.service
```
```
[Unit]
Description=node_exporter
Documentation=https://prometheus.io/
After=network.target
[Service]
ExecStart=/usr/local/node_exporter-1.4.0-rc.0.linux-amd64/node_exporter
WorkingDirectory=/usr/local/node_exporter-1.4.0-rc.0.linux-amd64/
Restart=on-failure
[Install]
WantedBy=multi-user.target
```
```
systemctl start node_exporter
systemctl status node_exporter
systemctl enable node_exporter
```
##### 4> 集成进prometheus（master节点）
```
vim /usr/local/prometheus-2.37.0.linux-amd64/prometheus.yml
```
```
scrape_configs:
  # 采集node exporter监控数据
  - job_name: 'node'
    static_configs:
      - targets: ['152.32.170.211:9100']
```
##### 5> 查看端口
```
netstat -npal|grep 9100
```
##### 6> 访问
```
http://152.32.170.211:9100/metrics
```
