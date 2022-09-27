FROM gcr.io/distroless/base:nonroot
WORKDIR /
COPY cj-killproxy /
USER nonroot:nonroot
EXPOSE 15020
ENTRYPOINT [ "/cj-killproxy" ]