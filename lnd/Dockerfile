# If you change this value, please change it in the following files as well:
# /.travis.yml
# /dev.Dockerfile
# /make/builder.Dockerfile
# /.github/workflows/main.yml
# /.github/workflows/release.yml
FROM golang:1.18.8-alpine as builder

# Force Go to use the cgo based DNS resolver. This is required to ensure DNS
# queries required to connect to linked containers succeed.
ENV GODEBUG netdns=cgo

# Pass a tag, branch or a commit using build-arg.  This allows a docker
# image to be built from a specified Git state.  The default image
# will use the Git tip of master by default.
ARG checkout="e45a59fc6c860f4de23d43ac3e0736d4b392c36a"
ARG repo="https://github.com/bottlepay/lnd"

# Install dependencies and build the binaries.
RUN apk add --no-cache --update alpine-sdk \
    git \
    make \
    gcc \
&&  git clone $repo /go/src/github.com/lightningnetwork/lnd \
&&  cd /go/src/github.com/lightningnetwork/lnd \
&&  git checkout $checkout \
&&  make install

# Start a new, final image.
FROM alpine as final

# Define a root volume for data persistence.
VOLUME /root/.lnd

# Add utilities for quality of life and SSL-related reasons. We also require
# curl and gpg for the signature verification script.
RUN apk --no-cache add \
    bash \
    jq \
    ca-certificates \
    gnupg \
    curl

# Copy the binaries from the builder image.
COPY --from=builder /go/bin/lncli /bin/
COPY --from=builder /go/bin/lnd /bin/
COPY --from=builder /go/src/github.com/lightningnetwork/lnd/scripts/verify-install.sh /
COPY --from=builder /go/src/github.com/lightningnetwork/lnd/scripts/keys/* /keys/

COPY run_lnd.sh /bin/

# Store the SHA256 hash of the binaries that were just produced for later
# verification.
RUN sha256sum /bin/lnd /bin/lncli > /shasums.txt \
  && cat /shasums.txt

# Expose lnd ports (p2p, rpc).
EXPOSE 9735 10009

ENTRYPOINT ["/bin/run_lnd.sh"]
