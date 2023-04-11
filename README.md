--------------------------------------------------------------
# Exercise 1 Solution
--------------------------------------------------------------

## Password Generator API
This is a prototype of a REST API that generates secure passwords. It takes four input parameters in the request:

* `min_length`: the minimum length of the generated password
* `special_chars`: the number of special characters to include in the password
* `numbers`: the number of numbers to include in the password
* `num_passwords`: the number of passwords to generate

The program starts an HTTP server that listens for incoming requests on port 8000. There are three endpoints available: /genpass for generating passwords, /healthz for checking the health of the application, and /readyz for checking the readiness of the application.

The /genpass endpoint accepts a POST request with the user's input as a json payload. It generates secure passwords based on the user's input and returns them in an array.

The /healthz endpoint accepts a GET request and returns a 200 OK status code to indicate that the application is healthzy.

The /readyz endpoint also accepts a GET request and returns a 200 OK status code only if the application is ready. It uses an atomic variable to track the readiness status and a middleware function to check the status before each request.

### Prerequisites
Before running this application, you must have the following installed:

- Go (to build localy)
- Docker (to build and test Localy)
- kubectl, helm (optional, for k8s deployment)

### Set up

To set up the password generator API, follow these steps:

1. Clone the repository:
```
git clone https://github.com/abdo-farag/password-generator.git
```
2. Build the Docker image:
```
cd password-generator && make docker-build
```
3. Run the Docker container: 
```
make docker-run
```
4. Run application test:
```
make test
```
5. Uninstall test docker setup:
```
make clean-docker
```

### deploy the appy to kubernetes cluster using Makefile
1. Build docker image, push it and deploy the application to kubernetes cluster using `kubectl apply`.
```
make all
```
2- Start port forwarding using:
```
kubectl port-forward -n password-generator deployments/password-generator  8000:8000
```

3- Test the api functionalty
```
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
           "min_length": 12,
           "special_chars": 2,
           "numbers": 2,
           "num_passwords": 3
         }' \
     http://localhost:8000/genpass
```

4- Clean k8s deployment
```
make clean-k8s
```
### Response example
The response will be a JSON object with an array of strings representing the generated passwords. The example response for the above request would be:
```
{
  "passwords": [
    "E)9xxP9)OljR",
    "fq8dl|2vxd.k",
    "xs;UpL:W24qV"
  ]
}
```
### Deployment using helm Chart
To deploy this application to Kubernetes, you can use the included Helm chart. Follow the steps below to deploy the application:
```
helm upgrade --install password-generator ./Charts/password-generator/ --namespace=password-generator --create-namespace
```

ingress and hpa autoscalling can be enabled using chart options:
```
--set ingress.enabled=true --set autoscaling.enabled=true
```

--------------------------------------------------------------
# Exercise 2 Solution
--------------------------------------------------------------

## Possible improvements to the Terraform code:
- Modularize the code to make it more manageable and reusable.
- Use variables and input parameters to make the code more dynamic and adaptable to different environments.
- Use remote state management to improve the collaboration between team members and avoid conflicts.
- Add more security controls, such as IAM policies, security groups, and encryption.
- Make sure the S3 bucket and DynamoDB table are only accessible by authorized users/roles. useing AWS IAM policies to restrict access to these ressources.
- Using official aws modules to create vpc, network, route53, s3, and create iam policy for kops to access this resources.
- Consider implementing backup and restore procedurues for your state storage resources. You can use AWS Backup service to automate the backup and restore process.

- Using the kops CLI to generate Terraform files with the --target=terraform optiion will result in a large, single Terrafrom configuration file. While this file may be unpractical, it is an Infrastructure as Code (IaC) file in the .tf format that can be easily maintained, updated using `kops update`, version-controlled and automate deployment using CI/CD pipeline



## Design monitoring, logging and alerting architecture for these environment

To design a monitoring, logging, and alerting architecture for these environments, I would consider the following components:

* Metrics Collection: I would use a metrics collection tool such as Prometheus to collect metrics from various sources such as the Kubernetes API server, cluster nodes, and application components running in the cluster. These metrics would be used to monitor the performance and health of the cluster and its workloads.

* Logging: For logging, I would use a centralized logging solution such as the Elastic Stack (Elasticsearch, Logstash, and Kibana) or Fluentd to collect and aggregate logs from various sources including Kubernetes API server, cluster nodes, and application compoents. The centralized logging solution should support log agregation, searching, and visualization.

* Alerting: I would use an alerting tool such as Prometheus Alertmanager or Grafana Loki to monitor the collected metrics and logs and send alerts when predefined thresholds or anomalies are detected. Alerting should be based on rules and connditions pre-defined, and should include escalation policyes and notification channnels such as email or Slack. Use incident management and alerting tools, such as PagerDuty, to manage and esclate incidents and alerts based on severity, priority, and impact.

* Dashboarding and Visualization: For dashboarding and visualiziation, I would use tools such as Grafana or Kibana to create custom dashboards and visualizations based on the collected metrics and logs. These dashboards would provide visibility into the overall health and performance of the cluster and its applications.

* Tracing: To troubleshoot and analyze issues with the applications, I would use a distributed tracing solution such as Jaeger or SigNoz to trace requests as they traverse the various components of the application.

## suggestion to improve previous solutions scalability

To improve the scalabillity of the deployment, monitoring, and loging architecture for dozens of installations, we can consider the following approaches:

Scalable monitoring infrastructure.
We can decouple the monitoring infrastrucuture from the main infrastructure and move it to a separate scalable cluster. This technic offers greater flxibility in scaling each infrastructure indepedently and helps to isolate the monitoring tolls and data from the main infrastructure, reducing the risk of resource contention and increasing security. We can also use dedicated monitoring clusters to manage monitoring infrastructure and its resources.

Migrating to a managed Kubernetes service (EKS)
We can consider migrating from Kops to a fully managed Kuberntes service like Amazon (EKS). EKS can help to simplify the management of the Kubernetes control plane, reducing the operational burden of managing a cluster at scale. EKS provides native integrations with other AWS services.
