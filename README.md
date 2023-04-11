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