# Deptech Test Backend

## How To Run
```
docker compose up -d --build
```

## ERD
![ERD](ERD.png)

## API Doc
[https://www.postman.com/gold-eclipse-18650/workspace/deptech-test-backend/collection/26140093-7062186d-1c17-44e5-86dd-5ca110930a52?action=share&creator=26140093&active-environment=26140093-a55638b6-e14e-4faf-a0f6-5aaeba4fab6a](https://www.postman.com/gold-eclipse-18650/workspace/deptech-test-backend/collection/26140093-7062186d-1c17-44e5-86dd-5ca110930a52?action=share&creator=26140093&active-environment=26140093-a55638b6-e14e-4faf-a0f6-5aaeba4fab6a)

## To Do

### Deploy
- [x] Setup database (mysql) using docker compose
- [x] Setup object storage (minio) using docker compose
- [x] Setup redis using docker compose
- [x] Setup Dockerfile

### Documentation
- [x] Draw ERD
- [x] Write API Doc
- [x] Write how to run app
- [ ] Record demo video

### Setup Repository
- [x] Init project
- [x] Setup codebase using go standard & clean architecture
- [x] Setup go validator
- [x] Test mysql connection
- [x] Test minio connection
- [x] Design User model
- [x] Design Category model
- [x] Design Product model
- [x] Design Transaction model
- [x] Design Transaction Item model
- [x] Encrypt password (bcrypt)
- [x] Setup authentication middleware (JWT)
- [ ] Setup logging

### User Feature
- [x] Create user
- [x] Read user
- [x] Update user
- [x] Delete user
- [x] List user
- [x] Login user
- [x] Logout user

### Category Feature (need auth)
- [x] Create category
- [x] Read category
- [x] Update category
- [x] Delete category
- [x] List category

### Product Feature (need auth)
- [x] Create product
- [x] Read product
- [x] Update product
- [x] Delete product
- [x] List product

### Transaction Feature (need auth)
- [x] Create transaction
- [x] List transaction