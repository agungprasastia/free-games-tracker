# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-07-07 06:07 UTC_

> 📊 **6** games tracked · **IDR 1,019,995** total value saved · Epic Games: 6

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| River City Girls 2 | Epic Games | IDR 276,999 | Jul 09, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/river-city-girls-2-77af3a) |
| I Have No Mouth, and I Must Scream | Epic Games | IDR 69,999 | Jul 09, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/i-have-no-mouth-and-i-must-scream-95c5c2) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Cyberpunk 2077 | **-70%** | ~~IDR 699,999~~ | **IDR 209,999** | [View](https://store.steampowered.com/app/1091500/) |
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Red Dead Redemption 2 | **-75%** | ~~IDR 879,000~~ | **IDR 219,750** | [View](https://store.steampowered.com/app/1174180/) |
| Sons Of The Forest | **-70%** | ~~IDR 245,999~~ | **IDR 73,799** | [View](https://store.steampowered.com/app/1326470/) |
| Sekiro™: Shadows Die Twice - GOTY Edition | **-50%** | ~~IDR 891,000~~ | **IDR 445,500** | [View](https://store.steampowered.com/app/814380/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
