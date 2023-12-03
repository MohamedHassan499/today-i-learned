# Replication controller:

Replication controller is a resource that ensures a specified number of pods are running at any one time. It also allows you to easily scale up or down by adding or removing replicas.

Imagine that you have a pod running a web server. You want to make sure that there are always 3 replicas of that pod running. If one of the pods crashes, the replication controller will automatically start a new one to replace it. If you want to scale up to 5 replicas, you can just tell the replication controller to do so.

It will autoscaler by spawning new pods if the current pods are using too much CPU or memory. Maybe It will spawn that pods on another nodes if the current nodes are using too much CPU or memory.

## Replication controller vs ReplicaSet:

Replication controller is the older version of ReplicaSet. ReplicaSet is the newer version of Replication controller. ReplicaSet is more powerful than Replication controller. ReplicaSet is the next generation of Replication controller.

## Replication controller yaml file:

```yaml
apiVersion: v1
kind: ReplicationController
metadata:
  name: nginx-rc
spec:
    replicas: 3
    selector:
        app: nginx
    template:
        metadata:
        name: nginx-pod
        labels:
            app: nginx
        spec:
        containers:
        - name: nginx-container
            image: nginx
```

> The template here is the pods that you'll spwan with the replication controller. The selector is the label that you'll use to select the pods that you want to manage with the replication controller.

## Create a replication controller:

```bash
kubectl create -f replication-controller.yaml
```

## Get replication controller:

```bash
kubectl get rc
```

## Get pods:

```bash
kubectl get pods
```
