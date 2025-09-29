# Safe-return-backend

The backend for Safe Return. Handles user accounts, authentication tokens, and
making sure the user gets back home safe. In the app, there will be a *paranoid mode*, where
the app will assume the user is in danger if they go offline, or are unreachable, and will contact
the user's emergency contacts if they are. This is why, when the *paranoid mode* is activated,
the server will need to track the user's location periodically in case they go offline or their time
runs out.

## Run locally

Requirements:
- [Docker](https://www.docker.com/) or [Podman](https://podman.io/)
- `docker compose` or `podman-compose`

To run the server, simply run:  
```sh
docker compose up --build  
```
if you use docker, or  
```sh
podman-compose up --build
```
if you use podman.

> [!NOTE]
> If you're using `podman-compose` without root, then you will get an error message from
> `nginx`. You can safely ignore this, just add `:8080` to any time you might access localhost

## API documentation

The documentation for all the endpoints can be found on `localhost/docs` while the server is
running.  
*(Remember to add the :8080 if you're using Podman!)*
