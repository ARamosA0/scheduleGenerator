FROM rust:1.82

WORKDIR /app
COPY . .

# Instala cargo-watch para hot-reload (opcional)
#RUN cargo install cargo-watch

EXPOSE 8088

# Default: ejecuta el binario en modo debug
#CMD ["cargo", "watch", "-x", "run"]
CMD ["cargo", "run"]
