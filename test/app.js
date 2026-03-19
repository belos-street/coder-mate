let wasmModule = null;

async function loadWasm() {
    try {
        const response = await fetch('../bin/coder-mate.wasm');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const wasmBuffer = await response.arrayBuffer();

        const go = new Go();
        wasmModule = await WebAssembly.instantiate(wasmBuffer, go.importObject);

        go.run(wasmModule.instance);

        console.log('WASM loaded successfully');
        document.getElementById('result').textContent = 'WASM module loaded successfully! Click the button to test.';
        document.getElementById('result').className = 'success';
        document.getElementById('callBtn').disabled = false;
    } catch (error) {
        console.error('Error loading WASM:', error);
        document.getElementById('result').textContent = `Error loading WASM: ${error.message}`;
        document.getElementById('result').className = 'error';
    }
}

window.callWasm = function () {
    if (!wasmModule) {
        document.getElementById('result').textContent = 'WASM module not loaded yet. Please wait...';
        return;
    }

    try {
        const code = `const x = 42;
function hello() {
  // This is a comment
  return "Hello World";
}`;

        const result = window.highlightCode(code, 'javascript', 'dark');

        document.getElementById('result').textContent = result;
        document.getElementById('result').className = 'success';
    } catch (error) {
        console.error('Error calling WASM function:', error);
        document.getElementById('result').textContent = `Error calling WASM: ${error.message}`;
        document.getElementById('result').className = 'error';
    }
};

window.onload = loadWasm;