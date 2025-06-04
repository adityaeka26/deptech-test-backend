# Deptech Test Backend

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
- [ ] Write how to run app
- [ ] Record demo video

### Setup Repository
- [x] Init project
- [x] Setup codebase using go standard & clean architecture
- [x] Setup go validator
- [x] Test database connection
- [ ] Test object storage connection
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
- [ ] Create category
- [ ] Read category
- [ ] Update category
- [ ] Delete category
- [ ] List category

### Product Feature (need auth)
- [ ] Create product
- [ ] Read product
- [ ] Update product
- [ ] Delete product
- [ ] List product

### Transaction Feature (need auth)
- [ ] Process transaction
- [ ] List transaction