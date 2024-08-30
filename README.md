# Beego-catgallery

Beego-catgallery is a web application that showcases cat images using The Cat API. It allows users to view random cat images, browse cat breeds, vote on their favorite cats, and save cats to their favorites list. This project demonstrates the use of the Beego web framework for Go, along with frontend interactivity using JavaScript.
This application uses Go channels for API calls to demonstrate concurrent processing.

- [Prerequisites](#prerequisites)
- [Features](#features)
- [Project Structure](#project-structure)
- [Setup and Installation](#setup-and-installation)
- [Important Notes](#important-notes)
- [Contributing](#contributing)

## Prerequisites

- Go installed (minimum version: 1.21). If needed, install from here [https://go.dev/doc/install](https://go.dev/doc/install)
- Beego and Bee tool installed (bee)
- A working GOPATH environment variable. If not set, instructions are provided.

## Features

- Display random cat images
- Browse cat breeds and view breed information
- Vote on cat images (upvote/downvote)
- Add cats to favorites
- View favorite cats in a gallery

## Project Structure

```bash
beego-catgallery/
├── conf/
│   ├── app.conf
│   └── app.sample.conf
├── controllers/
│   ├── cat_api.go
│   └── default.go
├── routers/
│   └── router.go
├── static/
│   ├── css/
│   │   └── style.css
│   └── js/
│       ├── main.js
│       └── reload.min.js
├── tests/
├── views/
├── .gitignore
├── beego-catgallery
├── go.mod
├── go.sum
├── main.go
└── README.md

```

## Setup and installation

1. **Clone the repository**

   ```sh
   git clone https://github.com/nthalt/beego-catgallery
   cd beego-catgallery
   ```

2. **Install Go Dependencies**
   Ensure you have `bee` installed:

   ```sh
   go install github.com/beego/bee/v2@latest
   ```

3. **Add `GOPATH` to your PATH:**

   Set GOPATH Temporarily (For the Current Terminal Session)

   ```sh
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin
   ```

   or

   Set GOPATH Permanently (For All Sessions)

   ```sh
   # for bash
   echo 'export GOPATH=$HOME/go' >> ~/.bashrc
   echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

   ```sh
   # for zsh
   echo 'export GOPATH=$HOME/go' >> ~/.zshrc
   echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
   source ~/.zshrc
   ```

4. **Install project dependencies:**

   ```bash
   go mod tidy
   ```

5. **Copy the Sample Configuration File**

   Copy `app.sample.conf` to `conf/app.conf` and change `cat_api_key` and `user_id`. `user_id` is used to separate users:

   ```bash
   cp conf/app.sample.conf conf/app.conf
   ```

6. **Run the project:**

   ```sh
   bee run
   ```

The application will be running at [http://localhost:8080](http://localhost:8080)

## Important Notes

- The `user_id` in the configuration is used as the `sub_id` for votes and favorites. Used to separate users.

<details>
<summary>

## Contributing

</summary>

We welcome contributions to this project. To ensure a smooth collaboration, please follow these guidelines:

1. **Fork the Repository**: Start by forking the repository on GitHub.

2. **Clone the Repository**: Clone your forked repository to your local machine using:

   ```bash
   git clone https://github.com/username/property-manager-django.git
   ```

3. **Create a Branch**: Create a new branch for your feature or bug fix:

   ```bash
   git checkout -b feature-or-bugfix-description
   ```

4. **Make Changes**: Implement your changes in the codebase. Ensure your code adheres to the project's coding standards and includes appropriate tests.

5. **Commit Changes**: Commit your changes with a clear and descriptive commit message:

   ```bash
   git add .
   git commit -m "Description of the feature or bug fix"
   ```

6. **Push to GitHub**: Push your branch to your forked repository on GitHub:

   ```bash
   git push origin feature-or-bugfix-description
   ```

7. **Create a Pull Request**: Go to the original repository on GitHub and create a pull request. Provide a clear and detailed description of your changes.

8. **Review Process**: Wait for the project maintainers to review your pull request. Be prepared to make any necessary changes based on feedback.

Thank you for your contributions! Your help is greatly appreciated.

</details>
