# ImageSort-Go

ImageSort-Go is a command-line tool for sorting and renaming media files based on their EXIF metadata or by ascending index. It supports various image and video formats.

## Features

- Rename media files based on their EXIF creation date.
- Rename media files by ascending index.
- Supports a wide range of image and video formats, including `.jpg`, `.png`, `.mov`, `.mp4`, and more.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/vin-rmdn/imagesort-go.git
   cd imagesort-go
   ```

2. Build the project:

   ```sh
   go build -o imagesort_go
   ```

## Usage

### Sort Media by EXIF Date

Sort and rename media files in a folder based on their EXIF creation date:

```sh
./imagesort_go sort-media /path/to/destination/folder
```

### Rename Media by Index

Rename media files in a folder by ascending index:

```sh
./imagesort_go rename /path/to/destination/folder
```

## Configuration

The supported file extensions are defined in [`config/config.go`](config/config.go). You can modify the `RecognizedImageExtensions` variable to add or remove supported formats.

## Development

### Prerequisites

- Go 1.24 or later

### Run Tests

To run the tests, use:

```sh
go test ./...
```

### Project Structure

- **`cmd/`**: Contains the CLI commands.
- **`config/`**: Configuration for supported file extensions.
- **`exif/`**: Handles EXIF metadata extraction and date parsing.
- **`media/`**: Media renaming logic.
- **`tool/`**: Utility functions, such as safe renaming.
- **`testdata/`**: Sample test files.

## Dependencies

This project uses the following dependencies:

- [go-exiftool](https://github.com/barasher/go-exiftool): For extracting EXIF metadata.
- [cobra](https://github.com/spf13/cobra): For building the CLI.
- [testify](https://github.com/stretchr/testify): For testing.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## Acknowledgments

Special thanks to the authors of the libraries used in this project.