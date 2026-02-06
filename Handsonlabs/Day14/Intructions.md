## Routing
- Create muxObject
- NewServerMux
- Attach NewHandlerFunc with your mux object
  - Health
  - Users
  - Print
- Write funcitonality for Handler using call back funciton
- Print log 
- Print Log for Start and Serve 
## Middle Ware
- Need create function with parameter of type http.Handler
- return tyep is also same http.Handler
- return a function 
  - http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)){
  - }
  - Write own logic
  - Don't forget to call next.ServeHTTP(w,r)
  - How to Call
## DataBase
- Open Database
  - sql.Open("sqlitelite","file:data14days.sql")
    - error
    - create a object sql
- Create Table
  - Write Query to Create Table in the form of string using string template way
  - Execute with sql.Exe(query)
    - err 
    - success
- Insert Data into Table
  - Write query to insert data 
  - execute it
    -  check error
- Read Data from Table
  - Read data from Table based on query
  - Parse and Print the data 

## Clean API
- 


- ## Clean Rest API
- Separation of concern
- Lose Coupling 
- Dry (Don't Repeat your self)
- Versioning
- Backward Compatiability 
- Scalability
- Consistancy 
- Layering 
- SOLID Principle
- Resource focused 
- Security 

