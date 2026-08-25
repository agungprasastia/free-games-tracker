# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-08-25 02:56 UTC_

> 📊 **18** games tracked · **IDR 2,874,670** total value saved · Epic Games: 18

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Cardpocalypse Standard Edition | Epic Games | IDR 187,000 | Aug 27, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/cardpocalypse) |
| Epic Mage Bundle | Epic Games | 0 | Aug 27, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/albion-online-epic-mage-bundle-2ceb19) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Grand Theft Auto V Enhanced | **-56%** | ~~IDR 659,000~~ | **IDR 289,960** | [View](https://store.steampowered.com/app/3240220/) |
| Ryse: Son of Rome | **-71%** | ~~IDR 116,000~~ | **IDR 33,640** | [View](https://store.steampowered.com/app/302510/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
