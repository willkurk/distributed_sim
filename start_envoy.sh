set -e

docker build -t grpcweb/envoy \
  -f docker/Dockerfile.envoy .
docker run -d -p 8080:8080 --network="host" grpcweb/envoy
