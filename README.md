## API server to access geolocation data

### How to run

1. `export GOPRIVATE=github.com/goku321`

2. `export ACCESS_TOKEN={YOUR_GIT_ACCESS_TOKEN}`

3. `make start-server`

    - It spins up a postgres container and runs `pg-init.sql` on startup. It also starts a api server container which serves on port `:8080` and talk to postgres container to look for requested ip.

    - Hit the api

    ```
    ▶ curl http://localhost:8080/geolocation/200.106.141.15
    {"IP":"200.106.141.15","CountryCode":"SI","Country":"Nepal","City":"DuBuquemouth","Latitude":-84.87503094689836,"Longitude":7.206435933364332,"MysteryValue":7823011346}
    ```

### Stop the server
`make stop-server`

### Run e2e test
`make test-e2e`

### How to do local development
1. start the database container (add more rows inside `pg-init.sql` if needed)

    `docker-compose up postgres`

2. make changes to api and start the server

    `docker-compose up api-server`

3. Hit the api using a browser or CURL

