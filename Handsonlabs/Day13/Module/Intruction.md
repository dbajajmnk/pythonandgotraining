# Day 13
## Modules Lokeshwara Simple (Congralutions )
  1. go mod init namemodule 
  2. it create file with go.mod file
  3. [Go Package Link](https://pkg.go.dev/) take some package form this sit
  4. go get modulename (replace this name with your package name which select form above site)
  5. go.sum file 
  6. you can import the module/pakages inside your code file main.go
## Package  Ranjan and Sai Pavan (Configurations)
  1.  Create new folder for your project
  2.  Perform module steps
  3.  Create directory as follow
      1.  cmd (Folder)
          1.  app
              1.  main.go
      2.  pkg (Folder)
          1.  util (Folder)
              1.  validationutil.go
      3.  internal (Folder)
          1.  service (Folder)
              1.  myService.go
  4.  Follow
      1.  Util file has been used in service
      2.  Service code as used in main 
  5.  Your applicaiton should as it
## File Handling Harivardan and Sunil
  1. Create File
     1. file,error = os.Create("fileName")
     2. check error if
     3. stop or kill the app
     4. close you file safety with defer close on your file object
  2. Write File
     1. need to call WriteString and pass string which need to write
  3. Append
     1. intiaize object os OpenFile method with following 
        1. filename
        2. Mode os.Append | Wrongly 
        3. 0664
     2. it open a file in append mode
        1. if error go for panic
     3. file.WriteString("asdsadsasaasd\n")
     4. close file with differ
  4. Buffer
     1. create bufferobject with bufio.NewWriter(file)
     2. object.WriteString("asddas")
     3. object.Flush()
  5. Read file 
     1. os.ReadFile(file)
     2. data and error
     3. string(data)
##  JSON Handling  Hari priya and Pavan (Congratultions)
     1. need to import encoding/json 
     2. Create Two Stuct 
     3. Do mapping or Seriralation
     4. json.MarshalIndent
        1. Strcut
        2. prefix
        3. Indent
     5. Need to create file using os.Create(filename.json)
     6. close with defer
     7. json.NewEncoder(fileReference).Encode(struct with value)
   
## Logging Ahish and Mrudula
   1. log 
      1. SetPrefix
      2. SetFlags
         1. log.Ldatae
         2. log.Ltime
         3. log.Lmicroseconds 
      3. Println() show logs on tremail using logger object
   2. Putting logs in file 
      1. Open file with 
         1. crate
         2. wrongly
         3. append
         4. 0664
      2. close the file with defer f.close()
      3. use this file as output for your logger
         1. SetOutput(filereference)
      4. Add log after that with Println
   3. slog
      1. Create a Handler
         1. slog.NewJSONHandler
            1. os.Stdout,
            2. &slog.HanlderOptions{Leavel:slog.LevelInfo})
   
      2. Logger
         1. slog.New(handler)
      4. reqLogger
         1. Attach a id
            1. create id variabl
         2. newObjectfor this
            1. logger.With("",idvariable)
      5. Put Error in logs
      6. Warn
      7. Info
   

   



  . 




Create a module with go mod init module
h.go.dev/
go to google and find some go module and check how we can those
Add into our project 
Check the impacted files 
Run your code 