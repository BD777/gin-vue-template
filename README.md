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
1. `bash ./devrun.sh`
2. `cd fronend && npm run serve`

## File Tree
- 📁 gin-vue-template
  - 📁 conf
    - 📄 app.ini // some configuration items of course
  - 📁 constants // some constants
    - 📄 ...
  - 📁 middleware // middleware for gin
    - 📄 auth_middleware.go // authorize, abort if unauthorized
  - 📁 models // connecting to db
    - 📄 ...
  - 📁 pkg
    - 📁 e // errcode
        - 📄 e.go
    - 📁 infra
        - 📄 redis.go // redis
    - 📁 logic // do some logic for your api
    - 📁 setting // read configuration item from conf/app.ini
    - 📁 utils // maybe put some utils here
  - 📁 routers
    - 📁 api // yep, http api
    - 📄 router.go // router definition
  - 📁 frontend // vue proj



## Reference
 - [gin-vue](https://github.com/Bingjian-Zhu/gin-vue)