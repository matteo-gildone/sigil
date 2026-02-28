// Package main is the entry point for sigil, a local-first encrypted
// secrets manager for developers.
//
// Secrets are stored encrypted on disk under the XDG data directory and
// injected as environment variables directly into processes via execve.
// No cloud, no daemon, no account required.
//
// Usage:
//
//	sigil set KEY VALUE [-project name]
//	sigil get KEY [-project name]
//	sigil list [-project name]
//	sigil delete KEY [-project name]
//	sigil exec [-project name] -- command [args...]
//
// Secrets are encrypted with AES-256-GCM, keyed from a passphrase via
// PBKDF2-SHA256. The store file on disk is never readable as plain text.
package main
