# aws-letter-sender
Program which sends many letters with aws engine

## Fast links
- [How-to-run](##how-to-run) - 

## How to run
### Runnig
1. Prepare `.env` file (You will do it only 1 time). See `example.env` file and fill `.env` file, or run program with `env` params.

2. Prepare `data.csv` file. Example:
```csv
EMAIL,name,TEMPLATE_FILE,SUBJECT
example@example.com,Dias1c,template/example.html,My Example subject
example@example.ccom,MyName,template/example.text,Text letter
example@example.ccom,,,
```
> Explaining. We will use this file for template to fill it and send it to emails. This file is `csv` type, and first line contains only `keys`, and all next lines contains values to `keys`.

Keys:
- `EMAIL` - 
- `name` - 
- `TEMPLATE_FILE` - 
- `SUBJECT` - 

> Program uses go builtin [`text/template`](https://pkg.go.dev/text/template), [`html/template`](https://pkg.go.dev/html/template) (See the usage examples).

3. Run:
The program for sending letters based on the submitted data
```
go run ./cmd/quick/main.go --data-file="data.csv"
```

On running program, it uses params. And this params we can set from different ways. Params also has priority.
Priority of params from:
```md
1. Data file
2. Args
3. .env file
```
> If we set `subject` as flag argument, and set it in `data.csv` file, program will use `subject` from `data.csv`.