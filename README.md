# distributed_sim
A distributed golang game server. It currently implements a simple population simulation. Has a browser based client with an extremely simple implementation.

This is because a backend services handles the rendering work.

# BUILD

```
./build_protos.sh
./sampleclienta/update_protos.sh
./build_docker.sh

docker-compose -f compose/system-test.yml up -d
```
