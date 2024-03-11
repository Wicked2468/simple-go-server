# Simple Go HTTP Server Example

This is a simple Go program that demonstrates how to create a basic HTTP server using the standard `net/http` package.

## Features

- Serves static files from the "./static" directory.
- Handles form submissions on the "/form" path.
- Responds with "Hello!!!!!!" on the "/hello" path for GET requests.

## Usage

1. Clone the repository:
    https://github.com/Wicked2468/simple-go-server.git

2. Run the server:
    go run main.go

3. Access the server in your browser:

- Static files can be accessed by navigating to `http://localhost:8000/static-file-name`.
- The form can be accessed by navigating to `http://localhost:8000/form`.
- The hello message can be accessed by navigating to `http://localhost:8000/hello`.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This example is created for educational purposes.
- Thanks to the Go team for providing an excellent standard library for building web applications.
