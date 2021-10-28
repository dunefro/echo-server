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