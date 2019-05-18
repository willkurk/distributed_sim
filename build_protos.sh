protoc -I protos/ protos/world.proto --go_out=plugins=grpc:protos
protoc -I protos/ protos/position.proto --go_out=plugins=grpc:protos
protoc -I protos/ protos/world_render.proto --go_out=plugins=grpc:protos

protoc -I=protos/ world.proto \
--js_out=import_style=commonjs:protos/js \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:protos/js

protoc -I=protos/ position.proto \
--js_out=import_style=commonjs:protos/js \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:protos/js

protoc -I=protos/ world_render.proto \
--js_out=import_style=commonjs:protos/js \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:protos/js
