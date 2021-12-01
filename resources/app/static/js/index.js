const globals = {
  timer: null,
  isAstilectronReady: false,
  resultNode: null,
  currentText: "",
};

function docReady(fn) {
  if (
    document.readyState === "complete" ||
    document.readyState === "interactive"
  ) {
    setTimeout(fn, 1);
  } else {
    document.addEventListener("DOMContentLoaded", fn);
  }
}

function debaunce(fn, time = 200) {
  if (globals.timer) clearTimeout(globals.timer);
  globals.timer = setTimeout(fn, time);
}

function formatText(delta, oldDelta) {
  let text = oldDelta.ops[0].insert;
  let pointer = text.length - 2;

  delta.ops.forEach((op) => {
    if (op.insert)
      text = `${text.substring(0, pointer)}${op.insert}${text.substring(
        pointer
      )}`;
    if (op.retain) pointer = op.retain;
    if (op.delete)
      text =
        pointer <= 0 || pointer + op.delete >= text.length
          ? ""
          : `${text.substring(0, pointer)}${text.substring(
              pointer + op.delete
            )}`;
  });
  return text;
}

function sendChange(text) {
  if (!text) {
    globals.resultNode.dispatchEvent(new CustomEvent("empty"));
    return;
  }
  globals.currentText = text;

  const message = { name: "process", payload: text };
  astilectron.sendMessage(message, function (message) {
    console.log(message);
    if (typeof message.payload[0] === "string") {
      globals.resultNode.dispatchEvent(
        new CustomEvent("error", { detail: message.payload })
      );
      return;
    }
    globals.resultNode.dispatchEvent(
      new CustomEvent("update", { detail: message.payload })
    );
  });
}

function configureQuill() {
  globals.resultNode = document.querySelector("#res");
  const quill = new Quill("#editor", {
    formats: [],
  });

  quill.on("text-change", function (delta, oldDelta, source) {
    // globals.resultNode.dispatchEvent(new CustomEvent("proc"))
    if (source != "user" || !globals.isAstilectronReady) return;
    debaunce(() => sendChange(formatText(delta, oldDelta)));
  });
}

function defineData() {
  Alpine.data("results", () => ({
    segments: [],
    res: false,
    plain: false,
    proc: "",
    error: "",
    result: {
      ["@update"]({ detail }) {
        this.res = true;
        this.error = "";
        this.segments = detail;
        this.proc = "";
        this.plain = false;
      },
      ["@plain"]({ detail }) {
        this.res = true;
        this.error = "";
        this.segments = detail;
        this.proc = "";
        this.plain = true;
      },
      ["@empty"]() {
        this.res = false;
        this.error = "";
        this.proc = "";
      },
      ["@proc"]() {
        this.res = false;
        this.error = "";
        this.proc = "Esperando . . .";
      },
      ["@error"]({ detail }) {
        this.res = false;
        this.error = detail;
        this.proc = "";
      },
    },
  }));
}

document.addEventListener("alpine:init", defineData);
document.addEventListener("astilectron-ready", () => {
  globals.isAstilectronReady = true;
  astilectron.onMessage(function (msg) {
    let handler = ()=>{}
    if(msg==="run"){
      handler = (result)=>{
        console.log(result)
        globals.resultNode.dispatchEvent(
          new CustomEvent("plain", { detail: result.payload })
        );
      }
    }
    console.log(globals.currentText);
    const message = { name: msg, payload: globals.currentText };
    astilectron.sendMessage(message, handler);
    return;
  });
});
docReady(configureQuill);
