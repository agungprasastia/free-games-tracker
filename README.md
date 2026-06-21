# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-06-21 00:17 UTC_

> 📊 **2** games tracked · **IDR 246,998** total value saved · Epic Games: 2

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Citizen Sleeper | Epic Games | IDR 137,999 | Jun 25, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/citizen-sleeper-944858) |
| ROBOBEAT | Epic Games | IDR 108,999 | Jun 25, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/robobeat-5f084b) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Cyberpunk 2077 | **-70%** | ~~IDR 699,999~~ | **IDR 209,999** | [View](https://store.steampowered.com/app/1091500/) |
| The Witcher 3: Wild Hunt - Complete Edition | **-80%** | ~~IDR 449,000~~ | **IDR 89,800** | [View](https://store.steampowered.com/app/124923/) |
| Dead by Daylight | **-60%** | ~~IDR 219,890~~ | **IDR 87,956** | [View](https://store.steampowered.com/app/381210/) |
| Dead Space | **-90%** | ~~IDR 659,000~~ | **IDR 65,900** | [View](https://store.steampowered.com/app/1693980/) |
| DayZ | **-55%** | ~~IDR 649,999~~ | **IDR 292,499** | [View](https://store.steampowered.com/app/221100/) |
| FINAL FANTASY VII REBIRTH | **-60%** | ~~IDR 729,000~~ | **IDR 291,600** | [View](https://store.steampowered.com/app/2909400/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
