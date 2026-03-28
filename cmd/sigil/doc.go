// Package main is the entry point for sigil, a local-first encrypted
// secrets manager for developers.
//
// Secrets are stored encrypted on disk under the XDG data directory.
// No cloud, no daemon, no account required.
//
// Usage:
//
//	sigil set    [-project] KEY
//	sigil get    [-project] [-clip] [-clear 15] KEY
//	sigil list   [-project]
//	sigil delete [-project] KEY
//
// set prompts for both the passphrase and the secret value interactively —
// nothing sensitive is passed as a command-line argument or written to
// shell history.
//
// get copies the secret value directly to the clipboard and clears it
// after a configurable timeout. The value is never printed to the terminal.
//
// Secrets are encrypted with AES-256-GCM, keyed from a passphrase via
// PBKDF2-SHA256. The store file on disk is never readable as plain text.
package main
