"use strict";

async function initWasm() {
  const statusEl = document.getElementById("loading-status");

  try {
    const go = new Go();
    const result = await WebAssembly.instantiateStreaming(
      fetch("main.wasm"),
      go.importObject
    );
    go.run(result.instance);
    statusEl.textContent = "WASM loaded. Ready.";
  } catch (err) {
    console.error("Failed to load WASM:", err);
    statusEl.textContent = "Failed to load WASM: " + err.message;
  }
}

document.addEventListener("DOMContentLoaded", initWasm);
