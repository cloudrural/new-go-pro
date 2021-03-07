# Base Image for Go program
FROM golang:alpine

# Environment varaibles
ENV WEBD_ROOT=/opt/webD

# Labels for WebD
LABEL WEBD_VERSION="$WEBD_VERSION"

# Build Args
ARG WEBD_VERSION

# Copy the application
COPY . $WEBD_ROOT

#Workdir
WORKDIR $WEBD_ROOT

# Port exposed
EXPOSE 8080

#Default command
CMD ["go","run","webD.go"]
