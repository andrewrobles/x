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
x reminders 
```

Each archive is saved as a separate Markdown file in `~/Reminders`, named
with its local ISO 8601 datetime with second precision and no timezone suffix:
`YYYY-MM-DDTHH-mm-ss.md` (for example, `2026-10-15T05-18-45.md`).

Existing archive files are never appended to or overwritten. Reminders are deleted only after the new archive is written, verified, and committed.

## Archive notepad

    x notepad

Reads `~/Notepad/.md` and archives its exact contents to
`~/Notepad/YYYY-MM-DDTHH-mm-ss.md`, using local time with second precision
and no timezone suffix.

The archive is written, verified, and committed to the Git repository in
`~/Notepad` with an empty commit message. The source file is then emptied,
not deleted. Missing source files produce an error; empty files are skipped.
Existing archives are never overwritten.

Like the reminders command, this stages all changes in the archive directory.
The source is cleared after the commit, so its clearing remains an uncommitted
change until the next commit.

Do not edit the source file while archiving.