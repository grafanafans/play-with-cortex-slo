# play-with-cortex-slo

An slo service example with cortex cluster and tools.


## Architecture

<img width="721" alt="image" src="https://user-images.githubusercontent.com/1459834/233758625-701aa171-19d0-41cd-8ca5-706aae21ee89.png">

- Deploy three Cortex instances with ruler and alertmanager.
- Using Cortex SLO service to auto generate prometheus rules and store them to Cortex.
- Using Cortex Tools to management tenant SLOs.
- Using `MyService` to simulates error request with any ratio.
- Using Grafana to view MyService SLO detail Dashboard and related rules.
- Using Echo receiver to print alertmanager notify.


## How to run

- start apps

```
git clone https://github.com/grafanafans/play-with-cortex-slo.git
cd play-with-cortex-slo
docker-compose up -d
```

- download cortextool

```
wget https://github.com/grafanafans/play-with-cortex-slo/releases/download/v0.1.0/cortextool-amd64-darwin.tar.gz 
tar xvf cortextool-amd64-darwin.tar.gz 
chmod +x cortextool-amd64-darwin  && mv cortextool-amd64-darwin /usr/local/bin/cortextool
```

- load slos

```
export CORTEX_TENANT_ID=demo
export CORTEX_ADDRESS=http://localhost:6666

cortextool slos load-windows  config/slos/windows/*.yaml
cortextool slos load config/slos/myservice.yml --windows google-30d
```

- set `myservice` error rate


```
curl http://localhost:8081/errrate?value=0.05
```

- load sloth grafana dashboard

Go `http://localhost:3000/dashboards` page and new dashboard with import by ids(14348 and 14643).

slo high level dashboard:

![sloth-high.png](/images/sloth-high.png)


slo detail level dashboard:

![sloth-detail.png](/images/sloth-detail.png)