# sigil

Local-first, encrypted secrets for developers. No cloud. No daemon. No account.

Store credentials locally and inject them directly into any process — without touching a `.env` file or opening a password manager.

```bash
sigil set DB_URL "postgres://localhost/myapp"
```

## Why

`.env` files get committed. Password managers break your flow. Cloud secret managers are overkill for local development.

`sigil` stores your secrets encrypted on disk and injects them as environment variables when you need them, then gets out of the way.

## Install

### Homebrew

```bash
brew install yourname/tap/sigil
```

### Go

```bash
go install github.com/yourname/sigil/cmd/sigil@latest
```

### Download

Grab a binary from the [releases page](https://github.com/yourname/sigil/releases).

## Usage

### Store a secret

```bash
sigil set KEY VALUE
sigil set DB_URL "postgres://localhost/myapp"
sigil set STRIPE_KEY "sk_test_..."
```

### Retrieve a secret

```bash
sigil get KEY
sigil get DB_URL
```

### List all keys

```bash
sigil list
```

### Delete a secret

```bash
sigil delete KEY
```

### Inject secrets into a process

```bash
sigil exec -- <command> [args...]

sigil exec -- go run main.go
sigil exec -- npm run dev
sigil exec -- python app.py
```

All stored secrets are injected as environment variables. The process receives them directly — nothing is written to disk or exposed in your shell history.

### Projects

Secrets are namespaced by project. The default project is `default`.

```bash
sigil set DB_URL "postgres://..." -project myapp
sigil exec -project myapp -- go run main.go
```

## How it works

Secrets are stored in `~/.local/share/sigil/<project>/store.enc` (XDG-compliant). Each store is encrypted with AES-256-GCM, keyed from your passphrase via PBKDF2. The file on disk is never readable as plain text.

When you run `sigil exec`, your secrets are decrypted in memory and passed directly to the target process via `execve`. `sigil` replaces itself with the target process — it does not stay running as a parent.

## Security

- Encryption: AES-256-GCM (authenticated)
- Key derivation: PBKDF2-SHA256
- No secrets are ever written to disk in plaintext
- No secrets appear in shell history
- `sigil exec` uses `execve` — no parent process retains secrets after exec

## Requirements

- Go 1.21+
- macOS, Linux, or Windows

## License

MIT
