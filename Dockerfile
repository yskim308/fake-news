# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Lambda runtime image
FROM public.ecr.aws/lambda/go:latest

# Copy Go binary
COPY --from=builder /src/main ${LAMBDA_TASK_ROOT}

# Copy HTML templates
COPY view/form.html ${LAMBDA_TASK_ROOT}/view/form.html
COPY view/main.html ${LAMBDA_TASK_ROOT}/view/main.html

# Lambda will invoke "main"
CMD [ "main" ]

