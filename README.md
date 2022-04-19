# Sample Continuous Integration Application for Kubernetes written on Go

This application demonstrates the lightweight result of Go and consumes a PostgreSQL StatefulSet.

It includes NetworkPolicies and IngressController which defines the baseuri path on /app-golang.

This is the CI part of the application, Jenkins will build any further change to the code and will push here the change on the Deployment. 

After that, ArgoCD will deploy the objects on the manifest.

![Jenkins Pipeline](/images/Jenkins.png)


![Docker Hub image](/images/DockerHub.png)
