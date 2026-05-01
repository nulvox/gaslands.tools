"use strict";

/**
 * JSON and HTML import/export functionality.
 */

var ImportExport = (function () {

  // 11.1: JSON export — download team as .gaslands.json
  function exportJSON(team) {
    if (!team) return;

    var jsonStr;
    if (typeof window.exportTeamJSON === "function") {
      jsonStr = window.exportTeamJSON(JSON.stringify(team));
    } else {
      jsonStr = JSON.stringify(team, null, 2);
    }

    var name = (team.name || "team").replace(/[^a-zA-Z0-9_-]/g, "_");
    downloadFile(name + ".gaslands.json", jsonStr, "application/json");
  }

  // 11.2: JSON import — file picker, parse, create tab
  function importJSON() {
    var input = document.createElement("input");
    input.type = "file";
    input.accept = ".json,.gaslands.json";
    input.addEventListener("change", function () {
      if (!input.files || !input.files[0]) return;

      var reader = new FileReader();
      reader.onload = function (e) {
        var contents = e.target.result;
        try {
          var result;
          if (typeof window.importTeam === "function") {
            result = JSON.parse(window.importTeam(contents));
          } else {
            result = { team: JSON.parse(contents), warnings: [] };
          }

          if (result.error) {
            alert("Import error: " + result.error);
            return;
          }

          // Create a new tab with the imported team
          Tabs.importTeam(result.team, result.warnings);
        } catch (err) {
          alert("Failed to import: " + err.message);
        }
      };
      reader.readAsText(input.files[0]);
    });
    input.click();
  }

  // 11.4: HTML export — generate styled HTML and download
  function exportHTML(team) {
    if (!team) return;

    if (typeof window.exportTeamHTML !== "function") {
      alert("WASM not loaded — cannot export HTML");
      return;
    }

    var html = window.exportTeamHTML(JSON.stringify(team));
    if (html.indexOf('"error"') === 0 || html.indexOf('{"error"') === 0) {
      alert("Export error: " + html);
      return;
    }

    var name = (team.name || "team").replace(/[^a-zA-Z0-9_-]/g, "_");
    downloadFile(name + "-roster.html", html, "text/html");
  }

  function downloadFile(filename, content, mimeType) {
    var blob = new Blob([content], { type: mimeType });
    var url = URL.createObjectURL(blob);
    var a = document.createElement("a");
    a.href = url;
    a.download = filename;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
    URL.revokeObjectURL(url);
  }

  return {
    exportJSON: exportJSON,
    importJSON: importJSON,
    exportHTML: exportHTML
  };
})();
