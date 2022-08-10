### **部署Prometheus（master）**
##### 1> 获取prometheus二进制安装包
```
wget https://github.com/prometheus/prometheus/releases/download/v2.37.0/prometheus-2.37.0.linux-amd64.tar.gz
```
```
tar xzvf prometheus-2.37.0.linux-amd64.tar.gz \
&& mv prometheus-2.37.0.linux-amd64 /usr/local/
&& cd /usr/local/prometheus-2.37.0.linux-amd64
```
##### 2> 构建prometheus.yml配置文件
```
vim /usr/local/prometheus-2.37.0.linux-amd64/prometheus.yml
```
```
# 全局配置
global:
  # 设置抓取间隔,默认1分钟,配置是15秒
  scrape_interval: 15s
  # 估算规则的默认周期,默认1分钟,配置是15秒
  evaluation_interval: 15s 
  # 抓取超时时间,默认10秒
  scrape_timeout: 10s

# Alertmanager configuration
alerting:
  alertmanagers:
    - static_configs:
        - targets:
          # - alertmanager:9093

# 规则文件列表,使用  evaluation_interval 间隔去抓取
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# 抓取节点配置,使用 scrape_interval 间隔去抓取
scrape_configs:
  # prometheus默认的节点配置
  - job_name: "prometheus"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
      - targets: ["152.32.170.211:9090"]
```
##### 3> 运行
```
./prometheus --config.file=./prometheus.yml --web.listen-address=0.0.0.0:9090 --storage.tsdb.path=/usr/local/prometheus-2.37.0.linux-amd64/data
```
##### 4> 转为系统服务
```
vim /etc/systemd/system/prometheus.service
```
```
[Unit]
Description=prometheus
Documentation=https://prometheus.io/
After=network.target
[Service]
ExecStart=/usr/local/prometheus-2.37.0.linux-amd64/prometheus --config.file=/usr/local/prometheus-2.37.0.linux-amd64/prometheus.yml --storage.tsdb.path=/usr/local/prometheus-2.37.0.linux-amd64/data --web.listen-address=0.0.0.0:9090
WorkingDirectory=/usr/local/prometheus-2.37.0.linux-amd64/
Restart=on-failure
[Install]
WantedBy=multi-user.target
```
```
systemctl start prometheus
systemctl status prometheus
systemctl enable prometheus
```
##### 5> 查看端口
```
netstat -npal|grep 9090
```
##### 6> 访问
```
http://152.32.170.211:9090/
```
