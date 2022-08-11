### **部署Blackbox-exporter（slave）**
##### 1> 获取blackbox-exporter二进制安装包
```
wget https://github.com/prometheus/blackbox_exporter/releases/download/v0.22.0/blackbox_exporter-0.22.0.linux-amd64.tar.gz
```
```
tar xzvf blackbox_exporter-0.22.0.linux-amd64.tar.gz \
&& mv blackbox_exporter-0.22.0.linux-amd64 /usr/local/ \
&& cd /usr/local/blackbox_exporter-0.22.0.linux-amd64
```
##### 2> 配置blackbox.yml文件
```
vim /usr/local/blackbox_exporter-0.22.0.linux-amd64/blackbox.yml
```
```
modules:
  http_2xx: # 这个名字是随便写的，但是需要在 prometheus.yml 配置文件中对应起来。
    prober: http # 进行探测的协议，可以是 http、tcp、dns、icmp
    timeout: 10s
  http_post_2xx:
    prober: http
    http:
      method: POST
  tcp_connect: # tcp检测，比如检测某个端口是否在线
    prober: tcp
  icmp: # icmp 检测，比如ping某个服务器
    prober: icmp
```
##### 3> 转为系统服务
```
vim /etc/systemd/system/blackbox_exporter.service
```
```
[Unit]
Description=blackbox_exporter
Documentation=https://prometheus.io/
After=network.target
[Service]
ExecStart=/usr/local/blackbox_exporter-0.22.0.linux-amd64/blackbox_exporter --config.file="/usr/local/blackbox_exporter-0.22.0.linux-amd64/blackbox.yml" --web.listen-address=":9098" --log.level=debug
WorkingDirectory=/usr/local/blackbox_exporter-0.22.0.linux-amd64/
Restart=on-failure
[Install]
WantedBy=multi-user.target
```
```
systemctl start blackbox_exporter
systemctl status blackbox_exporter
systemctl enable blackbox_exporter
```
##### 4> 集成进prometheus（master节点）
```
vim /usr/local/prometheus-2.37.0.linux-amd64/prometheus.yml
```
```
scrape_configs:
  - job_name: 'blackbox_http_2xx' # 配置get请求检测
    scrape_interval: 30s
    metrics_path: /probe
    params:
      module: [http_2xx]
    static_configs:
      - targets:         # 测试如下的请求是否可以访问的通
        - 152.32.170.211:10005
        - http://152.32.170.211:10005/hello/zhangsan
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 152.32.170.211:9098 # blackbox-exporter 服务所在的机器和端口
  - job_name: 'blackbox_http_post_2xx' # 配置post请求检测
    scrape_interval: 30s
    metrics_path: /probe
    params:
      module: [http_post_2xx]
    static_configs:
      - targets:              # 测试如下的post请求是否可以访问的通，该post请求不带参数
        - 152.32.170.211:10005
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 152.32.170.211:9098 # blackbox-exporter 服务所在的机器和端口
  - job_name: 'blackbox_http_ping' # 检测是否可以ping通某些机器
    scrape_interval: 30s
    metrics_path: /probe
    params:
      module: [icmp]
    static_configs:
      - targets:
        - 152.32.170.211
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 152.32.170.211:9098 # blackbox-exporter 服务所在的机器和端口
  - job_name: 'blackbox_tcp_connect' # 检测某些端口是否在线
    scrape_interval: 30s
    metrics_path: /probe
    params:
      module: [tcp_connect]
    static_configs:
      - targets:
        - 152.32.170.211:10006
        - 152.32.170.211:10005
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: 152.32.170.211:9098 # blackbox-exporter 服务所在的机器和端口

```
##### 5> 查看端口
```
netstat -npal|grep 9098
```
##### 6> 访问
```
http://152.32.170.211:9098/metrics
```
