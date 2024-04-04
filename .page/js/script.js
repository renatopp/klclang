$(function () {

  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("./js/klc.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
  });

  $('#terminal').terminal(function (command, term) {

    if (command === '') {
      term.echo('')
      return
    }

    const res = klc(command)
    var div = ''
    if (res.startsWith('File')) {
      div = $('<span class="response response-error">' + res + '</span>')
    } else if (res.startsWith('<')) {
      div = $('')
    } else {
      div = $('<span class="response response-ok">' + res + '</span>')
    }
    term.echo(div);
  }, {
    greetings: "",
    prompt: '? ',
  }).focus();

  $.terminal.syntax('haskell')
});