# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-07-15 04:51 UTC_

> 📊 **8** games tracked · **IDR 1,454,994** total value saved · Epic Games: 8

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Nova Lands | Epic Games | IDR 165,999 | Jul 16, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/nova-lands-4d1788) |
| Tattoo Tycoon | Epic Games | IDR 269,000 | Jul 16, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/tattoo-tycoon-b4352c) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Grand Theft Auto V Enhanced | **-50%** | ~~IDR 439,000~~ | **IDR 219,500** | [View](https://store.steampowered.com/app/3240220/) |
| Red Dead Redemption 2 | **-75%** | ~~IDR 879,000~~ | **IDR 219,750** | [View](https://store.steampowered.com/app/1174180/) |
| The Outlast Trials | **-70%** | ~~IDR 299,999~~ | **IDR 89,999** | [View](https://store.steampowered.com/app/1304930/) |
| No Man's Sky | **-60%** | ~~IDR 449,999~~ | **IDR 179,999** | [View](https://store.steampowered.com/app/275850/) |
| Assassin’s Creed Shadows | **-55%** | ~~IDR 799,000~~ | **IDR 359,550** | [View](https://store.steampowered.com/app/3159330/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
