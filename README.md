# Developer Productivity

A KubeAcademy course for developers new to Kubernetes.

## Lesson 1: Setting Up Your Workstation

Install Docker Engine on Linux:
https://docs.docker.com/install/linux/docker-ce/ubuntu/

Install Docker Desktop on Mac:
https://docs.docker.com/docker-for-mac/install/

Install KinD:
https://github.com/kubernetes-sigs/kind#installation-and-usage

Install kubectl:
https://kubernetes.io/docs/tasks/tools/install-kubectl/

Install Kustomize:
https://github.com/kubernetes-sigs/kustomize/blob/master/docs/INSTALL.md

Install Helm:
https://helm.sh/docs/intro/install/

Install Skaffold
https://skaffold.dev/docs/install/

## Lesson 2: Building Container Images

Optional: Build developer-productivity app.  Note: this step requires having the Go
programming language installed:

    go build

Optional: Run the app to ensure it works:

    ./developer-productivity

Visit http://localhost:8000/ in your browser.

Build docker image:

    docker build -t <image repo> .

Tag it with a version tag:

    docker tag <image id> <image repo>:0.1

List your images:

    docker images -f reference=<image repo>

Log in to registry:

    docker login

Push image to registry:

    docker push <image repo>:0.1

## Lesson 3: Running Kubernetes Loccally

Create a KinD cluster:

    kind create cluster --name kubeacademy --config kind-config.yaml

Examine the cluster nodes:

    kubectl get nodes

Examine pods in the cluster:

    kubectl get pods -A

Load image onto KinD cluster nodes:

    kind load docker-image --name kubeacademy <image repo>:0.1

Create the pod:

    kubectl apply -f pod.yaml

Follow the pod logs:

    kubectl logs developer-productivity-pod -f

In other terminal forward a port to the pod:

    kubectl port-forward pod/developer-productivity-pod 8080:8000

Use your browser to make a request to your pod by visiting
http://localhost:8080/

Clean up:

    kubectl delete -f pod.yaml

## Lesson 4: Deploying Your Application

* at least make mention of configmaps and secrets, demo if time

Create deployment:

    kubectl apply -f deployment.yaml

Create service:

    kubectl apply -f service.yaml

Port forward local 8000 to 80 on service:

    kubectl port-forward service/developer-productivity-svc 8000:80

Again, use your browser to make a request to your pod by visiting
http://localhost:8080/

Clean up:

    kubectl delete -f service.yaml -f deployment.yaml

Generate combined manifest with Kustomize:

    kustomize build base

Generate manifest for production:

    kustomize build overlays/production

Create resources with overlays:

    kustomize build overlays/production | kubectl apply -f -

Clean up:

    kustomize build overlays/production | kubectl delete -f -

## Lesson 5: Packaging Your Application

