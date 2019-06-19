var storage = window.localStorage;

// vue
var app = new Vue({
  el: '#app',
  data: {
    Identity: "",
    statusColor: "",
    status: "Loading ...",
    marked: false,
    info: {},
    tag: "test",
    marks:[],
    ticket: "",
    isChecked: false,
    isToday: false,
  },
  // created: function() {
  //   this.ticket = GetQueryVariable("t");
  //   if (!this.ticket) window.location.href = "./index.html";
  //   this.getInfo();
  //   this.getMarks();
  // },
  methods: {
    mark: function(type) {
      // check all kinds of conditions
      if (this.links === []) return;
      if (type == 'TODAY' && this.isToday) return;
      if (type == 'CHECK' && this.isChecked) return;
      if (type == 'SELF' && this.marked) return;
      let admintoken = storage["XYZUA_Token"];
      if (type != "SELF" && !admintoken) return;

      if (type == 'TODAY') type = GetDateString();  
      axios // Add mark
        .post("./v1/util", {
          admintoken: admintoken,
          t: this.ticket,
          type: type
        })
        .then((resp) => {
          if (resp.data.Success) this.getMarks();
          else this.errorFunc(resp.data.Message);
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
            this.tag = this.identity.Raw.Tag;
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
              if (m.Type == "SELF") this.marked = true;
              if (m.Type == "CHECK") this.isChecked = true;
              if (m.Type == GetDateString()) this.isToday = true;
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
      this.message(error, "#57261F");
    }
  }
})