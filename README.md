# Deptech Test Backend

## How To Run
```
docker compose up -d --build
```

## To Do

### Deploy
- [x] Setup database (mysql) using docker compose
- [x] Setup object storage (minio) using docker compose
- [ ] Setup redis using docker compose
- [x] Setup Dockerfile
- [ ] Deploy to VPS using docker
- [ ] Setup DNS
- [ ] Setup nginx
- [ ] Setup HTTPS

### Documentation
- [ ] Write ERD
- [ ] Write API Doc
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
- [ ] Logout user

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
- [ ] Process transaction
- [ ] List transaction