# lfm

lfm is a last.fm client written in go

## Usage

Inside a root of project run

```bash
 go build
 go install
```

To list what users have been listening, you can issue

```bash
lfm -u goranche,poohica,milann89 -p 12month
```

To get top tracks from artist you can run

```bash
lfm -a Drake
```

Currently working on piping commands so you can do things like this

```bash
lfm -t rap | lfm -a Drake
```

## Idea 

I created this project cause I wanted to learn go. I will try to cover as much of the language as possible. That is why some parts of the code will be overengineered. Topics I want to cover are:

- basic structures
- concurency
- io
- testing
- logging
- regex
- json
- package organization
- os tooling
- go lang patterns


## License

The MIT License (MIT) - see LICENSE for more details
