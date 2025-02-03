# HNG-PROJECT-2-GO-backend-stage-1

## Description
This is the second intership project for backend track stage 1. This project allows users to enter an integer and the project returns the following in a JSON format
- number
- is prime number 
- is perfect number
- its properties eg(["odd", "even"]) etc
- sum of the digits of the number and
- Fun fact from the numberApi

----
## Setup Instruction
### Prerequisites
Ensure you have the following installed:
- **Go** (`1.22` or later) → [Download Go](https://go.dev/dl/)
- **Git** → [Download Git](https://git-scm.com/downloads)
- **VS Code** (Recommended) → [Download VS Code](https://code.visualstudio.com/)

### Installation & Running the Project

1. **Clone the repository**:
   ```sh
   git clone https://github.com/kiddo9/HNG-PROJECT-2-GO-backend-stage-1.git
2. **Navigate to the project directory**:
    ```sh
    cd HNG-PROJECT-2-GO-backend-stage-1

3. **Start the server locally**
    ```sh
    go run main.go
4. **Test the API**:
    *Open Postman or run the following cURL command*:
    ```sh
    curl -X GET http://localhost:8080

***API Documentation***:
1. Local URL:
    ```
    http://localhost:8080

2. Online URL:
    ```
    hng-project-2-go-backend-stage-1-production.up.railway.app

3. Get Intern Details
    
    Endpoint: 
    This is the url endpoint. you can change the number or integer to any number of choose
    ```
    GET '/api/classify-number?number=371'


 4.  Expected Response Format:
        ```
        {
            "number": 371,
            "is_prime": false,
            "is_perfect": false,
            "properties": ["armstrong", "odd"],
            "digit_sum": 11,  // sum of its digits
            "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371" //gotten from the numbers API
        }

5. Example Usage (cURL):
    ```
        curl -X GET http://localhost:8080/api/classify-number?number=371

Folder Structure

    HNG-PROJECT-2-GO-backend-stage-1/
    |-- controllers /
    |   ------ mathController.go
    |-- models /
    |   ------ mathStruct.go
    │── main.go
    │── go.mod
    │── README.md

Programming Language/Stack backlink:
```sh
https://hng.tech/hire/golang-developers

