# expense-tracker
A golang based expense tracker backend 

### DOCS

**Commands**

1. Copy the Repo
```
git clone https://github.com/SugamKuber/expense-tracker && cd expense-tracker
```
2. Set env 
```
DB_URI=<postgresql uri>
JWT_SECRET=<long secret>
```

3. Start in Local:
- put ENV VARS in `.env`
- run `make start` to start all servers
- run `make stop` to stop all servers
- run `make check` to check health of all servers
- run `make restart` to restart all servers
- run `make clean` to clean up logs saved

4. Start in docker: (Avoid this, Improvements needed)

- run `docker-compose build` to build images
- run `docker-compose up` to start servers
- run `docker-compose stop` to stop servers
- run `docker-compose down` to remove containers
- run `docker-compose restart` to restart servers

**Api Guide**:

***Auth***

- ***Signup***: POST /signup - Register a new user.
- ***Login***: POST /login - Authenticate user and receive a JWT token.
- ***Get User***: GET /me - Retrieve the authenticated user's information.
- ***Change Password***: POST /change-password - Change the user's password.

***Expenses***
- ***Add Expense***: POST /add - Add a new expense.

Below are the types we can add, check postman doc for full API body
```
split_method:percentage
split_method:equal
split_method:exact
```
- ***Get My Expenses***: GET /track/me - Retrieve the authenticated user's expenses.
- ***Get All Expenses***: GET /track/all - Get all related user's expenses.
- ***Get All Expense as Admin***: GET /track/all/admin - Get every users in db  expenses.

***Download***

- ***Download My Expense***: GET /download/me - Download the authenticated user's expense in excle.
- ***Download All Expenses***: GET /download/all - Download all users' expense in excle.
