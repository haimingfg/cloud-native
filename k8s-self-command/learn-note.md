
# kubernetes 是什么？


# 组成

## 控制平面
控制平面节点是运行控制平面组件的机器， 包括 etcd （集群数据库） 和 API Server （命令行工具 kubectl 与之通信）。

## pods

## nodes

# 工具
kubectl 使用命令控制k8s集群

kind 在本地机器运行k8s

minikube 跟kind一样，构建单节点k8s

kubeadm 能够创建k8s集群


# demo 步骤
0、服务安装服务，[centos-7-install.md](centos-7-install.md)

1、使用kind按照本地集群
sudo kind create cluster

通过下面指令检测
// 获取节点
sudo kubectl get nodes

// 集群信息
sudo  kubectl cluster-info

ex：https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/cluster-interactive/

2、部署应用
The Deployment instructs Kubernetes how to create and update instances of your application
https://kubernetes.io/docs/tutorials/kubernetes-basics/deploy-app/deploy-intro/

sudo kubectl create deployment nginx-deploy --image=nginx:latest --port=80

sudo kubectl get deployments

export POD_NAME=$(sudo kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')

curl http://localhost:8001/api/v1/namespaces/default/pods/nginx-deploy-84657b87-xmtx9/

sudo kubectl status

// 第二个窗口打开进行测试
sudo kubectl proxy



export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')

curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/

3、通过 proxy 访问 pods

https://kubernetes.io/docs/tutorials/kubernetes-basics/explore/explore-intro/

// 获取pods信息，一个pods可以有多个容器
kubectl get pods

// 获取pods 更多信息，如images， ip
kubectl describe pods


// 开通pods外部访问的代理 
kubectl proxy


export POD_NAME=$(sudo kubectl get pods  -l app=nginx-deploy  -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')

curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/proxy/

curl http://localhost:8001/api/v1/namespaces/default/pods/proxy/$POD_NAME

$POD_NAME=nginx
// 查看pod启动
sudo kubectl logs $POD_NAME
// 查看 pod里面的环境变量
kubectl exec $POD_NAME -- env

// 进入pod里面的的bash模式 
kubectl exec -ti $POD_NAME -- bash


4、Using a Service to Expose Your App
https://kubernetes.io/docs/tutorials/kubernetes-basics/expose/expose-intro/

service 是一个抽象层能够定义pods被外部流量访问，负载均衡，服务发现pod

// 获取服务入口
kubectl get services

// 通过 expose 把应用向外开放
kubectl expose deployment/nginx-deploy --type="NodePort" --port 80

// -l标签, 作为筛选条件查询
kubectl get pods -l app=nginx-deploy
kubectl get services -l app=nginx-deploy 

// 打标签
export POD_NAME=$(sudo kubectl get pods -l app=nginx-deploy -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')
kubectl label pods $POD_NAME version=v1



// 查询情况
kubectl describe pods $POD_NAME
```
Name:         nginx-deploy-84657b87-xmtx9
Namespace:    default
Priority:     0
Node:         kind-control-plane/172.18.0.2
Start Time:   Sat, 30 Oct 2021 15:12:04 +0000
Labels:       app=nginx-deploy
              pod-template-hash=84657b87
              version=v1

```

// 通过version字段标签进行筛选
kubectl get pods -l version=v1


// 获取节点端口
export NODE_PORT=$(sudo kubectl get services/nginx-deploy -o go-template='{{(index .spec.ports 0).nodePort}}')
echo NODE_PORT=$NODE_PORT

// 访问
kubectl get nodes -o yaml 获取ip

curl 172.18.0.2:$NODE_PORT

// 删除服务
kubectl delete service -l app=nginx-deploy



5、autoscale

// 获取副本 
kubectl get rs

// 命令修改副本数
kubectl scale deployments/nginx-deploy --replicas=2

// 调整负载均衡
kubectl describe services nginx-deploy
export NODE_PORT=$(sudo kubectl get services/nginx-deploy -o go-template='{{(index .spec.ports 0).nodePort}}')
echo NODE_PORT=$NODE_PORT