# echo-server
This is an echo server based on GO

# Build
```
export DOCKER_BUILDKIT=1
docker image build -t echo-server:v0.1.0 .
```

# Deploy

```
helm install echo echo-server
```

# To test
```
SVC_IP=$(kubectl get svc -l app.kubernetes.io/name=echo-server -ojsonpath='{.items[0].spec.clusterIP}')
SVC_PORT=$(kubectl get svc -l app.kubernetes.io/name=echo-server -ojsonpath='{.items[0].spec.ports[0].port}')

# 'apple is the text that will be echoed'
curl $SVC_IP:$SVC_PORT/apple
```
If we pass multiple endpoints after a single `/` then `/` will be replaces by a space
```
$ curl $SVC_IP:$SVC_PORT/apple/banana/papaya
apple banana papaya
```