// Package main is the entry point for sigil, a local-first encrypted
// secrets manager for developers.
//
// Secrets are stored encrypted on disk under the XDG data directory and
// injected as environment variables directly into processes via execve.
// No cloud, no daemon, no account required.
//
// Usage:
//
//	sigil set    [-project] KEY VALUE
//	sigil get    [-project] KEY
//	sigil list   [-project]
//	sigil delete [-project] KEY
//	sigil clip   [-project] [-clear 15]KEY
//	sigil exec   [-project] -- command [args...]
//
// Secrets are encrypted with AES-256-GCM, keyed from a passphrase via
// PBKDF2-SHA256. The store file on disk is never readable as plain text.
package main
