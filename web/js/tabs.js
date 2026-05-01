"use strict";

/**
 * Tab management for multi-team support.
 * Each tab corresponds to a team. Team data is stored as parsed JSON.
 */

var Tabs = (function () {
  var teams = new Map(); // id -> team JSON object
  var activeTeamId = null;

  function getTabBar() {
    return document.getElementById("tab-bar");
  }

  function createTeamTab() {
    if (typeof window.createTeam !== "function") {
      console.error("WASM not loaded yet");
      return;
    }

    var teamJSON = window.createTeam();
    var team = JSON.parse(teamJSON);
    teams.set(team.id, team);

    var tabBar = getTabBar();
    var btn = document.createElement("button");
    btn.className = "tab";
    btn.dataset.teamId = team.id;
    btn.textContent = team.name || "New Team";
    btn.addEventListener("click", function () {
      activateTab(team.id);
    });

    // Insert before the + button
    var newBtn = document.getElementById("tab-new");
    tabBar.insertBefore(btn, newBtn);

    activateTab(team.id);
    console.log("Created team:", teamJSON);
  }

  function activateTab(teamId) {
    if (!teams.has(teamId)) return;

    activeTeamId = teamId;

    // Update tab styles
    var tabs = getTabBar().querySelectorAll(".tab");
    tabs.forEach(function (t) {
      t.classList.toggle("active", t.dataset.teamId === teamId);
    });

    // Render using TeamUI
    var team = teams.get(activeTeamId);
    if (team) {
      TeamUI.refreshTeam(team);
    }
  }

  function getActiveTeam() {
    return activeTeamId ? teams.get(activeTeamId) : null;
  }

  function updateActiveTeam(team) {
    if (team && team.id) {
      teams.set(team.id, team);
      // Update tab label
      var tabs = getTabBar().querySelectorAll(".tab");
      tabs.forEach(function (t) {
        if (t.dataset.teamId === team.id) {
          t.textContent = team.name || "New Team";
        }
      });
    }
  }

  function showEmptyState() {
    var content = document.getElementById("main-content");
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

  function restoreFromStorage() {
    var savedTeams = Storage.listTeams();
    if (savedTeams.length === 0) {
      showEmptyState();
      return;
    }

    var lastId = null;
    savedTeams.forEach(function (entry) {
      var team = Storage.loadTeam(entry.id);
      if (!team) return;

      teams.set(team.id, team);

      var tabBar = getTabBar();
      var btn = document.createElement("button");
      btn.className = "tab";
      btn.dataset.teamId = team.id;
      btn.textContent = team.name || "New Team";
      btn.addEventListener("click", function () {
        activateTab(team.id);
      });

      var newBtn = document.getElementById("tab-new");
      tabBar.insertBefore(btn, newBtn);

      lastId = team.id;
    });

    if (lastId) {
      activateTab(lastId);
    }
  }

  function init() {
    document.getElementById("tab-new").addEventListener("click", createTeamTab);
    restoreFromStorage();
  }

  function importTeam(team, warnings) {
    if (!team || !team.id) return;

    teams.set(team.id, team);

    var tabBar = getTabBar();
    var btn = document.createElement("button");
    btn.className = "tab";
    btn.dataset.teamId = team.id;
    btn.textContent = team.name || "New Team";
    btn.addEventListener("click", function () {
      activateTab(team.id);
    });

    var newBtn = document.getElementById("tab-new");
    tabBar.insertBefore(btn, newBtn);

    activateTab(team.id);
    Storage.saveTeam(team);
  }

  return {
    init: init,
    createTeamTab: createTeamTab,
    activateTab: activateTab,
    getActiveTeam: getActiveTeam,
    updateActiveTeam: updateActiveTeam,
    importTeam: importTeam
  };
})();
