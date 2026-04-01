# running the base scripts initial ( before adding validaitions )
http POST http://localhost:8082/api/students

# running the base scripts initial ( after adding validaitions ) - name , email and age 
http POST http://localhost:8082/api/students id=1 name=rakesh email=a@b.com age:=12
 - the age must be added with := 
    - age = 12 --> "age"="12" (string)
    - age := 12 --> "age"=12 (integer)
