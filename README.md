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

Import `api.postman_collection.json` in postman, Provide TOKEN=<JWT TOKEN FROM LOGIN\> in Headers

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

***Health***
- - ***Get Health***: GET <services>/h - Check the health of services.

### Improvements & Updates needed
- Handle the services intraction with gRPC or kafka
- Better file structure & code re use for all services
- Handle cloud deployment automatically (pipelines & build) and configure servers for scalablity
- Unit & Integration Tests

### Sample Architecture
![image](https://github.com/user-attachments/assets/d4ac0ffe-11c5-4aed-b7d8-f46a80683dc6)

Why different services ?
- User auth service is related to only user stuff, Future improvements like OTP verification, OAuth can be integrated easily, This can also be used for org's Other services where auth is needed 
- Expense Tracker services has bit of more processing than user due to DB quries & validations, Which is also independent & can improve/scale easily in future  
- File manager can take more computation than user auth & expense tracker, soo Its better to handle it in difference service, Will be helpful to handle file related updates in future as well


### Sample excel screenshot 
(ignore the user names and values)

![image](https://github.com/user-attachments/assets/872e54e7-679e-4f68-ab06-0b046e92aa23)

Get All users expenses (related to the creator only, Not everyone in database)

![image](https://github.com/user-attachments/assets/338f522d-d9ca-471b-9c67-bba059bf002d)

Get my expenses 
