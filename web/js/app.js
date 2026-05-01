"use strict";

async function initWasm() {
  var statusEl = document.getElementById("loading-status");

  try {
    var go = new Go();
    var result = await WebAssembly.instantiateStreaming(
      fetch("main.wasm"),
      go.importObject
    );
    go.run(result.instance);

    // WASM loaded — initialize tabs
    if (statusEl) {
      statusEl.remove();
    }
    Tabs.init();
  } catch (err) {
    console.error("Failed to load WASM:", err);
    if (statusEl) {
      statusEl.textContent = "Failed to load WASM: " + err.message;
    }
  }
}

document.addEventListener("DOMContentLoaded", initWasm);
