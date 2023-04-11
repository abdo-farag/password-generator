# Define variables
GO_FLAGS := CGO_ENABLED=0 GOOS=linux GO111MODULE=on
IMAGE_NAME := abdofarag/password-generator
IMAGE_TAG := latest
DEPLOY_FILE := password-generator.yml

# Run all targets
all: docker-build docker-push k8s-deploy

# Build the Go binary
build:
	@$(GO_FLAGS) go build -o password_generator ./cmd

# Build the Docker image
docker-build:
	@docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .

# Push the Docker image to Docker Hub
docker-push:
	@docker push $(IMAGE_NAME):$(IMAGE_TAG)

# Start the application conatiner
docker-run:
	@docker run -tid --name genpass -p 8000:8000 $(IMAGE_NAME):$(IMAGE_TAG)

# test application
test:
	@curl -X POST -H "Content-Type: application/json" -d '{ "min_length": 12, "special_chars": 2, "numbers": 2, "num_passwords": 3 }' http://localhost:8000/genpass

# Apply the Kubernetes deployment
k8s-deploy:
	@kubectl apply -f $(DEPLOY_FILE)

# Forward pod port 8000 to localhost
k8s-port-forward: 
	@kubectl port-forward -n password-generator deployments/password-generator 8000:8000

# Clean up
clean-docker:
	@docker rm -f genpass
	

# clean k8s
clean-k8s:
	@kubectl delete -f $(DEPLOY_FILE) --ignore-not-found=true

# Clean up
clean-all: clean-k8s clean-docker
	@rm -rf password_generator


.PHONY: build docker-build docker-push docker-run test k8s-deploy k8s-port-forward clean-docker clean-k8s clean-all 
