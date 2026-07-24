# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-07-24 05:19 UTC_

> 📊 **11** games tracked · **IDR 2,092,674** total value saved · Epic Games: 11

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Foretales | Epic Games | IDR 298,682 | Jul 30, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/foretales-d6c5bd) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| ICARUS | **-80%** | ~~IDR 284,999~~ | **IDR 56,999** | [View](https://store.steampowered.com/app/1149460/) |
| METAL GEAR SOLID: MASTER COLLECTION Vol.1 | **-50%** | ~~IDR 729,000~~ | **IDR 364,500** | [View](https://store.steampowered.com/app/886313/) |
| Warhammer 40,000: Space Marine 2 | **-70%** | ~~IDR 549,000~~ | **IDR 164,700** | [View](https://store.steampowered.com/app/2183900/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
