# go-hexagonal
To start the application using Docker Compose, you can run the following command:
```
docker-compose up -d
```
This will start the containers defined in the `docker-compose.yml` file in detached mode.

To enter the container where the `appproduct` service is running, you can use the `docker exec` command with the `-it` flags and specify the container name or ID:
```
docker exec -it appproduct bash
```
This will open a bash shell inside the container, allowing you to interact with it.

Once inside the container, you can use `go mod tidy` to download and manage the necessary Go dependencies. This command will analyze the `go.mod` file and ensure that all required dependencies are present and up to date:
```
go mod tidy
```
Make sure to run this command in the appropriate directory where the `go.mod` file is located.

By following these steps, you will be able to start the application using Docker Compose, access the container, and manage the Go dependencies using `go mod tidy`.

