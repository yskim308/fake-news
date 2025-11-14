FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main .


# final image
FROM public.ecr.aws/lambda/provided:al2

# Copy the binary from the "builder" stage
COPY --from=builder /main /var/task/main

# Copy HTML templates
RUN mkdir -p /var/task/view
COPY view/form.html /var/task/view/form.html
COPY view/main.html /var/task/view/main.html

CMD [ "main" ]
