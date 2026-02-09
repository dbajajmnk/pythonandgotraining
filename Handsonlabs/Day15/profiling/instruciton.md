## Prfiling
### Imports
    - imports net/http
    - imports net/http/pprof
    - time
    - log
### Setup a run with go Routine
    go func(){
        log.
        log. (http.listenAndServer("))
    }
## Apply a load
    - Iteration over million of records
    - You sleep method of time to apply some delay

### How To Run 
  - Run your server first
  - Check the profile with pprof in form of test out
    - go tool pprof http://localhost:6060/debug/pprof/profile
  - Check Profile with pprof in the form of UI out (Not working right now)
    -  go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile