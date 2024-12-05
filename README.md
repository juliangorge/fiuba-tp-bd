# fiuba-tp-bd

### Stack:

Frontend - NextJS (React)

Backend - GO

### Integrantes:
- Julián Gorge - 104286
- Martin Pata Fraile de Manterola - 106226
- Joaquín Velurtas - 109655
- Rodrigo Souto - 97649

## Instructions

1. Go to the root directory of the repo.

2. Stop all running containers (to prevent conflicts)
```
docker-compose stop
```

3. Setup the `.env` file, here is a dummy working example:
```
POSTGRES_PASSWORD=postgres123
POSTGRES_USER=postgres
POSTGRES_DB=postgres
POSTGRES_HOST=postgres
MONGO_INITDB_ROOT_USERNAME=root
MONGO_INITDB_ROOT_PASSWORD=toor
MONGODB_HOST=mongodb
```

4. Build and launch the services and databases.
```
docker-compose up --build`
```

5. Now you can access the website with to http://localhost:3000
