// vue
var app = new Vue({
  el: '#app',
  data: {
    Identity: "",
    statusColor: "",
    status: "Loading ...",
    marked: false,
    info: {},
    marks:[],
    ticket: "",
  },
  created: function() {
    this.ticket = GetQueryVariable("t");
    if (!this.ticket) window.location.href = "./index.html";
    this.getInfo();
    this.getMarks();
  },
  methods: {
    mark: function() {
      if (this.marked) return;
      if (this.links === []) return;
      axios // Add mark
        .post("./v1/util", {
          t: this.ticket,
          type: "SELF"
        })
        .then((resp) => {
          if (resp.data.Success) this.getMarks();
          else this.message(resp.data.Message, "#57261F");
        })
        .catch(this.errorFunc);
    },
    getInfo: function() {
      axios // get info
        .get('./v1/info/' + this.ticket)
        .then((resp) => {
          if (resp.data.Success) {
            this.message("Loading ... ", "");
            this.identity = resp.data.Message;
            this.info = JSON.parse(this.identity.Raw.Info);
          } else {
            this.errorFunc(resp.data.Message);
          }
        })
        .catch(this.errorFunc);
    },
    getMarks: function() {
      axios // get marks
        .get('./v1/util/' + this.ticket)
        .then((resp) => {
          if (resp.data.Success) {
            this.marks = resp.data.Message;
            for (i in resp.data.Message) {
              let m = resp.data.Message[i];
              if (m.Type == "SELF") {
                this.marked = true;
                break;
              }
            }
          } else {
            this.errorFunc(resp.data.Message);
          }
        })
        .catch(this.errorFunc);
    },
    // Show message
    message: function(msg, color) {
      console.log(msg);
      this.status = msg;
      this.statusColor = "font-size: 1rem; background-color: " + color + ";";
      setTimeout(() => { // recovery
        this.status = "XYZUA";
        this.statusColor = "";
      }, 1000);
    },
    // wrapped catch function
    errorFunc: function(error) {
      console.log(error);
      this.message(error, "#57261F");
    }
  }
})