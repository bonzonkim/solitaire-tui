# Quickstart: Card Rendering Implementation

**Date**: 2026-01-18  
**Feature**: 007-card-rendering

## Prerequisites

- Go 1.23+
- Terminal with Unicode support (monospace font)

## Build & Run

```bash
cd /Users/bumgukang/bonzonkim/github.com/solitaire-tui

# Build
go build -o solitaire ./cmd/solitaire

# Run
./solitaire
```

## Verification

### Run Tests
```bash
go test ./...
```

### Manual Visual Check
1. Launch the game: `./solitaire`
2. Verify cards display as 11×7 boxes with Unicode borders
3. Check face-up cards show rank in corners and suit centered
4. Check face-down cards show ░ pattern fill
5. Verify "10" rank doesn't break alignment

## Key Files

| File | Purpose |
|------|---------|
| `internal/ui/view.go` | Card rendering logic (modified) |
| `internal/ui/styles/styles.go` | Card dimensions and colors |
| `internal/game/card.go` | Card model (unchanged) |
