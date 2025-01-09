# URL Shortener in Go

A simple and efficient URL shortener built with Go, HTML templates, and Redis for data storage. This application allows users to input a long URL and generate a shorter version. The shortened URL redirects users to the original URL when accessed.

## Features
- **Shorten URLs**: Input any valid URL to generate a shortened version.
- **Redirect Functionality**: Use the shortened URL to seamlessly redirect to the original link.
- **In-Memory Storage**: Redis ensures quick data access and efficient URL storage.
- **Simple Interface**: Clean and minimal HTML templates for a straightforward user experience.
- **Built with Go**: Uses Go's `net/http` package for handling requests and responses.

## Prerequisites
- [Go](https://go.dev/dl/) (version 1.20 or later recommended)
- [Redis](https://redis.io/download) installed and running locally or on a server

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/ashab-k/go_url_shortner.git
   cd go_url_shortner
