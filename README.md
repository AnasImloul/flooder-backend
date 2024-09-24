# flooder-backend
A Golang-based backend for the flooder game.

## Installation

### Prerequisites
Make sure you have the following installed on your system:
- [Docker](https://www.docker.com/get-started)

### Steps to Build and Run the Project

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/flooder-backend.git
   cd flooder-backend
   ```

2. **Build the Docker Image**:
   Use the provided Dockerfile to build the Docker image.
   ```bash
   docker build -t flooder-backend .
   ```

3. **Run the Docker Container**:
   After building the image, run the container using the following command:
   ```bash
   docker run -p 8080:8080 flooder-backend
   ```

   The backend service will be accessible at `http://localhost:8080`.

## Contributing

Feel free to fork this repository and submit pull requests for improvements or bug fixes.