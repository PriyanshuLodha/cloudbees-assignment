CLOUDBEES Project
Overview
Welcome to the CLOUDBEES project! This repository comprises the server and client 
components of a sophisticated train service application. The server manages user interactions, ticketing, 
and offers a range of functionalities. The client, on the other hand, communicates with the server to carry 
out actions such as adding users, purchasing tickets, and viewing user details by section.

Project Structure
plaintext
Copy code
CLOUDBEES/
│
├── server/
│   ├── main.go
│   └── main_test.go
│
├── client/
│   ├── main.go
│   └── main_test.go
│
├── run_server.bat
└── run_client.bat
Instructions
Server
Run Server Tests:

Execute run_server.bat to run the server tests.
bash
Copy code
C:\path\to\project\> run_server.bat
The script will provide feedback on the success or failure of the server tests.
Run Server:

If the server tests pass, execute main.go to start the server.
bash
Copy code
go run server/main.go
Client
Run Client Tests:

Execute run_client.bat to run the client tests.
bash
Copy code
C:\path\to\project\> run_client.bat
The script will provide feedback on the success or failure of the client tests.
Run Client:

If the client tests pass, execute main.go to launch the client.

go run client/main.go
Author
Priyanshu Lodha
