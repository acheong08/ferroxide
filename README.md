# ferroxide

A third-party, open-source ProtonMail bridge. For power users only, designed to
run on a server.

ferroxide supports CardDAV, IMAP and SMTP.

Rationale:

* No GUI, only a CLI (so it runs in headless environments)
* Standard-compliant (we don't care about Microsoft Outlook)
* Fully open-source

Feel free to join the IRC channel: #emersion on Libera Chat.

## How does it work?

ferroxide is a server that translates standard protocols (SMTP, IMAP, CardDAV)
into ProtonMail API requests. It allows you to use your preferred e-mail clients
and `git-send-email` with ProtonMail.

    +-----------------+             +-------------+  ProtonMail  +--------------+
    |                 | IMAP, SMTP  |             |     API      |              |
    |  E-mail client  <------------->  ferroxide  <-------------->  ProtonMail  |
    |                 |             |             |              |              |
    +-----------------+             +-------------+              +--------------+

## Setup

### Go

ferroxide is implemented in Go. Head to [Go website](https://golang.org) for
setup information.

### Installing

Start by installing ferroxide:

```shell
git clone https://github.com/acheong08/ferroxide.git
go build ./cmd/ferroxide
```

Then you'll need to login to ProtonMail via ferroxide, so that ferroxide can
retrieve e-mails from ProtonMail. You can do so with this command:

```shell
ferroxide auth <username>
```

Once you're logged in, a "bridge password" will be printed. Don't close your
terminal yet, as this password is not stored anywhere by ferroxide and will be
needed when configuring your e-mail client.

Your ProtonMail credentials are stored on disk encrypted with this bridge
password (a 32-byte random password generated when logging in).

## Usage

ferroxide can be used in multiple modes.

> Don't start ferroxide multiple times, instead you can use `ferroxide serve`.
> This requires ports 1025 (smtp), 1143 (imap), and 8080 (carddav).

### SMTP

To run ferroxide as an SMTP server:

```shell
ferroxide smtp
```

Once the bridge is started, you can configure your e-mail client with the
following settings:

* Hostname: `localhost`
* Port: 1025
* Security: none
* Username: your ProtonMail username
* Password: the bridge password (not your ProtonMail password)

### CardDAV

You must setup an HTTPS reverse proxy to forward requests to `ferroxide`.

```shell
ferroxide carddav
```

Tested on GNOME (Evolution) and Android (DAVDroid).

### IMAP

⚠️  **Warning**: IMAP support is work-in-progress. Here be dragons.

For now, it only supports unencrypted local connections.

```shell
ferroxide imap
```

## License

MIT
