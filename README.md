
<h1 align="center">Url Shortner System</h1>

## It's an url shortner system ,which is able to short an url from given long url and redirect to the long url address.You also allowed to see all the short urls you have created along with it's orginal long url and id respectively , even you can delete any record through id.Here i have used Mysql Databese to store the records.if you use this Url Shortner System on local server in your pc it would work fine but if you wants to share this link you have upload it on live server.



## 1. Install Packages

- github.com/gorilla/mux
- github.com/jinzhu/gorm
- github.com/jinzhu/gorm/dialects/mysql
- github.com/speps/go-hashids



##  You have to install this packages and you can simply do this by using them as git command on  git bash  

# Connect Database
- I have worked on localhost .My username on mysql database was "root" ,and password was blank "".You have to configure it accordiing hosting and database.
- ( dataSourceName := "your_username:your_password@tcp(localhost:3306)/your_database_name?parseTime=True" )

## There is no need to create database and table manually ,system would create automatically 



# Create Url 
## (router.HandleFunc("/create/url", createurl).Methods("PUT") )Route For Create url 
## If you are on localhost ,You have to use this (http://localhost:8080/create/url) on POSTMAN to create short url. Use PUT Request 
## On Body => 
{
    "longURL":"Your Long url"
}

# Redirect Url 
## ( router.HandleFunc("/{Id}", redirect).Methods("GET") )Route For Redirect url 
## There is nothing to check it  on POSTMAN , if took any of your created short and find it's working ,that's mean this route working properly

# Get Urls 
## ( router.HandleFunc("/urls", geturls).Methods("GET")) Route For Get all url 
## If you are on localhost ,You have to use this ( http://localhost:8080/urls ) on POSTMAN to see all the url short url along with id and long url respectively .Use Get Request 


# Delete Url 
## (router.HandleFunc("/url/{Id}", deleteurl).Methods("DELETE") )Route For Delete url 
## If you are on localhost ,You have to use this (http://localhost:8080/url/Your id to delete ) on POSTMAN to Delete url record from database. Use Delete Request 

 

# Contact & social

- **[Facebook](https://www.facebook.com/safahait.sawon)**
- **[Github](https://github.com/Safayet-Shawn/)**
- **[Gmail:safayetshawn95@gmail.com](safayetshawn95@gmail.com)**
- **[Phone :01726681903](01726681903)**
 

