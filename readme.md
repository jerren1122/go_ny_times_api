#overview
This is a POC of nesting 1 api interaction inside another

The surfaced API is written in Go, The hidden API interaction is written in Ruby

The surfaced API(Go) passes a date to the hidden API Interaction(Ruby) to the NYTimes NYT retrieves the results for us and the ruby script chooses to only return the abstract of the article for the given year. 

This allows us to build a network of API Connections that are not dependent on knowing how the hidden
API's function. 

#building the go project
This link is very helpful: https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
Additionally if you have go installed 
utilizing 'go build' and './go-ny-times-api' (as this is the name of generated executable after build is called) will locate the main package and execute accordingly. 

#debugging
VSCode has a debugging functionality for Go. We can run the go script in the debugger and then interact with the API via another tool like postman this will allow us to hit the breakpoints. 

To interact in Postman once the API is up and running we would contact: 
http://localhost:8081/get-article (as this is running on your local)
for the body we would pass {
                           "Year": "1989" 
                           }
                           
                           
