
# ascii-art-export

ascii-art-export is a webpage, that allows create ascii representation of input string and export it as file in txt format

Your can choose template for ascii representation:
- standard
- shadow
- thinkertoy

## Usage/Examples
Clone the repository and start the server
```bash
go run ./web
2022/07/18 14:58:52 web server is on http://localhost:8080
```
go to http://localhost:8080





Write your input string:

![App Screenshot](Screenshots/1.png?raw=true "Title")

Choose one of the ascii-template:

![App Screenshot](Screenshots/2.png?raw=true "Title")

Press **SUBMIT** button:

![App Screenshot](Screenshots/3.png?raw=true "Title")

After **SUBMIT** button has been clicked press **Download** button to export file in txt format.



## Implementation details

`main.go` 
Creates multiplexer and starts server on 8080 port.
When the request reaches the server, a multiplexer will inspect the URL being
requested and redirect the request to the correct handler fucntion

`handlers.go`

Handlers fucntions process requests. When the processing is complete, the handler passes the
data to the template engine, which will use templates to generate HTML to be
returned to the client. 

## Authors

- @ddarzox
- @robertt3kuk


