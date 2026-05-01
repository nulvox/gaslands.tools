"use strict";

/**
 * URL-based team sharing.
 * Encodes team JSON (base64) in the URL hash.
 */

var Share = (function () {
  var PREFIX = "team=";

  function shareTeam(team) {
    if (!team) return;

    try {
      var json = JSON.stringify(team);
      var encoded = btoa(unescape(encodeURIComponent(json)));
      var url = window.location.origin + window.location.pathname + "#" + PREFIX + encoded;

      if (navigator.clipboard && navigator.clipboard.writeText) {
        navigator.clipboard.writeText(url).then(function () {
          alert("Share URL copied to clipboard!");
        }).catch(function () {
          prompt("Copy this URL to share:", url);
        });
      } else {
        prompt("Copy this URL to share:", url);
      }
    } catch (e) {
      alert("Failed to generate share URL: " + e.message);
    }
  }

  function checkForSharedTeam() {
    var hash = window.location.hash;
    if (!hash || hash.indexOf("#" + PREFIX) !== 0) return;

    try {
      var encoded = hash.substring(1 + PREFIX.length);
      var json = decodeURIComponent(escape(atob(encoded)));

      var result;
      if (typeof window.importTeam === "function") {
        result = JSON.parse(window.importTeam(json));
      } else {
        result = { team: JSON.parse(json), warnings: [] };
      }

      if (result.error) {
        console.error("Shared team import error:", result.error);
        return;
      }

      Tabs.importTeam(result.team, result.warnings);

      // Clear the hash to avoid re-importing on reload
      history.replaceState(null, "", window.location.pathname);
    } catch (e) {
      console.error("Failed to load shared team:", e.message);
    }
  }

  return {
    shareTeam: shareTeam,
    checkForSharedTeam: checkForSharedTeam
  };
})();
