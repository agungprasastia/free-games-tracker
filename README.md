# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-07-18 04:48 UTC_

> 📊 **10** games tracked · **IDR 1,793,992** total value saved · Epic Games: 10

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Luto | Epic Games | IDR 165,999 | Jul 23, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/luto-0a4ab3) |
| Echo Generation: Midnight Edition | Epic Games | IDR 172,999 | Jul 23, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/echo-generation-midnight-edition-069026) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Red Dead Redemption 2 | **-75%** | ~~IDR 879,000~~ | **IDR 219,750** | [View](https://store.steampowered.com/app/1174180/) |
| Warhammer 40,000: Space Marine 2 | **-70%** | ~~IDR 549,000~~ | **IDR 164,700** | [View](https://store.steampowered.com/app/2183900/) |
| The Outlast Trials | **-70%** | ~~IDR 299,999~~ | **IDR 89,999** | [View](https://store.steampowered.com/app/1304930/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
