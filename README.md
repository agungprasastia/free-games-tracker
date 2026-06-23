# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-06-23 06:06 UTC_

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
| Horizon Forbidden West™ Complete Edition | **-50%** | ~~IDR 879,000~~ | **IDR 439,500** | [View](https://store.steampowered.com/app/2420110/) |
| Don't Starve Together | **-90%** | ~~IDR 95,999~~ | **IDR 9,599** | [View](https://store.steampowered.com/app/322330/) |
| Dead by Daylight | **-60%** | ~~IDR 219,890~~ | **IDR 87,956** | [View](https://store.steampowered.com/app/381210/) |
| The Witcher 3: Wild Hunt - Complete Edition | **-80%** | ~~IDR 449,000~~ | **IDR 89,800** | [View](https://store.steampowered.com/app/124923/) |
| It Takes Two | **-70%** | ~~IDR 479,000~~ | **IDR 143,700** | [View](https://store.steampowered.com/app/1426210/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
