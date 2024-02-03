# K8s networking

## 1. Pod networking

### 1.1. Pod IP addresses

- Each Pod is assigned a unique IP address.
- Every container in a Pod shares the network namespace, including the IP address and network ports (Which means they can share the same localhost interface).
- Containers inside a Pod can communicate with one another using localhost.
- Containers outside the Pod can also communicate with a Pod if they know its IP address.

### 1.2. Pod-to-Pod communications

- Pods can communicate with other Pods in the same node.
- Pods can communicate with other Pods in different nodes.
- Pods can communicate with other Pods in different clusters.

### 1.3. Pod-to-Service communications

- Pods can communicate with Services in the same cluster.
- Pods can communicate with Services in different clusters.
