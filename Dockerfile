# Build the manager binary
FROM golang:1.21 AS backend-builder
COPY backend/ /build/backend
COPY build.sh /build/
WORKDIR /build
RUN sh -x ./build.sh --backend && ls /build/bin/ && sleep 10

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM gcr.io/distroless/static:nonroot
FROM node:21 AS frontend-builder
COPY frontend/ /build/frontend
COPY build.sh /build/
WORKDIR /build/frontend
RUN npm install
RUN npm run build

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
#FROM gcr.io/distroless/static:nonroot
FROM debian:latest
WORKDIR /src
COPY --from=backend-builder /build/bin/cfs-gui .
COPY --from=backend-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=frontend-builder /build/frontend/dist ./dist/
EXPOSE 6007
ENTRYPOINT ["/src/cfs-gui","-c","/src/config/"]