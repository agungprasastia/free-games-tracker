# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-07-12 05:27 UTC_

> 📊 **8** games tracked · **IDR 1,454,994** total value saved · Epic Games: 8

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Nova Lands | Epic Games | IDR 165,999 | Jul 16, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/nova-lands-4d1788) |
| Tattoo Tycoon | Epic Games | IDR 269,000 | Jul 16, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/tattoo-tycoon-b4352c) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Grand Theft Auto V Enhanced | **-56%** | ~~IDR 659,000~~ | **IDR 289,960** | [View](https://store.steampowered.com/app/3240220/) |
| Red Dead Redemption 2 | **-75%** | ~~IDR 879,000~~ | **IDR 219,750** | [View](https://store.steampowered.com/app/1174180/) |
| Planet Zoo | **-95%** | ~~IDR 440,278~~ | **IDR 22,013** | [View](https://store.steampowered.com/app/703080/) |
| Kingdom Come: Deliverance II | **-60%** | ~~IDR 641,000~~ | **IDR 256,400** | [View](https://store.steampowered.com/app/1771300/) |
| Grand Theft Auto: The Trilogy – The Definitive Edition | **-67%** | ~~IDR 649,000~~ | **IDR 214,170** | [View](https://store.steampowered.com/app/817628/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
