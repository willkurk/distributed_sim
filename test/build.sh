set -e

mkdir -p vendor/github.com/willkurk/distgame/protos
cp -r ../protos vendor/github.com/willkurk/distgame/
go build main.go
