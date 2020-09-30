# Go Simple Clean Architecture
**Personal version of clean architecture boilerplate.**  


## About
I created this project to learn more about clean architecture
but i simplify it so it probably look more beginner-friendly.  
I build it with a bottom-up approach (lol idk how to name it) that means
every rule from the bottom is responsible with the higher level until it reaches the top level,
and higher level can't access their lower level because the more high the level is the more user-level
developer can write the code e.g: `AddUpvoterFromPost`  

## Logic
Now we just focus in internal folder which is the heart of application.  

We already know that most of application use to logic such as:  
- Application-logic  
  which responsible to be the mediator among pure business logic and low level application  
  this logic usually not understandable for product and business team  
  because this logic more tend to application-specific  
  e.g: parse json byte to struct of byte and insert it to database
- Business-logic  
  is logic that every team can understand because this logic can be described by human-level language  
  and usually wrap one or more application logic to construct complex logic.  
  e.g: Reikaa(my lovely waifu) must checkout stuff before she can pay it,  
  and check if Reikaa balance is more or same with stuff total price.  

You get the point right?  

## Architecture

list of rule from low to high:
- infrastructure  
  this layer responsible for low-level operation  
  not contain application and business logic  
  such as:
  - connect this application to port 6379 which is the port for redis-server
  - get default database connection that already set from config file
- interfaces  
  interface mostly used for parse, transform from incoming and outgoing data
  this layer tend to pure application logic  
  such as:
  - parse this byte from request to specified struct type
  - deserialize byte from database connection
- usecase  
  usecase wrap interface to implement higher level data-manipulation
  this layer tend to application but with a bit of business logic
  and single responsibility for each schema  
  such as:
  - user model which contain function insert data and if exist by id  
  - checkout model which contain function count to check how many checkout is created  
- domain  
  finally the most understandable level that wrap one or many usecase model  
  and domain layer is mostly responsible for business logic even the non-tech person understand what are you writing for  
  this function can construct complex logic but with simple input and of course with validation too  
  such as:
  - user A like your photos please call function `AddLikeFromPost`  
  - user C want to delete their account so we will use `PurgeSpecifiedAccount` to delete all their data like: photos, friends which need aggregate and transaction context  


## Folder Structure  
---

```
├── cmd                          # contains implementation from internal core package
│   └── application              # construct server side
│       ├── handler              # handler/controller to access domain logic
│       │   ├── helpers.go       # known human-level return code from domain logic
│       │   └── insert_user.go   # example of implementation domain logic
│       └── server.go            # file endpoint to run the server
├── config                       # config file for specific environment
│   └── dev.env
├── docs                         # docs file like swagger, insomnia, etc
├── internal                     # heart of this application i already notice about
│   ├── domain
│   ├── infrastructure
│   ├── interfaces
│   └── usecase
├── main.go                      # this file used to run main application from cmd and autodetect to build server like heroku
├── makefile                     # few shortcut command to dockerize, build, and run app
├── pkg                          # shareable package idk but i felt comfort while creating this
│   ├── structs
│   ├── system
│   └── transform
└── store                        # folder i create for storing seed (as a json file)
```

## Technologies
- [Echo](github.com/labstack/echo)
- [Resty](github.com/go-resty/resty)
- [Json-iterator](github.com/json-iterator/go)
- [Kamva](github.com/Kamva/mgm/v3)


### Sorry But i'll update this soon
  send issue if my language can't understandable
