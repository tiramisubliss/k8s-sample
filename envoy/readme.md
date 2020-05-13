```
docker run --name=proxy -d \
  -p 80:10000 \
  -v $(pwd)/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
  envoyproxy/envoy:latest
```


```
docker run --name=proxy-with-admin -d \
    -p 9901:9901 \
    -p 10000:10000 \
    -v $(pwd)/envoy/envoy.yaml:/etc/envoy/envoy.yaml \
    envoyproxy/envoy:latest
```

```
docker run --name proxy1 -p 80:8080 --user 1000:1000 -v /root/envoy.yaml:/etc/envoy/envoy.yaml envoyproxy/envoy
```
