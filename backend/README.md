# Go-Micro Project Template

## Install Micro
```bash
go install github.com/devexps/go-micro/cmd/micro/v2@latest
```

## Create a project
```bash
# Create a template project
micro new my_project

# Jump right in the `helloword` default service
cd my_project/helloworld

# Build and run the default service
make all
make run
```

## Create a service
```bash
# Move into your new project first
cd my_project

# Create a service
micro new my_service --service-only

# Jump right in `my_service` service
cd my_service

# Build and run `my_service` service
make all
make run
```

## Docker
```bash
# build
docker build \
  -t go-layout/helloworld:0.1.0 \
  -f .docker/Dockerfile \
  --build-arg SERVICE_RELATIVE_PATH=helloworld \
  --build-arg GRPC_PORT=9090 \
  --build-arg HTTP_PORT=8080 \
  .
	
# run
docker run --rm \
  -p 8080:8080 -p 9090:9090 \
  -v ${PWD}/helloworld/configs:/data/conf \
  go-layout/helloworld:0.1.0
```

You can also use makefile in each services to build and run docker. It looks like:
```bash
# Move into your new service first
cd my_project/my_service

# build
make docker

# run
make docker-run
```
