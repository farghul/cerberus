# Cerberus

Cerberus is an interactive CLI tool for S3 Object Storage. It is named after the [multi-headed dog](https://en.wikipedia.org/wiki/Cerberus) that guarded the gates of the underworld in Greek mythology.

![Dog](cerberus.webp)

## 📚 Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

## 🏭 Function

Cerberus creates an S3 client to facilitate the interaction with object store buckets and objects.

## 📂 Project Structure

Inside of your Cerberus project, you'll see the following files:

``` zsh
.
├── .gitignore
├── cerberus.webp
├── go.mod
├── go.sum
├── LICENSE.md
├── main.go
├── README.md
├── tasks.go
├── tools.go
```

## 🚧 Build

From the root folder containing `main.go`:

``` zsh
go build -o cerberus .
```

## 🎏 Available Flags

| Command|Action|
|:-|:-|
|`-help`|Display help information|
|`-list`|List all buckets|
|`-create`|Create new S3 bucket(s)|
|`-bktd`|Delete S3 bucket(s)|
|`-objd`|Delete s3 object(s)|
|`-get`|Download s3 object(s)|
|`-put`|Upload S3 object(s)|
|`-large`|Upload large S3 object(s)|

## 🏃 Run

``` zsh
cerberus -put [BUCKET-NAME] [SOURCE] [DESTINATION]
```

## 🎭 Example

``` zsh
cerberus -put my-bucket ~/Pictures/image.jpg photos/2024/image.jpg
```

## 🎫 License

Code is distributed under [The Unlicense](https://github.com/farghul/cerberus/blob/main/LICENSE.md) and is part of the Public Domain.