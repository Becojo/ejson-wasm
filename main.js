const go = new Go();
WebAssembly.instantiateStreaming(fetch("dist/ejson.wasm"), go.importObject).then(async (result) => {
  go.run(result.instance);
});

document.querySelector('#encrypt').addEventListener('click', () => {
  var input = document.querySelector('#input');
  var encrypted = ejson.encrypt(input.value);

  if(encrypted != false) {
    input.value = encrypted;
  }
});
