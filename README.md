# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-08-17 02:57 UTC_

> 📊 **16** games tracked · **IDR 2,687,670** total value saved · Epic Games: 16

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Caravan SandWitch | Epic Games | IDR 140,000 | Aug 20, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/caravan-sandwitch-05ff58) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| PEAK | **-50%** | ~~IDR 69,999~~ | **IDR 34,999** | [View](https://store.steampowered.com/app/3527290/) |
| KINGDOM HEARTS -HD 1.5+2.5 ReMIX- | **-70%** | ~~IDR 569,000~~ | **IDR 170,700** | [View](https://store.steampowered.com/app/2552430/) |
| Cult of the Lamb | **-60%** | ~~IDR 149,999~~ | **IDR 59,999** | [View](https://store.steampowered.com/app/1313140/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
