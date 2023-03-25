#### Docker commands ( Note : the Dockerfile can be improved)

```bash
# build the image
docker build -t {{$PROJECT_NAME}} .
# run and publish with the name of {{ $PROJECT_NAME }}
docker run --publish {{ $GOLANG_PORT }}:{{ $GOLANG_PORT }} --name {{ $PROJECT_NAME }} {{ $PROJECT_NAME }}
# stop
docker stop {{ $PROJECT_NAME }}
# remove 
docker image rm {{ $PROJECT_NAME }}
```
