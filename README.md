## Product Requirements
 
    
  •From scratch setup a Full Stack application with ReactJS as Frontend and Go as Backend:
- [x] Create the folder structure
- [x] Basic blocks for both frontend(basic app...) and backend(server, system)
- [x] Create a First API endpoint and a First View that shows this endpoint (Library choices are free)
 
 
  •The frontend part shall:
- [x] Implement a routing system
- [x] have an error catching system to catch errors from the virtual DOM
- [x] Implement a welcome page

  •The backend app shall:  
- [x] implement a routing system
- [x] implement a unique endpoint. It shall return "hello world"
- [x] return 404 if an requested endpoint does not exist
- [x] All requests to the API shall be done though HTTP protocol and response body shall be in JSON format

  •Nice to have:   
- [x] he front should look nice and to do that you are free to use any frameworks, but you will need to justify your choices.
- [x] On both part, have a linter installed for local development, we are waiting to see a code readable and understandable.


## Context

The answer to the task is divided into 2 parts

- API - Simple backend service to create data and API's

- CLIENT - Frontend side to show data sent from backend

## Instructions

- Clone this repo.
- Open the project folder using Visual Studio Code

## API side

- A simple api working at http://localhost:8081/articles was produced with go(lang).

### To Start "api"

1. Open a new terminal
2. cd to folder api
3. Execute "go run main.go" to start the server. Server runs in port 8081


---

## CLIENT side

1.  React error boundary package is used for error boundary. [React Error Boundary](https://www.npmjs.com/package/react-error-boundary)


### To Start "client"

1. Open a new terminal
2. cd to folder client
3. Execute "npm install" to get the dependencies
4. Execute "npm start" to start the project.

### That's it!

Salute!

