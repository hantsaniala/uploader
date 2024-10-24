# Uploader

Just an helper tool that will upload the content of a folder across multiple FTP server.

**Important:** Only FTP server that use linux filesystem is supported.

## Installation

Just grab the latest binary release that is compatible to your OS from [here](https://github.com/hantsaniala/uploader/releases/latest) and put it anywhere you want.

## Configuration

Before running this application, you need a `config.json` file next to your binary. The following should be the content structure of your config file:

```json
{
  "source": "source",
  "destination": "/tmp",
  "server": [
    {
      "host": "ftp.example.com",
      "username": "username",
      "password": "Supe3RP@ssw0rd",
      "port": 21
    }
  ]
}
```

By default, it will be automatically generated if it doesn't exist. If that's your case, update the config file corresponding your credentials and folder configuration and just re-run the app again when everything is ready.

## Folder 

### Source

Depending on your `config.json` file, the source folder will be created if doesn't exit.

Put all file that you want to upload inside it.

### Destination

Must be linux based path and not end with slash.

For now, this folder must be created first.

## Run

Just double click to run this app. As simple as that. 😉

## Author

[Hantsaniala](https://t.me/hantsaniala3) 2024