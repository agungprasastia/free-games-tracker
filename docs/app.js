/* ═══════════════════════════════════════════════════════════════
   Free Games Tracker — App Logic
   Fetches JSON data, renders cards, handles filter/search/sort
   ═══════════════════════════════════════════════════════════════ */

// ── State ──────────────────────────────────────────────────────
let allGames = [];
let allDeals = [];
let history = [];
let activeFilter = 'all';
let searchQuery = '';
let sortBy = 'newest';

// ── Data Fetching ──────────────────────────────────────────────

/**
 * Fetch JSON from multiple possible paths (handles local dev + GitHub Pages).
 * Tries 'data/<file>' first (GitHub Pages), then '../data/<file>' (local dev).
 */
async function fetchJSON(filename) {
  const paths = [`data/${filename}`, `../data/${filename}`];
  for (const path of paths) {
    try {
      const resp = await fetch(path);
      if (resp.ok) return await resp.json();
    } catch (_) { /* try next path */ }
  }
  console.warn(`Could not load ${filename}`);
  return [];
}

async function loadAllData() {
  [allGames, allDeals, history] = await Promise.all([
    fetchJSON('games.json'),
    fetchJSON('deals.json'),
    fetchJSON('history.json'),
  ]);
}

// ── Theme Toggle ───────────────────────────────────────────────

function initTheme() {
  const saved = localStorage.getItem('theme');
  const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
  const isDark = saved ? saved === 'dark' : prefersDark;

  document.documentElement.classList.toggle('dark', isDark);
  updateThemeIcon(isDark);

  document.getElementById('theme-toggle').addEventListener('click', () => {
    const nowDark = document.documentElement.classList.toggle('dark');
    localStorage.setItem('theme', nowDark ? 'dark' : 'light');
    updateThemeIcon(nowDark);
  });
}

function updateThemeIcon(isDark) {
  document.getElementById('sun-icon').classList.toggle('hidden', !isDark);
  document.getElementById('moon-icon').classList.toggle('hidden', isDark);
}

// ── Stats Rendering ────────────────────────────────────────────

function parseIDR(priceStr) {
  if (!priceStr || priceStr === 'Free' || priceStr === 'N/A' || priceStr === '0') return 0;
  const match = priceStr.match(/[\d,]+/);
  if (!match) return 0;
  return parseInt(match[0].replace(/,/g, ''), 10) || 0;
}

function formatIDR(n) {
  return n.toLocaleString('en-US');
}

function renderStats() {
  const totalGames = history.length;
  const totalValue = history.reduce((sum, g) => sum + parseIDRPrice(g.original_price), 0);

  const platformCounts = {};
  history.forEach((g) => {
    platformCounts[g.platform] = (platformCounts[g.platform] || 0) + 1;
  });

  const freeCount = allGames.length;
  const dealsCount = allDeals.length;

  // Update header counts
  document.getElementById('header-free-count').textContent = freeCount;
  document.getElementById('header-deals-count').textContent = dealsCount;

  // Stats cards
  const statsHtml = `
    <div class="stats-card rounded-2xl p-5 flex items-center gap-4">
      <div class="w-12 h-12 rounded-xl bg-accent-green/10 flex items-center justify-center">
        <i data-lucide="gift" class="w-6 h-6 text-accent-green"></i>
      </div>
      <div>
        <p class="text-2xl font-extrabold">${freeCount}</p>
        <p class="text-xs text-gray-500 dark:text-gray-400 font-medium">Free games now</p>
      </div>
    </div>
    <div class="stats-card rounded-2xl p-5 flex items-center gap-4">
      <div class="w-12 h-12 rounded-xl bg-accent-red/10 flex items-center justify-center">
        <i data-lucide="tag" class="w-6 h-6 text-accent-red"></i>
      </div>
      <div>
        <p class="text-2xl font-extrabold">${dealsCount}</p>
        <p class="text-xs text-gray-500 dark:text-gray-400 font-medium">Active deals</p>
      </div>
    </div>
    <div class="stats-card rounded-2xl p-5 flex items-center gap-4">
      <div class="w-12 h-12 rounded-xl bg-accent-purple/10 flex items-center justify-center">
        <i data-lucide="bar-chart-3" class="w-6 h-6 text-accent-purple"></i>
      </div>
      <div>
        <p class="text-2xl font-extrabold">${totalGames}</p>
        <p class="text-xs text-gray-500 dark:text-gray-400 font-medium">Games tracked</p>
      </div>
    </div>
    <div class="stats-card rounded-2xl p-5 flex items-center gap-4">
      <div class="w-12 h-12 rounded-xl bg-accent-blue/10 flex items-center justify-center">
        <i data-lucide="wallet" class="w-6 h-6 text-accent-blue"></i>
      </div>
      <div>
        <p class="text-2xl font-extrabold">IDR ${formatIDR(totalValue)}</p>
        <p class="text-xs text-gray-500 dark:text-gray-400 font-medium">Total value saved</p>
      </div>
    </div>
  `;
  document.getElementById('stats-section').innerHTML = statsHtml;
  refreshIcons();
}

// ── Countdown Timer ────────────────────────────────────────────

function getCountdown(endDateStr) {
  if (!endDateStr) return null;

  const end = new Date(endDateStr);
  if (isNaN(end.getTime())) return null;

  const now = new Date();
  const diff = end - now;

  if (diff <= 0) return { text: 'Expired', urgent: true, expired: true };

  const days = Math.floor(diff / 86400000);
  const hours = Math.floor((diff % 86400000) / 3600000);
  const mins = Math.floor((diff % 3600000) / 60000);

  let text;
  if (days > 0) text = `${days}d ${hours}h left`;
  else if (hours > 0) text = `${hours}h ${mins}m left`;
  else text = `${mins}m left`;

  return {
    text,
    urgent: days < 1,
    warning: days >= 1 && days <= 2,
    expired: false,
  };
}

// ── Card Rendering ─────────────────────────────────────────────

function platformBadge(platform) {
  if (platform === 'Epic Games') {
    return `<span class="badge-epic text-xs font-bold px-2.5 py-1 rounded-lg">Epic Games</span>`;
  }
  return `<span class="badge-steam text-xs font-bold px-2.5 py-1 rounded-lg">Steam</span>`;
}

function gameCardHTML(game, index) {
  const countdown = getCountdown(game.end_date);
  const countdownHtml = countdown
    ? `<div class="flex items-center gap-1.5 text-xs font-semibold ${countdown.urgent ? 'countdown-urgent' : countdown.warning ? 'countdown-warning' : 'text-gray-500 dark:text-gray-400'}">
         <i data-lucide="clock" class="w-3.5 h-3.5"></i>
         ${countdown.text}
       </div>`
    : '';

  const thumb = game.thumbnail || '';
  const thumbHtml = thumb
    ? `<img src="${thumb}" alt="${escapeHtml(game.title)}" loading="lazy" class="card-thumb w-full h-44 object-cover" onerror="this.parentElement.style.display='none'">`
    : `<div class="w-full h-44 flex items-center justify-center bg-gradient-to-br from-accent-purple/20 to-accent-blue/20"><i data-lucide="gamepad-2" class="w-12 h-12 text-accent-purple/50"></i></div>`;

  return `
    <a href="${game.url || '#'}" target="_blank" rel="noopener"
       class="game-card rounded-2xl flex flex-col card-enter" style="animation-delay: ${index * 0.05}s">
      <div class="relative overflow-hidden rounded-t-2xl">
        ${thumbHtml}
        <div class="absolute top-3 left-3">${platformBadge(game.platform)}</div>
        <div class="absolute top-3 right-3">
          <span class="tag-free text-xs font-extrabold px-3 py-1.5 rounded-lg">FREE</span>
        </div>
      </div>
      <div class="p-4 flex flex-col gap-2 flex-1">
        <h4 class="font-bold text-base leading-tight line-clamp-2" title="${escapeHtml(game.title)}">${escapeHtml(game.title)}</h4>
        <div class="flex items-center justify-between text-sm mt-auto">
          <span class="text-gray-400 line-through text-xs">${escapeHtml(game.original_price || '')}</span>
          ${countdownHtml}
        </div>
      </div>
    </a>
  `;
}

function dealCardHTML(deal, index) {
  const thumb = deal.thumbnail || '';
  const thumbHtml = thumb
    ? `<img src="${thumb}" alt="${escapeHtml(deal.title)}" loading="lazy" class="card-thumb w-full h-44 object-cover" onerror="this.parentElement.style.display='none'">`
    : `<div class="w-full h-44 flex items-center justify-center bg-gradient-to-br from-accent-red/20 to-accent-purple/20"><i data-lucide="tags" class="w-12 h-12 text-accent-red/50"></i></div>`;

  return `
    <a href="${deal.url || '#'}" target="_blank" rel="noopener"
       class="game-card rounded-2xl flex flex-col card-enter" style="animation-delay: ${index * 0.05}s">
      <div class="relative overflow-hidden rounded-t-2xl">
        ${thumbHtml}
        <div class="absolute top-3 left-3">${platformBadge(deal.platform)}</div>
        <div class="absolute top-3 right-3">
          <span class="tag-discount text-xs font-extrabold px-3 py-1.5 rounded-lg">-${deal.discount_percent}%</span>
        </div>
      </div>
      <div class="p-4 flex flex-col gap-2 flex-1">
        <h4 class="font-bold text-base leading-tight line-clamp-2" title="${escapeHtml(deal.title)}">${escapeHtml(deal.title)}</h4>
        <div class="flex items-center justify-between text-sm mt-auto">
          <span class="text-gray-400 line-through text-xs">${escapeHtml(deal.original_price || '')}</span>
          <span class="font-extrabold text-accent-green">${escapeHtml(deal.discounted_price || '')}</span>
        </div>
      </div>
    </a>
  `;
}

// ── Filtering & Sorting ────────────────────────────────────────

function getFilteredGames() {
  let games = [...allGames];

  if (activeFilter !== 'all') {
    games = games.filter((g) => g.platform === activeFilter);
  }
  if (searchQuery) {
    const q = searchQuery.toLowerCase();
    games = games.filter((g) => g.title.toLowerCase().includes(q));
  }

  sortItems(games, 'game');
  return games;
}

function getFilteredDeals() {
  let deals = [...allDeals];

  if (activeFilter !== 'all') {
    deals = deals.filter((d) => d.platform === activeFilter);
  }
  if (searchQuery) {
    const q = searchQuery.toLowerCase();
    deals = deals.filter((d) => d.title.toLowerCase().includes(q));
  }

  sortItems(deals, 'deal');
  return deals;
}

function sortItems(items, type) {
  switch (sortBy) {
    case 'newest':
      items.sort((a, b) => {
        const da = new Date(a.scraped_at || 0);
        const db = new Date(b.scraped_at || 0);
        return db - da;
      });
      break;
    case 'expiring':
      items.sort((a, b) => {
        const ea = new Date(a.end_date || '9999-12-31').getTime();
        const eb = new Date(b.end_date || '9999-12-31').getTime();
        return ea - eb;
      });
      break;
    case 'discount':
      if (type === 'deal') {
        items.sort((a, b) => b.discount_percent - a.discount_percent);
      }
      break;
    case 'title':
      items.sort((a, b) => a.title.localeCompare(b.title));
      break;
  }
}

// ── Render Grids ───────────────────────────────────────────────

function renderGrids() {
  const games = getFilteredGames();
  const deals = getFilteredDeals();

  // Free games grid
  const gamesGrid = document.getElementById('free-games-grid');
  const gamesCount = document.getElementById('free-games-count');

  if (games.length > 0) {
    gamesGrid.innerHTML = games.map((g, i) => gameCardHTML(g, i)).join('');
    gamesCount.textContent = `(${games.length})`;
  } else {
    gamesGrid.innerHTML = `
      <div class="col-span-full text-center py-12 text-gray-400">
        <div class="flex justify-center mb-2"><i data-lucide="search-x" class="w-10 h-10"></i></div>
        <p>No free games found${searchQuery ? ' for "' + escapeHtml(searchQuery) + '"' : ''}.</p>
      </div>`;
    gamesCount.textContent = '(0)';
  }

  // Deals grid
  const dealsGrid = document.getElementById('deals-grid');
  const dealsCount = document.getElementById('deals-count');

  if (deals.length > 0) {
    dealsGrid.innerHTML = deals.map((d, i) => dealCardHTML(d, i)).join('');
    dealsCount.textContent = `(${deals.length})`;
  } else {
    dealsGrid.innerHTML = `
      <div class="col-span-full text-center py-12 text-gray-400">
        <div class="flex justify-center mb-2"><i data-lucide="search-x" class="w-10 h-10"></i></div>
        <p>No deals found${searchQuery ? ' for "' + escapeHtml(searchQuery) + '"' : ''}.</p>
      </div>`;
    dealsCount.textContent = '(0)';
  }

  // Show/hide empty state
  const emptyState = document.getElementById('empty-state');
  const hasContent = games.length > 0 || deals.length > 0;
  emptyState.classList.toggle('hidden', hasContent);
  refreshIcons();
}

// ── Filter & Search Setup ──────────────────────────────────────

function initFilters() {
  // Platform filter buttons
  document.querySelectorAll('.filter-btn').forEach((btn) => {
    btn.addEventListener('click', () => {
      document.querySelectorAll('.filter-btn').forEach((b) => b.classList.remove('active'));
      btn.classList.add('active');
      activeFilter = btn.dataset.filter;
      renderGrids();
    });
  });

  // Search input (debounced)
  let searchTimeout;
  document.getElementById('search-input').addEventListener('input', (e) => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
      searchQuery = e.target.value.trim();
      renderGrids();
    }, 200);
  });

  // Sort dropdown
  document.getElementById('sort-select').addEventListener('change', (e) => {
    sortBy = e.target.value;
    renderGrids();
  });
}

// ── Live Countdown Updates ─────────────────────────────────────

function startCountdownUpdater() {
  setInterval(() => {
    // Only re-render if there are visible countdown timers
    const hasTimers = document.querySelectorAll('.countdown-urgent, .countdown-warning').length > 0;
    if (hasTimers) {
      renderGrids();
    }
  }, 60000); // Update every minute
}

// ── Utility ────────────────────────────────────────────────────

function escapeHtml(str) {
  if (!str) return '';
  const div = document.createElement('div');
  div.textContent = str;
  return div.innerHTML;
}

function refreshIcons() {
  if (window.lucide) lucide.createIcons();
}

function parseIDRPrice(priceStr) {
  if (!priceStr || priceStr === 'Free' || priceStr === 'N/A' || priceStr === '0') return 0;
  const match = priceStr.match(/[\d,]+/);
  if (!match) return 0;
  return parseInt(match[0].replace(/,/g, ''), 10) || 0;
}

// ── Last Updated ───────────────────────────────────────────────

function setLastUpdated() {
  let updatedStr = 'Recently';
  if (allGames.length > 0 && allGames[0].scraped_at) {
    const d = new Date(allGames[0].scraped_at);
    if (!isNaN(d.getTime())) {
      updatedStr = d.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        timeZone: 'UTC',
      }) + ' UTC';
    }
  }
  document.getElementById('header-updated').textContent = `Updated: ${updatedStr}`;
  document.getElementById('footer-updated').textContent = `Last updated: ${updatedStr}`;
}

// ── Init ───────────────────────────────────────────────────────

async function init() {
  initTheme();
  initFilters();
  refreshIcons(); // Render static HTML icons

  await loadAllData();

  // Clear loading skeletons
  document.getElementById('free-games-grid').innerHTML = '';
  document.getElementById('deals-grid').innerHTML = '';

  renderStats();
  renderGrids();
  setLastUpdated();
  startCountdownUpdater();
}

// Run on DOM ready
if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', init);
} else {
  init();
}
