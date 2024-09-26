projectName="orange"
projectDir="/data/webapps/orange"

serverUser="root"
serverHost="192.168.0.201"

.PHONY: build
build:
	@CGO_ENABLED=0 go build -ldflags="-w -s" -o ${projectName} .

.PHONY: build-doc
build-doc:
	@ swag init --parseDependency=true

# make build-docker-image
build-docker-image:
	@docker build -t ${projectName}:latest .
	@echo "build successful"

.PHONY: build-windows
build-windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${projectName} .
	@echo "build windows binary successful"

.PHONY: build-linux
build-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${projectName} .
	@md5 go-admin_linux
	@echo "build linux binary successful"

build-sqlite:
	go build -tags sqlite3 -ldflags="-w -s" -a -installsuffix -o ${projectName} .

.PHONY: deploy
deploy:
	@make build-doc
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${projectName} .
	@md5 ${projectName}
	@ssh ${serverUser}@${serverHost} "mkdir -p ${projectDir}"
	@rsync -vI ${projectName} ${serverUser}@${serverHost}:${projectDir}/${projectName}
	@rm -rf ${projectName}
	@scp resources/config.dev.yaml ${serverUser}@${serverHost}:${projectDir}/config.yaml
	@scp scripts/deploy.sh ${serverUser}@${serverHost}:${projectDir}/deploy.sh
	@ssh ${serverUser}@${serverHost} "cd ${projectDir} ; sh deploy.sh"
	@echo "upload ${projectName} success"


#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t go-admin:latest

# make deploy
deploy:

	#@git checkout master
	#@git pull origin master
	make build-linux
	make run