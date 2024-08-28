# Beego-catgallery

## Prerequisites

- Go installed (minimum version: 1.21)
- Beego and Bee tool installed (bee)
- (Optional) A working GOPATH environment variable. If not set, instructions are provided.

## Project Structure

```bash
beego-catgallery/
├── conf
│ └── app.conf
├── controllers
│ └── default.go
├── main.go
├── models
├── routers
│ └── router.go
├── static
├── tests
│ └── default_test.go
├── views
│ └── index.tpl

```

## Setup and installation

1. **Clone the repository**

   ```sh
   git clone https://www.github.com/nthalt/beego-catgallery
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

<!-- 4. **Move to `src` directory:**

   ```sh
   cd $GOPATH/src/
   ``` -->

4. **Install project dependencies:**

   ```bash
   go mod tidy
   ```

5. **Copy the Sample Configuration File**

   Copy `app.sample.conf` to `conf/app.conf` and change relevant values:

   ```bash
   cp conf/app.sample.conf conf/app.conf
   ```

6. **Run the project:**

   ```sh
   bee run
   ```

The application will be running at [http://localhost:8080](http://localhost:8080)
