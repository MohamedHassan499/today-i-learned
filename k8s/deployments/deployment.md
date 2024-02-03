## Deployment

### What is deployment

Deployment is a higher level object that manages ReplicaSets and provides declarative updates to Pods along with a lot of other useful features.

### Deployment strategies

- Recreate
- RollingUpdate

1. Recreate:

- All old pods are killed at once and then all new pods are created at once.
- It is not recommended for production. (as it will cause downtime)

2. RollingUpdate:

- It gradually replaces old pods with new ones.
- It is the default strategy.
- It is recommended for production.

> [!NOTE]
>
> - Deployment is the recommended way to manage the creation and scaling of Pods.
> - When you upgrade your application, you can deploy a new version of your image by updating the PodTemplateSpec of your Deployment.
> - The new PodTemplateSpec will then be used to create new ReplicaSets which will in turn create the Pods to run your application.

- You can use the rollout command to check the status of a rollout:
  ```bash
  kubectl rollout status deployment/nginx-deployment
  ```
- You can also undo a rollout or rollout to a previous revision:
  ```bash
  kubectl rollout undo deployment/nginx-deployment
  ```
- You can also pause and resume a rollout:
  ```bash
  kubectl rollout pause deployment/nginx-deployment
  kubectl rollout resume deployment/nginx-deployment
  ```
- You can also scale a Deployment by using the scale subcommand:
  ```bash
  kubectl scale deployment nginx-deployment --replicas=10
  ```
- You can also edit a Deployment by using the edit subcommand:
  ```bash
  kubectl edit deployment nginx-deployment
  ```
- You can also delete a Deployment by using the delete subcommand:
  ```bash
  kubectl delete deployment nginx-deployment
  ```

### Create a deployment

```bash
kubectl create deployment nginx --image=nginx
```

### Get deployments

```bash
kubectl get deployments
```

### Get deployment details

```bash
kubectl describe deployment nginx
```

### Get deployment yaml

```bash
kubectl get deployment nginx -o yaml
```

### RollUpdate deployment

```bash
kubectl set image deployment/nginx nginx=nginx:1.9.1 --record
```
