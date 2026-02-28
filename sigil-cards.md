# Sigil — Project Cards

---

## Card 01 — Project scaffold

Set up the module and directory structure. No logic yet, just the skeleton.

**Tasks**
- `go mod init github.com/yourname/sigil`
- Create the directory layout:
  ```
  sigil/
    cmd/sigil/
      main.go
    internal/
      cli/
      crypto/
      store/
      xdg/
    go.mod
    Makefile
    README.md
  ```
- Write a `main.go` that prints `"sigil"` and exits cleanly
- Write a `Makefile` with a `build` target

**Done when:** 
- [ ] `go build ./...` passes with no errors.

---

## Card 02 — CLI dispatcher

Build the command routing layer. No subcommand logic yet — just the skeleton that dispatches correctly.

**Tasks**
- Read `os.Args` and dispatch to `set`, `get`, `list`, `delete`, `exec`
- Each command gets its own `flag.FlagSet` with a `-project` flag defaulting to `"default"`
- Print a usage message for unknown commands
- Print per-command usage when called with no args

**Done when:** 
- [ ] `sigil set`, `sigil get`, `sigil exec` etc. each print their own placeholder message. `sigil foo` prints an error.

---

## Card 03 — XDG paths

Implement the XDG base directory logic so sigil knows where to read and write files.

**Tasks**
- Implement `DataDir() string` — returns `$XDG_DATA_HOME/sigil` or `~/.local/share/sigil` as fallback
- Implement `ConfigDir() string` — returns `$XDG_CONFIG_HOME/sigil` or `~/.config/sigil` as fallback
- Implement `ProjectPath(project string) string` — returns the full path to a project's store file
- Create directories if they don't exist (`os.MkdirAll`)
- Write table-driven tests for all three functions

**Done when:** 
- [ ] Tests pass. Calling `ProjectPath("default")` returns the correct path on your machine.

---

## Card 04 — Passphrase input

Implement secure passphrase reading from the terminal.

**Tasks**
- Add `golang.org/x/term` to the module
- Implement `PromptPassphrase(prompt string) ([]byte, error)` in `internal/cli`
- Print the prompt, read without echo, print a newline after
- Return an error if stdin is not a terminal

**Done when:** 
- [ ] Running a test binary prompts for a passphrase and the input is not visible.

---

## Card 05 — Encryption

Implement encrypt and decrypt using the stdlib. This is the hardest card — take your time.

**Tasks**
- Implement `Encrypt(passphrase, plaintext []byte) ([]byte, error)` in `internal/crypto`
  - Generate a random 16-byte salt
  - Derive a 32-byte key using `crypto/pbkdf2` with SHA-256
  - Encrypt with `crypto/aes` + `crypto/cipher` in GCM mode
  - Return `salt + nonce + ciphertext` concatenated
- Implement `Decrypt(passphrase, data []byte) ([]byte, error)`
  - Split the input back into salt, nonce, ciphertext
  - Re-derive the key, decrypt and authenticate
  - Return a clear error if the passphrase is wrong
- Write table-driven tests:
  - Encrypt then decrypt returns original plaintext
  - Wrong passphrase returns an error
  - Two encryptions of the same plaintext produce different ciphertexts (random nonce)

**Done when:** 
- [ ] All tests pass.

---

## Card 06 — Store

Implement the store: load, save, get, set, delete, list. No encryption wired in yet — work with plaintext JSON first.

**Tasks**
- Define a `Store` type backed by `map[string]string`
- Implement `Load(path string) (*Store, error)` — reads and unmarshals JSON
- Implement `Save(path string) error` — marshals to JSON and writes atomically via `os.CreateTemp` + `os.Rename`
- Implement `Set(key, value string)`, `Get(key string) (string, bool)`, `Delete(key string)`, `List() []string`
- Write table-driven tests for all operations
- Test that `Save` followed by `Load` round-trips correctly

**Done when:** 
- [ ] All tests pass. You can save a store to disk and load it back.

---

## Card 07 — Wire encryption into the store

Connect the crypto and store layers so the file on disk is always encrypted.

**Tasks**
- Update `Load` to accept a passphrase, read the raw bytes, decrypt them, then unmarshal
- Update `Save` to marshal to JSON, encrypt, then write the ciphertext atomically
- Update all callers to pass the passphrase through
- Test that the file on disk is not readable as plain JSON
- Test that loading with a wrong passphrase returns a clear error

**Done when:** 
- [ ] The store file on disk is encrypted. Tests pass.

---

## Card 08 — `set` and `get` commands

Wire the first two real commands end to end.

**Tasks**
- `sigil set KEY VALUE` — prompts for passphrase, loads store, sets key, saves
- `sigil get KEY` — prompts for passphrase, loads store, prints value or exits with error if not found
- Handle the case where the store file doesn't exist yet (`set` should create it)
- Print a clear error if the wrong passphrase is given

**Done when:** 
- [ ] You can `sigil set DB_URL "postgres://..."` and then `sigil get DB_URL` and get the value back.

---

## Card 09 — `list` and `delete` commands

**Tasks**
- `sigil list` — prompts for passphrase, prints all keys, one per line
- `sigil list -project myapp` — lists keys for a specific project
- `sigil delete KEY` — prompts for passphrase, removes key, saves, confirms deletion
- Print a clear error if the key doesn't exist

**Done when:** 
- [ ] All four commands work end to end.

---

## Card 10 — `exec` command (Unix)

Implement secret injection via `syscall.Exec`.

**Tasks**
- `sigil exec -- <command> [args...]` — prompts for passphrase, loads secrets, execs into the command
- Use `exec.LookPath` to resolve the binary
- Merge store keys into `os.Environ()` and pass to `syscall.Exec`
- Create `internal/cli/exec_unix.go` with build tag `//go:build !windows`
- Create `internal/cli/exec_windows.go` with build tag `//go:build windows` using `exec.Command` as fallback
- Test that the target process receives the injected env vars

**Done when:** 
- [ ] `sigil exec -- env | grep DB_URL` prints the stored value.

---

## Card 11 — Error handling audit

Go through every command and make sure errors are handled correctly and consistently.

**Tasks**
- All errors print to `os.Stderr`, not `os.Stdout`
- All error paths exit with `os.Exit(1)`
- No naked `_` ignoring errors anywhere
- Error messages are lowercase, no punctuation (Go convention)
- Wrong passphrase gives a clear message, not a raw crypto error
- Missing key gives a clear message

**Done when:** 
- [ ] `go vet ./...` passes clean. Every error path has been manually tested.

---

## Card 12 — Distribution

Make sigil installable and releasable.

**Tasks**
- Add a `Makefile` target that builds for `linux/amd64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`
- Write a `README.md` with installation instructions and usage examples for all commands
- Tag `v0.1.0`
- Set up a GitHub Actions workflow that builds and attaches binaries on tag push

**Done when:** 
- [ ] `v0.1.0` is tagged, binaries are attached to the GitHub release, README is complete.

---
