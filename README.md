# aws-letter-sender

Email sender using [AWS SES](https://aws.amazon.com/ses/) with [HTML](https://pkg.go.dev/html/template)/[text](https://pkg.go.dev/text/template) templates support.

## Quick Send

1. Create `.env` from `example.env`
2. Create your template in `templates` folder
3. Prepare `data.csv`:

```csv
EMAIL
user1@example.com
user2@example.com
```

3. Run:

```bash
go run ./cmd/quick/main.go --data-file="data.csv" --email-sender="sender@email.com" --subject="Subject" --tmpl-file="templates/example.html"
```

## Advanced Usage

### Running Without Flags

You can run program without any flags

```
go run ./cmd/quick/main.go
```

Required params will be taken from `data.csv` and `.env` files.

> [!TIP]
> Program supports flag `-h`
>
> ```sh
> go run ./cmd/quick/main.go -h  # Show all available options
> ```

### Data File

Data file (which is .csv) that defines email settings and template variables.

Advanced `data.csv` example:

```csv
EMAIL,TEMPLATE_FILE,SUBJECT,name
example@example.com,templates/example.html,My Example subject,Dias1c
example@example.com,templates/example.txt,Text letter,MyName
example@example.com,,,
```

Explaining variable keys:

| key             | variable | description                     |
| --------------- | -------- | ------------------------------- |
| `EMAIL`         | system   | variable, recipient email       |
| `TEMPLATE_FILE` | system   | variable, path to template file |
| `SUBJECT`       | system   | variable, letter subject        |
| `name`          | user     | variable                        |

> [!NOTE]
> For the last row with empty columns, the values ​​for those columns will be taken from the `program arguments` or `env` file.

#### System data Variables

| key             | type     | description           |
| --------------- | -------- | --------------------- |
| `EMAIL`         | required | recipient email       |
| `TEMPLATE_FILE` | optional | path to template file |
| `SUBJECT`       | optional | letter subject        |
| `SENDER_EMAIL`  | optional | aws sender email      |
| `SENDER_REGION` | optional | aws region            |

### Parameters Priority

Program uses priority system to resolve conflicts when same parameter is set in multiple places:

| Priority | Source        | Example                               |
| -------- | ------------- | ------------------------------------- |
| High     | Data file     | `data.csv`: SUBJECT=Welcome           |
| Medium   | Command flags | `--subject="Hello"` as flag on launch |
| Low      | .env file     | `SUBJECT=Hi` in .env                  |

### Templates

#### HTML

```html
<h1>Welcome {{.name}}!</h1>
<p>Your activation code {{.code}} for {{.EMAIL}}</p>
```

#### Text

```
--- {{.EMAIL}}
Dear {{.name}},
Your order #{{.orderId}} completed
```
