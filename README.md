# Project X

Run

```sh
go run .
```

Install


```sh
go build
sudo cp x /usr/local/bin/x
```

## Archive reminders

```sh
x archive
```

Each archive is saved as a separate Markdown file in `~/Reminders`, named
with its local ISO 8601 datetime with second precision and no timezone suffix:
`YYYY-MM-DDTHH-mm-ss.md` (for example, `2026-10-15T05-18-45.md`).

Existing archive files are never appended to or overwritten. Reminders are deleted only after the new archive is written, verified, and committed.