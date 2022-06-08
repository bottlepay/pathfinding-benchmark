FROM node:16 as build-web-admin

WORKDIR /build

RUN git clone https://github.com/L2-Technology/sensei.git . && git checkout 97a6f7af94571df51b5b9c9df3ef9417554cafb4

WORKDIR /build/web-admin
RUN npm install
RUN npm run build

FROM rust:1.56 as build-sensei

WORKDIR /build

# copy your source tree
RUN git clone https://github.com/L2-Technology/sensei.git . && git checkout 97a6f7af94571df51b5b9c9df3ef9417554cafb4

COPY --from=build-web-admin /build/web-admin/build/ /build/web-admin/build/

RUN rustup component add rustfmt

RUN cargo build --verbose --release

# our final base
FROM rust:1.56

# copy the build artifact from the build stage
COPY --from=build-sensei /build/target/release/* .

# set the startup command to run your binary
CMD ["./senseid"]
