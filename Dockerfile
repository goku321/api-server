FROM golang
ARG ACCESS_TOKEN
ENV GIT_TERMINAL_PROMPT=1
ENV GOPRIVATE="github.com/goku321"
RUN git config --global url."https://${ACCESS_TOKEN}:@github.com/".insteadOf "https://github.com/"
# Set the Current Working Directory inside the container
WORKDIR /app
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY app/main.go .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o api-server .

# Command to run when starting the container
CMD [ "./api-server" ]