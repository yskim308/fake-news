# --- Stage 1: The Builder ---
# This stage builds the Go binary.
# We use a specific Go version. Update "1.21" if you use a different one.
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /src

# Copy the go.mod and go.sum files first
# This leverages Docker's layer caching.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your source code (the "root")
COPY . .

# Build the binary for Lambda
# CGO_ENABLED=0 creates a static binary
# GOOS=linux is required by the Lambda runtime
# -o /main outputs the compiled binary to a file named "main" at the root
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .


# --- Stage 2: The Final Image ---
# This is the actual image you will deploy to Lambda.
# It uses the official Amazon Lambda base image for Go.
FROM public.ecr.aws/lambda/go:1

# The LAMBDA_TASK_ROOT is /var/task. This is where Lambda runs your code.
# We copy the built binary from the "builder" stage into this directory.
COPY --from=builder /main /var/task/main

# Copy your HTML templates into the image, matching the path in your Go code
# Your code looks for "./view/form.html", so we create that directory structure.
# We must create the 'view' directory first.
RUN mkdir -p /var/task/view
COPY view/form.html /var/task/view/form.html
COPY view/main.html /var/task/view/main.html

# Set the command to run when the container starts.
# This tells Lambda to execute your compiled binary named "main".
CMD [ "main" ]
