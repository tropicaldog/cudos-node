FROM debian

RUN apt -y update
RUN apt -y install clang build-essential wget

# GET FROM https://github.com/rust-lang/docker-rust-nightly
ENV RUSTUP_HOME=/usr/local/rustup \
    CARGO_HOME=/usr/local/cargo \
    PATH=/usr/local/cargo/bin:$PATH


RUN url="https://static.rust-lang.org/rustup/dist/x86_64-unknown-linux-gnu/rustup-init"; \
    wget "$url"; \
    chmod +x rustup-init; \
    ./rustup-init -y --no-modify-path --default-toolchain 1.51.0; \
    rm rustup-init; \
    chmod -R a+w $RUSTUP_HOME $CARGO_HOME; \
    rustup --version; \
    cargo --version; \
    rustc --version;

# allow non-root user to download more deps later
RUN chmod -R 777 /usr/local/cargo

## COPY BUILD SCRIPTS

WORKDIR /code

# COPY /project-wasmvm /code

# COPY guest/*.sh /opt/
# RUN chmod +x /opt/*.sh

RUN mkdir /.cargo
RUN chmod +rx /.cargo
COPY /project-wasmvm/builders/guest/cargo-config /.cargo/config

CMD ["/bin/bash", "-c", "chmod +x ./builders/guest/build_linux.sh && ./builders/guest/build_linux.sh"] 