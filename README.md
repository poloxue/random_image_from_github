# Random Image

The Random Image Service is a lightweight, serverless API designed to randomly select and return image URLs from a specified category. Built with Go and leveraging GitHub as an image repository and jsDelivr as a CDN, this service efficiently serves random images suitable for a wide range of applications, from dynamic website content to testing and placeholders.

## Features

- **Serverless architecture** for scalability and cost-effectiveness.
- **Category-based image selection** to return relevant images.
- **Integration with GitHub** for easy image management.
- **Utilization of jsDelivr CDN** for fast and reliable image delivery.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (Version 1.18 or later recommended)
- Access to a GitHub repository for storing images
- An AWS account if deploying the service to AWS Lambda (optional)

### Installation

1. **Clone the repository**

```sh
git clone https://github.com/yourgithubuser/random-image-service.git
cd random-image-service
```

2. **Set up your image repository**

Upload your images to a GitHub repository, organizing them into folders/categories as needed.

3. **Configure the service**

Edit the service configuration to point to your GitHub repository and set any necessary parameters such as the AWS credentials if deploying to AWS Lambda.

4. **Run locally**

```sh
go run main.go
```

This starts the service locally. You can access it at `http://localhost:8080/image/random/{category}`.

### Deployment

To deploy this service to AWS Lambda, follow these steps:

1. Package your application.
2. Use the AWS CLI or AWS Management Console to create a new Lambda function.
3. Upload your package and set the handler to your function.
4. Configure API Gateway to expose your Lambda function.

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

