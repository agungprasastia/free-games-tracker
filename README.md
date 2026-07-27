# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-07-27 05:49 UTC_

> 📊 **11** games tracked · **IDR 2,092,674** total value saved · Epic Games: 11

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Foretales | Epic Games | IDR 298,682 | Jul 30, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/foretales-d6c5bd) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| METAL GEAR SOLID Δ: SNAKE EATER | **-60%** | ~~IDR 904,000~~ | **IDR 361,600** | [View](https://store.steampowered.com/app/2417610/) |
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| ICARUS | **-80%** | ~~IDR 284,999~~ | **IDR 56,999** | [View](https://store.steampowered.com/app/1149460/) |
| METAL GEAR SOLID: MASTER COLLECTION Vol.1 | **-50%** | ~~IDR 729,000~~ | **IDR 364,500** | [View](https://store.steampowered.com/app/886313/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
