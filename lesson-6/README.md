# Improving Development Workflow with Skaffold

```sh
skaffold init
```

View the created `skaffold.yaml` file.

Before running skaffold, verify which cluster `kubectl` is configured with.

```sh
kubectl config current-context
```

Start skaffold in development mode.

```sh
skaffold dev
```

Verify Deployment created Pods.

```sh
kubectl get pods
```

Update `src/main.go` to reflect a new version number.

```go
const version = "0.2"`
```

View logs to verify the new version has been deployed.

How do we verify the change by accessing the API?

Stop `skaffold` and rerun with the `--port-forward` flag.

```sh
skaffold dev --port-forward
```

Notice the port that gets reported in the logs.

```sh
# NOTE: Your port may differ.
curl 127.0.0.1:4503
```

What happens when we update the Kubernetes manifests?
Update the `k8s/deployment.yaml` manifest to specify 5 replicas.

```yaml
spec:
  replicas: 5
```

Note that in the skaffold logs, we are not rebuilding our image.
Because our Kubernetes manifests live outside of our configured Docker context, skaffold knows that it does not need to rebuild our images.

