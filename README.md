# gin-vue-template
 - backend: gin
 - frontend: vue

## Quick Start
### Dependency
1. install golang
2. install nodejs and npm
3. `cd frontend && npm install`
4. install nodemon via npm
5. install mysql
6. install redis

### DEV
1. `bash ./go_generate.sh`
2. `bash ./devrun.sh`
3. `cd frontend && npm run serve`

## File Tree
- 📁 gin-vue-template
  - 📁 cmd
    - 📄 main.go // program entry
  - 📁 conf
    - 📄 app.yaml // some configuration items of course
  - 📁 config // read configuration item from conf/app.ini
    - config.go
  - 📁 pkg
    - 📁 constants // some constants
      - 📄 ...
    - 📁 e // errcode
        - 📄 e.go
    - 📁 infra
      - 📁 redis // redis
      - 📁 logger // init logrus and impl some function for log
    - 📁 logic // do some logic for your api
    - 📁 models // connecting to db
      - 📄 ... // some models and DAO functions
    - 📁 service // http service provider for wire gen
    - 📁 utils // maybe put some utils here
    - 📁 wire // wire
  - 📁 routers
    - 📁 api // yep, http api
    - 📁 middleware // middleware for gin
      - 📄 init_middleware.go // initialize, generate logid for example
      - 📄 auth_middleware.go // authorize, abort if unauthorized
    - 📄 router.go // router definition
  - 📁 frontend // vue proj


## Reference
 - [gin-vue](https://github.com/Bingjian-Zhu/gin-vue)
