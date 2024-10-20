# Rule Engine with AST
## Introduction
This application serves as a Rule Based Engine for validating information through JSON format. It is written in Golang and utilizes Abstract Syntax Trees to validate JSON data.

## Tools used
1. MongoDB : To store json data
2. Golang: To create a server and server static files.

## Why use MongoDB?
1. It is a NoSQL database, providing flexible schema to dump irregular Abstract Syntax Trees.
2. MongoDB’s document-based model allows storing complex data.

## Functionalities:
### Rule Addition
- Rule addition requires users to send their rule string through the web interface, which gets validated and converted to AST and gets stored in the DB.
- Rule addition can be through API `http://localhost:8080/create`
    ``` bash
    POST /create
    body : {
        rule : "((age > 30 AND department = 'Sales') OR (age < 25 ANDdepartment = 'Marketing')) AND (salary > 50000 OR experience >5)"
    }
    ```

### Data Evalation against Rules
- Data can be evaluated against the created rules, by sending them in a JSON format from the frontend
- Data can also be evaluated through API `http://localhost:8080/evaluate`
    ``` bash
    POST /evaluate
    body : {
        "rule_id":"67137426a45b9f271a4b3ec8",
        "data":{
		    "age":        25,
		    "department": "Sales",
		    "salary":     5000000,
		    "experience": 8
	    }
    }
    ```

### Combining Rules
- Rules can also be combined into a single AST using `OR` conjunction. Once the user adds a rule, immidiately they are given the option to combine it with any existing rule
- Rules can also be combined through API `http://localhost:8080/combine`
    ``` bash
    body: {
        "rules" :[
            "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience >5)",
            "((age > 30 AND department = 'Marketing')) AND (salary >20000 OR experience > 5)"
        ]   
    }
    ```

## What is the flow of the application?
1. User heads to `http://localhost:8080/static/`and accessses the web interface.
2. User enters the rule into the input box and hits submits
3. The rule is validated at the backend, and an AST is constructed through a parser.
4. The AST is converted to a map type and inserted into MongoDB.
5. The user is notified of the successful operation.
6. The user is prompted to combine the newly added rule to existing rules.
7. The user enters a JSON data that is to be evaluated, in the input box below the rule.
8. The user is recieves response whether the JSON data satisfies the rules or not through a boolean value. 

## Requirements to run the application
1. Go [How to install go?](https://go.dev/doc/install)
2. Docker (Optional) [How to install Docker?](https://docs.docker.com/engine/install/)
3. MongoDB Atlas cluster [How to create MongoDb Altas Cluster?](https://www.mongodb.com/docs/guides/atlas/cluster/)
4. MongoDB Connection string [How to get MongoDB connection string?](https://www.geeksforgeeks.org/how-to-get-the-database-url-in-mongodb/)

## How to run the application through go compiler
1. Clone this repository
2. Open the project directory locally
3. Create a `.env` file.
4. Enter the following details in the .env file
    ``` 
    MONGO_URL=<Your MongoDB connection url>
    ```
    Example 
    ``` 
    MONGO_URL=mongodb://username:password@host1:port1,host2:port2/database?option1=value1&option2=value2
    ```
5. To test the application you can run
    ``` bash
    go test ./...
    ```
6. Run the following command to generate a go executable
    ``` bash
    go build -o main .
    ```
7. A go executable called `main` must have been generated, to run the executable, run the following on your termninal
    ``` bash
    ./main
    ```
8. If the application has started successfully, you'll get following message. By default, the server will run on port 8080
    ``` bash
    2024/10/20 09:32:11 Pinged your deployment. You successfully connected to MongoDB!
    2024/10/20 09:32:11 Server Started at port 8080
    ```
9. Head over to `http://localhost:8080/static/` to launch the web interface

## How to run the application through Docker
1. Clone this repository
2. Open the project directory locally
3. Create a `.env` file.
4. Enter the following details in the .env file
    ``` 
    MONGO_URL=<Your MongoDB connection url>
    ```
    Example 
    ``` 
    MONGO_URL=mongodb://username:password@host1:port1,host2:port2/database?option1=value1&option2=value2
    ```
5. Run the following command to build a docker image
    ``` bash
    docker build -t rule-engine .
    ```
6. Once the docker image is created run the following command to start the server at portt 8080
    ``` bash
    docker run -it -p 8080:8080 --rm --name rule-engine rule-engine
    ```
7. Once the server starts successfully, you'll get following response
    ``` bash
    2024/10/20 04:09:56 Pinged your deployment. You successfully connected to MongoDB!
    2024/10/20 04:09:56 Server Started at port 8080
    ```
8. Head over to `http://localhost:8080/static/` to launch the web interface, which instantly subscribes to the weather updates

## Screenshots
### Web Interface

