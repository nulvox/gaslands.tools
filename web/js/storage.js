"use strict";

/**
 * localStorage persistence for team data.
 * Teams stored as gaslands_team_{id}, manifest at gaslands_manifest.
 */

var Storage = (function () {
  var MANIFEST_KEY = "gaslands_manifest";
  var TEAM_PREFIX = "gaslands_team_";

  function isAvailable() {
    try {
      localStorage.setItem("__test__", "1");
      localStorage.removeItem("__test__");
      return true;
    } catch (e) {
      return false;
    }
  }

  function getManifest() {
    try {
      var data = localStorage.getItem(MANIFEST_KEY);
      return data ? JSON.parse(data) : [];
    } catch (e) {
      return [];
    }
  }

  function saveManifest(manifest) {
    try {
      localStorage.setItem(MANIFEST_KEY, JSON.stringify(manifest));
    } catch (e) {
      showStorageWarning();
    }
  }

  function saveTeam(team) {
    if (!team || !team.id) return false;

    try {
      localStorage.setItem(TEAM_PREFIX + team.id, JSON.stringify(team));

      // Update manifest
      var manifest = getManifest();
      var found = false;
      for (var i = 0; i < manifest.length; i++) {
        if (manifest[i].id === team.id) {
          manifest[i].name = team.name || "New Team";
          found = true;
          break;
        }
      }
      if (!found) {
        manifest.push({ id: team.id, name: team.name || "New Team" });
      }
      saveManifest(manifest);
      return true;
    } catch (e) {
      showStorageWarning();
      return false;
    }
  }

  function loadTeam(id) {
    try {
      var data = localStorage.getItem(TEAM_PREFIX + id);
      return data ? JSON.parse(data) : null;
    } catch (e) {
      return null;
    }
  }

  function listTeams() {
    return getManifest();
  }

  function deleteTeam(id) {
    try {
      localStorage.removeItem(TEAM_PREFIX + id);
      var manifest = getManifest().filter(function (m) { return m.id !== id; });
      saveManifest(manifest);
    } catch (e) {
      // Ignore errors on delete
    }
  }

  function showStorageWarning() {
    // Only show once
    if (document.getElementById("storage-warning")) return;

    var banner = document.createElement("div");
    banner.id = "storage-warning";
    banner.style.cssText =
      "position:fixed;bottom:0;left:0;right:0;background:var(--rust-dark);" +
      "color:var(--text-primary);padding:10px 16px;font-size:12px;text-align:center;" +
      "z-index:200;border-top:2px solid var(--rust);font-family:'Share Tech Mono',monospace";
    banner.textContent = "Warning: Unable to save to localStorage. Your changes may not persist.";

    var closeBtn = document.createElement("button");
    closeBtn.style.cssText =
      "background:none;border:none;color:var(--amber);cursor:pointer;" +
      "margin-left:12px;font-family:'Share Tech Mono',monospace;font-size:12px";
    closeBtn.textContent = "[dismiss]";
    closeBtn.addEventListener("click", function () { banner.remove(); });
    banner.appendChild(closeBtn);

    document.body.appendChild(banner);
  }

  return {
    isAvailable: isAvailable,
    saveTeam: saveTeam,
    loadTeam: loadTeam,
    listTeams: listTeams,
    deleteTeam: deleteTeam
  };
})();
