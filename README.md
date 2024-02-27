# Random Image

The Random Image is a lightweight, restful API designed to return image URLs from a specified category. Built with Go and leveraging GitHub as an image repository and jsDelivr as a CDN, this service efficiently serves random images suitable for a wide range of applications, from dynamic website content to testing and placeholders.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (Version 1.18 or later recommended)
- Access to a GitHub repository for storing images

### Installation

1. **Clone the repository**

```sh
git clone https://github.com/yourgithubuser/random-image-service.git
cd random-image-service
```

2. **Set up your image repository**

Upload your images to a GitHub repository, organizing them into folders/categories as needed.

3. **Run locally**

```sh
go run *.go
```

This starts the service locally. You can access it at `http://localhost:8080/image/random/{category}`.

## Usage

To get a random image, make a GET request to the service with the desired category:

```http
GET /image/random/{category}
```

The service will return a JSON response with the image URL:

```json
{
  "image": "https://cdn.jsdelivr.net/gh/yourgithubuser/yourrepo@latest/category/image.jpg"
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

