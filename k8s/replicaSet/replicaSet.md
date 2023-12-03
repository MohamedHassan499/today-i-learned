# Introduction

Let's say that we have an installed pod in our node which runs our Node.js API, what If that pod receives a lot of requests and it can't handle all of them? We can't scale it up because it's a single pod, so we need to create more pods to handle the requests. But how can we do that? We can do that with a ReplicaSet.

Also, If one of our pods fails, we need to create another one to replace it, and that's another thing that a ReplicaSet can do for us.

# ReplicaSet/ReplicationController

A ReplicaSet is a Kubernetes object that helps us to create and maintain a set of pods that are identical. It's a controller that ensures that a specified number of pod replicas are running at any given time.

> Note:
> Both ReplicaSet and ReplicationController are the same, but the ReplicaSet is the new version of the ReplicationController.

## Write YAML file

Let's create a ReplicaSet that will create 3 replicas of our Node.js API.

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: nodejs-api
spec:
    replicas: 3
    selector:
        matchLabels:
        app: nodejs-api
    template:
        metadata:
        labels:
            app: nodejs-api
        spec:
        containers:
            - name: nodejs-api
            image: nodejs-api
            ports:
                - containerPort: 3000
```

> Template is the pod template that will be used to create the pods.

and to create it we need to run the following command:

```bash
kubectl apply -f nodejs-api-replicaset.yaml
```

## How about scaling replicas?

We can scale our replicas with the following command:

```bash
kubectl scale replicaset nodejs-api --replicas=5
```

But notice that it doesn't modify the number of replicas in our YAML file, it just scales it up or down.

## Delete, list or update a ReplicaSet

To delete a ReplicaSet we can run the following command:

```bash
kubectl delete replicaset nodejs-api
```

To update a ReplicaSet we can run the following command:

```bash
kubectl replace -f nodejs-api-replicaset.yaml
```

To list all ReplicaSets we can run the following command:

```bash
kubectl get replicaset
```

> Note:
> When we delete pod from a ReplicaSet, the ReplicaSet will create another pod to replace it.
> When we create pod with the same name as a pod in a ReplicaSet, the ReplicaSet will delete the old pod and create a new one with the same name.
