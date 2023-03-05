### Project setup

```
go run main.go

Listening on localhost:9527

```

### API

```

path /file

query:
* localSystemFilePaty: your path of file
* orderBy: [fileName, lastModified, size]
* orderByDirection: [Descending, Ascending]
* filterByName: your file name

```

### Docker

```
# build docker image

docker build . -t "image_name"

docker run -it --name "container_name" -p "your_port":9527 "image_name"

```
