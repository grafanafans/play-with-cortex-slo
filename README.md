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


## How to start

```
git clone https://github.com/grafanafans/play-with-cortex-slo.git
cd play-with-cortex-slo
docker-compose up -d
```

- download cortextools

```
// comming soon
```
