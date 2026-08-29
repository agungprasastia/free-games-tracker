# 🎮 Free Games Tracker

Automatically tracks free games from **Epic Games** & **Steam** — updated daily via GitHub Actions.

_Last updated: 2026-08-29 08:38 UTC_

> 📊 **20** games tracked · **IDR 3,246,669** total value saved · Epic Games: 20

## 🔥 Current free games

| Game | Platform | Normal Price | Available Until | Link |
|------|----------|-------------|----------------|------|
| Breathedge | Epic Games | IDR 199,000 | Sep 03, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/breathedge) |
| Rival Stars Horse Racing : Desktop Edition | Epic Games | IDR 172,999 | Sep 03, 2026 15:00 UTC | [Claim](https://store.epicgames.com/en-US/p/rival-stars-horse-racing-dd09de) |

## 🏷️ Steam deals (>50% off)

| Game | Discount | Original | Sale Price | Link |
|------|----------|----------|------------|------|
| Grand Theft Auto V Enhanced | **-56%** | ~~IDR 659,000~~ | **IDR 289,960** | [View](https://store.steampowered.com/app/3240220/) |
| Red Dead Redemption 2 | **-75%** | ~~IDR 879,000~~ | **IDR 219,750** | [View](https://store.steampowered.com/app/1174180/) |
| Forza Horizon 5 | **-60%** | ~~IDR 699,000~~ | **IDR 279,600** | [View](https://store.steampowered.com/app/1551360/) |
| Hogwarts Legacy | **-90%** | ~~IDR 799,000~~ | **IDR 79,900** | [View](https://store.steampowered.com/app/990080/) |

## 📦 Data

- [`data/games.json`](data/games.json) — current free games
- [`data/deals.json`](data/deals.json) — current Steam deals (>50% off)
- [`data/history.json`](data/history.json) — all free games ever tracked

## 🤖 How it works

GitHub Actions runs every day at 09:00 WIB, scrapes Epic Games & Steam APIs, updates the data files, and commits the changes automatically.

Built with **Go** 🐹 for simplicity and performance.
