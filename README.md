# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-09-04 06:56 UTC_

> 📊 **21** games tracked · **IDR 3,316,668** total value saved · Epic Games: 21

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Alone With You | Epic Games | IDR 69,999 | Sep 10, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/alone-with-you-028a15) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Grand Theft Auto V Enhanced | **-56%** | ~~IDR 659,000~~ | **IDR 289,960** | [View](https://store.steampowered.com/app/3240220/) |
| Call of Duty®: Black Ops III | **-75%** | ~~IDR 891,000~~ | **IDR 222,750** | [View](https://store.steampowered.com/app/311210/) |
| Atomic Heart | **-80%** | ~~IDR 549,000~~ | **IDR 109,800** | [View](https://store.steampowered.com/app/668580/) |
| Warhammer 40,000: Space Marine 2 | **-75%** | ~~IDR 609,000~~ | **IDR 152,250** | [View](https://store.steampowered.com/app/2183900/) |
| Call of Duty®: Black Ops III | **-75%** | ~~IDR 891,000~~ | **IDR 222,750** | [View](https://store.steampowered.com/app/311210/) |
| Sons Of The Forest | **-70%** | ~~IDR 245,999~~ | **IDR 73,799** | [View](https://store.steampowered.com/app/1326470/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
