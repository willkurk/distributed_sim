set -e

docker build -t distgame/renderer \
  -f docker/Dockerfile.renderer .
docker build -t grpcweb/envoy \
  -f docker/Dockerfile.envoy .
docker build -t distgame/sampleclient \
  -f docker/Dockerfile.sampleclient .
