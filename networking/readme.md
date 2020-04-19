$ export CLUSTER_IP=$(kubectl get services/webapp1-clusterip-svc -o go-template='{{(index .spec.clusterIP)}}')
$ echo CLUSTER_IP=$CLUSTER_IP
$ curl $CLUSTER_IP:80

$ export CLUSTER_IP=$(kubectl get services/webapp1-clusterip-targetport-svc -o go-template='{{(index .spec.clusterIP)}}')
$ echo CLUSTER_IP=$CLUSTER_IP $ curl $CLUSTER_IP:8080

