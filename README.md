# timeandhostname

This application returns hostname and current time of the environment that containers are running on top of against GET requests.

Application can function on any Kubernetes environment, but i will share details how to run on Minikube.

Docker container image can be built seperately. I have already built and pushed to the following:

`guresonur/wtitawithn` (you can directly pull the image as it is public in Dockerhub)

start a local Minikube cluster with Ingress addon (assuming there is already a profile for your personal usage, I will guide you to start a new profile):

`minikube start -p newprofile`

`minikube -p newprofile addons enable ingress` (More info regarding ingress with Minikube https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)

`#cd ../kubernetes`

`#kubectl apply -f .`

verify that service, ingress and deployment is there:

`#kubectl get service`

`#kubectl get ingress`

`#kubectl get deployment`

We want this service can serve external requests, so we will start a tunnel for Minikube

`#minikube -p newprofile tunnel`

LoadBalancer port is 8080

Open your web browser and hit the `http://127.0.0.1:8080`

<img width="344" alt="image" src="https://user-images.githubusercontent.com/27812281/193452982-4992fbc7-5b94-4c55-9285-731b0d9093e0.png">

--- 

If you want to build Go application to run on local without containerization:
`go build -o main`
`./main`
