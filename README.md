# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-08-09 03:40 UTC_

> 📊 **15** games tracked · **IDR 2,547,670** total value saved · Epic Games: 15

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Beacon Pines | Epic Games | IDR 137,999 | Aug 13, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/beacon-pines-629fc3) |
| We Were Here Together | Epic Games | IDR 89,999 | Aug 13, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/we-were-here-together-6a6d66) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Cyberpunk 2077 | **-70%** | ~~IDR 699,999~~ | **IDR 209,999** | [View](https://store.steampowered.com/app/1091500/) |
| Tom Clancy's Ghost Recon® Wildlands | **-95%** | ~~IDR 515,000~~ | **IDR 25,750** | [View](https://store.steampowered.com/app/460930/) |
| Ready or Not | **-50%** | ~~IDR 255,999~~ | **IDR 127,999** | [View](https://store.steampowered.com/app/1144200/) |
| Tom Clancy's Ghost Recon® Wildlands | **-95%** | ~~IDR 515,000~~ | **IDR 25,750** | [View](https://store.steampowered.com/app/460930/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
