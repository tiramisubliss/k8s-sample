username=$(echo -n "admin" | base64)
password=$(echo -n "a62fjbd37942dcs" | base64)

kubectl create -f secret.yaml

kubectl get secrets

kubectl create -f secret-env.yaml

kubectl exec -it secret-env-pod env | grep SECRET_

kubectl get pods

kubectl create -f secret-pod.yaml

kubectl exec -it secret-vol-pod ls /etc/secret-volume

kubectl exec -it secret-vol-pod cat /etc/secret-volume/username

kubectl exec -it secret-vol-pod cat /etc/secret-volume/password
