# TT Ranking Calculator

After a tournament, I often want to work out how many ranking points I have gained (or lost), but they are only published once a month. So, I often find myself on the journey home jumping between the rankings list and points table totting it up in my head. I was tired of it, so I made this.

![Screenshot of site](/assets/screenshot.png "Screenshot of site")

## Development

This project uses [Templ](https://templ.guide/) which requires the CLI to generate the Go from the template files. To install the CLI...

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

...alternatively, other ways of installation are available from [their docs](https://templ.guide/quick-start/installation).

To generate from the templates, run...

```sh
templ generate
```

### Running the Go Binary Directly

Build...

```sh
go build
```

Run...


```sh
./tt-ranking-calculator
```

### Running via Docker

Assuming Docker is setup, build the image...

```sh
docker build . -t RowMur/tt-ranking-calculator
```

...and then to run...

```sh
docker run -p8080:8080 -t RowMur/tt-ranking-calculator
```

## Deployment

This project is self hosted on my [server provided by Hetzner](https://host.rowmur.dev/). There is a webhook triggered on push to master which then triggers a deploy.
