"use strict";

/**
 * Tab management for multi-team support.
 * Each tab corresponds to a team. Team data is stored as parsed JSON.
 */

const Tabs = (function () {
  const teams = new Map(); // id -> team JSON object
  let activeTeamId = null;

  function getTabBar() {
    return document.getElementById("tab-bar");
  }

  function getMainContent() {
    return document.getElementById("main-content");
  }

  function createTeamTab() {
    if (typeof window.createTeam !== "function") {
      console.error("WASM not loaded yet");
      return;
    }

    const teamJSON = window.createTeam();
    const team = JSON.parse(teamJSON);
    teams.set(team.id, team);

    const tabBar = getTabBar();
    const btn = document.createElement("button");
    btn.className = "tab";
    btn.dataset.teamId = team.id;
    btn.textContent = team.name || "New Team";
    btn.addEventListener("click", function () {
      activateTab(team.id);
    });

    // Insert before the + button
    const newBtn = document.getElementById("tab-new");
    tabBar.insertBefore(btn, newBtn);

    activateTab(team.id);
    console.log("Created team:", teamJSON);
  }

  function activateTab(teamId) {
    if (!teams.has(teamId)) return;

    activeTeamId = teamId;

    // Update tab styles
    const tabs = getTabBar().querySelectorAll(".tab");
    tabs.forEach(function (t) {
      t.classList.toggle("active", t.dataset.teamId === teamId);
    });

    renderActiveTeam();
  }

  function renderActiveTeam() {
    const content = getMainContent();
    const team = teams.get(activeTeamId);
    if (!team) {
      content.innerHTML = '<div class="empty-state"><h2>No team selected</h2></div>';
      return;
    }

    // Placeholder rendering — will be replaced by team-ui.js
    content.innerHTML =
      '<div class="team-header">' +
      '<div style="color:var(--text-muted);font-size:12px;letter-spacing:2px">TEAM: ' +
      (team.name || "NEW TEAM") +
      "</div>" +
      '<div style="color:var(--text-muted);font-size:11px;margin-top:4px">Vehicles: ' +
      (team.vehicles ? team.vehicles.length : 0) +
      "</div>" +
      "</div>";

    // Update budget bar
    updateBudgetBar(team);
  }

  function updateBudgetBar(team) {
    const bar = document.getElementById("budget-bar");
    if (!team) {
      bar.style.display = "none";
      return;
    }
    bar.style.display = "flex";

    // Calculate spent from WASM if available
    let spent = 0;
    let totalHull = 0;
    if (typeof window.validateTeam === "function") {
      // Re-serialize and get fresh data
      const result = window.updateTeam(JSON.stringify(team));
      const parsed = JSON.parse(result);
      if (parsed.team) {
        // Use the updated team
        teams.set(team.id, parsed.team);
      }
    }

    // Simple client-side calculation as fallback
    if (team.vehicles) {
      team.vehicles.forEach(function (v) {
        // We'll rely on WASM for accurate cost later
        spent = 0; // placeholder
      });
    }

    const remaining = team.budget - spent;
    const vehicleCount = team.vehicles ? team.vehicles.length : 0;

    bar.innerHTML =
      '<div class="stat-pill"><span class="label">Budget</span><span class="value">' +
      team.budget +
      "</span></div>" +
      '<div class="stat-pill"><span class="label">Spent</span><span class="value">' +
      spent +
      "</span></div>" +
      '<div class="stat-pill' +
      (remaining < 0 ? " over" : "") +
      '"><span class="label">Remaining</span><span class="value">' +
      remaining +
      "</span></div>" +
      '<div class="stat-pill"><span class="label">Vehicles</span><span class="value">' +
      vehicleCount +
      "</span></div>";
  }

  function getActiveTeam() {
    return activeTeamId ? teams.get(activeTeamId) : null;
  }

  function updateActiveTeam(team) {
    if (team && team.id) {
      teams.set(team.id, team);
      // Update tab label
      const tabs = getTabBar().querySelectorAll(".tab");
      tabs.forEach(function (t) {
        if (t.dataset.teamId === team.id) {
          t.textContent = team.name || "New Team";
        }
      });
    }
  }

  function showEmptyState() {
    const content = getMainContent();
    content.innerHTML =
      '<div class="empty-state">' +
      "<h2>WELCOME TO GASLANDS.TOOLS</h2>" +
      "<p>Build and manage your Gaslands Refuelled team rosters.<br>" +
      "Select a sponsor, set your budget, add vehicles, and equip them for the wasteland.</p>" +
      '<button class="create-btn" id="empty-create-btn">CREATE TEAM</button>' +
      "</div>";

    document.getElementById("empty-create-btn").addEventListener("click", createTeamTab);
    document.getElementById("budget-bar").style.display = "none";
  }

  function init() {
    document.getElementById("tab-new").addEventListener("click", createTeamTab);
    // Show empty state initially
    showEmptyState();
  }

  return {
    init: init,
    createTeamTab: createTeamTab,
    activateTab: activateTab,
    getActiveTeam: getActiveTeam,
    updateActiveTeam: updateActiveTeam,
    renderActiveTeam: renderActiveTeam,
  };
})();
