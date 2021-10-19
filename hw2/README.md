# 指令
```
docker build . -t httpserver_example:0.2

docker exec -it {containerid} sh


PID=$(docker inspect --format "{{ .State.Pid }}" {containerid})


nsenter --target $PID --mount --uts --ipc --net --pid

```

仓库连接(https://hub.docker.com/repository/docker/haimingli/httpserver_example)
