# What is it?
HPA Demo with kubernetes.
Further reading https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/

# Prerequisites
 
* Kubernetes
* kubectl

# How to run

```sh
kubectl apply -f deployment.yaml -f hpa.yaml
```

# How to test HPA

```sh
kubectl run -i --tty load-generator --rm --image=busybox --restart=Never -- /bin/sh -c "while sleep 0.0000001; do wget -q -O- http://hpa-simple-app; done"
```

# How to clean up

```sh
kubectl delete service hpa-simple-app
kubectl delete deployment hpa-simple-app
kubectl delete horizontalpodautoscaler hpa-simple-app
```