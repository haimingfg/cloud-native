- 启动指令

###  kubeadm 初始化控制平面节点 

sudo kubeadm init --image-repository registry.aliyuncs.com/google_containers  
// 指定ip
--apiserver-advertise-address=192.168.34.2 


### kubelet: The Kubernetes Node Agent




// 获取节点
sudo kubectl get nodes

// 集群信息
sudo  kubectl cluster-info

// Check that kubectl is configured to talk to your cluster

kubectl version -o json

// 获取pods信息，一个pods可以有多个容器
kubectl get pods

// 获取pods 更多信息，如images， ip
kubectl describe pods


// 开通pods外部访问的代理 
kubectl proxy

### kind
创建集群
sudo kind create cluster