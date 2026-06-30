# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-06-30 06:11 UTC_

> 📊 **4** games tracked · **IDR 672,997** total value saved · Epic Games: 4

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| RollerCoaster Tycoon 3 Complete Edition | Epic Games | IDR 288,000 | Jul 02, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/rollercoaster-tycoon-3-complete-edition) |
| Voidwrought | Epic Games | IDR 137,999 | Jul 02, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/voidwrought-ce8f4b) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Cyberpunk 2077 | **-70%** | ~~IDR 699,999~~ | **IDR 209,999** | [View](https://store.steampowered.com/app/1091500/) |
| Red Dead Redemption 2 | **-75%** | ~~IDR 879,000~~ | **IDR 219,750** | [View](https://store.steampowered.com/app/1174180/) |
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Sons Of The Forest | **-70%** | ~~IDR 245,999~~ | **IDR 73,799** | [View](https://store.steampowered.com/app/1326470/) |
| Ready or Not | **-50%** | ~~IDR 255,999~~ | **IDR 127,999** | [View](https://store.steampowered.com/app/1144200/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
