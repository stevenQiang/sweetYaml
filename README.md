# sweetYaml
read yaml file use env data or default value

# Introduction

I use [yaml](https://github.com/go-yaml/yaml) encountered some problems
- I want use env data
- I want use default value

# Solve
I checked a lot of information, and can solve two problems
- in yaml file you can use `${DATABASE_HOST}` can incoming.
```
 DATABASE_HOST=localhost run main.go
```
- default value in main.go
```
  var defaultBaseDBConfig = DBConfig {
    Host: "loclhost", 
    Port: "5432", 
    Username: "postgres", 
    Password: "postgres",
    Database: "postgres",
  }
```
- in this demo, i use two layers yaml

# Note
- This is a demo, if you have questions, you can contact me.
- If you use my code, i hope you can give me a start