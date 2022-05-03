FROM golang:1.18-alpine as builder

# Force Go to use the cgo based DNS resolver. This is required to ensure DNS
# queries required to connect to linked containers succeed.
ENV GODEBUG netdns=cgo

# Install dependencies and build the binaries.
RUN apk add --no-cache --update alpine-sdk \
    git \
    make \
    gcc

WORKDIR /root

COPY go.mod .
COPY go.sum .

# Get dependancies - will be cached if we won't change mod/sum
RUN go mod download

COPY . /root

RUN go build ./cmd/testrunner

# Start a new, final image.
FROM alpine as final

# Add bash, jq and ca-certs, for quality of life and SSL-related reasons.
RUN apk --no-cache add \
    bash \
    jq alsa-utils \
    ca-certificates \
    tzdata

# Copy the binaries from the builder image.
COPY --from=builder /root/testrunner /bin/

ADD cmd/testrunner/entrypoint.sh /
RUN chmod +x entrypoint.sh

# Specify the start command and entrypoint as the lnd daemon.
ENTRYPOINT ["/entrypoint.sh"]
