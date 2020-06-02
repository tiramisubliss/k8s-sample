# Canary Deployment Under Kubernetes

Create the first deployment, `deployment-red` into the Kubernetes cluster.

`kubectl apply -f deployment-red.yaml`

Create the second deployment, `deployment-green` into the Kubernetes cluster.

`kubectl apply -f deployment-green.yaml`

Let's retrieve a list of pods according to the label, `type=example_code`.

`kubectl get pods -l app=example_code`
`kubectl get pods -l app=example_code  -l color=red`
`kubectl get pods -l app=example_code  -l color=green`

Now it's time to bind the pods to some services. First, let's take a look at the contents of the first service as defined
in the yaml file, `service_01.yaml`.

`kubectl apply -f service_01.yaml`
`kubectl apply -f service_02.yaml`
`kubectl apply -f service_03.yaml`

`kubectl get services  --all-namespaces --field-selector metadata.name=echocolor-all`

`kubectl cluster-info`

`for i in {1..10}; do curl 172.17.0.40:30671; done`

`kubectl get services  --all-namespaces --field-selector metadata.name=echocolor-red`

`for i in {1..10}; do curl 172.17.0.40:31298; done`

The way we do this is to delete the `red` deployment like so:

`kubectl delete -f deployment-red.yaml`

`kubectl get services  --all-namespaces --field-selector metadata.name=echocolor-all`

`for i in {1..10}; do curl 172.17.0.40:30671; done`