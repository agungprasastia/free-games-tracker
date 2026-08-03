# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-08-03 05:40 UTC_

> 📊 **13** games tracked · **IDR 2,319,672** total value saved · Epic Games: 13

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| OTXO | Epic Games | IDR 103,999 | Aug 06, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/otxo-396b8b) |
| Sol Cesto | Epic Games | IDR 122,999 | Aug 06, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/sol-cesto-e9b803) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Marvel’s Spider-Man Remastered | **-60%** | ~~IDR 879,000~~ | **IDR 351,600** | [View](https://store.steampowered.com/app/1817070/) |
| Marvel’s Spider-Man: Miles Morales | **-60%** | ~~IDR 729,000~~ | **IDR 291,600** | [View](https://store.steampowered.com/app/1817190/) |
| Cyberpunk 2077 | **-70%** | ~~IDR 699,999~~ | **IDR 209,999** | [View](https://store.steampowered.com/app/1091500/) |
| Squad | **-60%** | ~~IDR 437,591~~ | **IDR 175,036** | [View](https://store.steampowered.com/app/393380/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
