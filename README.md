Golang Hosted Chatbot
=====================

This project is a simple web-based chatbot hosted using Gin in Golang. It communicates with a Python backend to process user input and return responses.

Project Structure
-----------------
```php

├── main.go                  # Golang server with Gin framework
├── static/                  # Static files (HTML, CSS, JS)
│   ├── index.html           # HTML file for frontend interface
├── .gitignore               # Files to be ignored by Git`
```
Features
--------

-   **Golang Server**: A backend server powered by Gin, serving the frontend and processing chat requests.
-   **Python Backend Communication**: The chatbot requests responses from a Python service running on a specified URL.
-   **Frontend Interface**: A basic HTML form for the user to input queries, with JavaScript to handle asynchronous communication.

Setup
-----

### 1\. Clone the Repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/RahulPatnaik/Golang-hosted-chatbot
cd Golang-hosted-chatbot
```
### 2\. Install Golang

Make sure that Go is installed on your system. If not, you can download it from the official Go website: <https://golang.org/dl/>

### 3\. Set Up Python Backend

The chatbot requires a Python backend running on port `8000` (or whichever port you set). If you haven't set up the Python backend, make sure to follow these steps:

1.  Clone or set up your Python backend service.

2.  Run the Python backend:

    ```python app.py  # or the appropriate command to run your Python server```

Ensure that the Python backend is running on `http://127.0.0.1:8000` (or adjust the URL in `main.go` if needed).

### 4\. Run the Golang Server

After installing Go and setting up the Python backend, you can run the Golang server:

bash

Copy code

```go run main.go```

This will start the server on port `8080` (or whichever port you set). You can now access the chatbot via:

```http://127.0.0.1:8080```

### 5\. Interact with the Chatbot

Once the server is running, you can open your browser and go to the URL above. You will be able to interact with the chatbot by typing messages into the input field and receiving responses from the Python backend.

Commands Overview
-----------------

### Run Golang Server

bash

Copy code

`go run main.go`

Notes
-----

-   The frontend interface (`index.html`) is located in the `static/` directory. You can modify this file to change the look and feel of the chatbot interface.
-   The `main.go` file contains the logic for the Gin server. You can customize routes and the interaction with the Python backend as needed.
-   Make sure your Python backend is properly configured to accept requests from the Golang server.
