document.addEventListener("DOMContentLoaded", function () {
    CodeMirror.fromTextArea(document.querySelector("#new_contents"),
			    {async: true,
			     lineNumbers: true,
			     matchBrackets: true,
			     autoCloseBrackets: true,
			     lineWrapping: true,
			     viewportMargin: Infinity});
});
