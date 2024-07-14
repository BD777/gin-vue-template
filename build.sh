bash go_generate.sh

mkdir -p ./build
mkdir -p ./build/conf
mkdir -p ./build/frontend
touch ./build/conf/app.yaml

cd frontend && npm run build && cp -r ./dist ../build/frontend && cd ..
go build -o ./build/gin-vue-template ./cmd