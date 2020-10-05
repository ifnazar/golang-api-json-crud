FROM golang

# Create the workspace directory
RUN mkdir /app

# Define the default workspace
WORKDIR /app

# Copy the local package files to the container's workspace
ADD . /app/

# Download all the dependencies
RUN go get -d -v .

# Build the app
RUN go build -o main .

CMD ["./main"]

# Document that the service listens on port 8080.
EXPOSE 3000