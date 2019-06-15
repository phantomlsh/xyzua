var storage = window.localStorage; //storage

// vue
var app = new Vue({
  el: '#app',
  data: {
    identity: "",
    statusColor: "",
    status: "Loading ...",
    marked: false,
    info: {},
    links: [],
  },
  created: function() {
    if (!storage["XYZUA_Token"]) window.location.href = "./index.html";
    // check jumping
    if (storage["Appname"]) {
      this.jumpApp(storage["Appname"], this.getInfo);
      storage.removeItem("Appname");
    } else {
      this.getInfo();
    }
  },
  methods: {
    logout: function() {
      Logout();
    },
    changePassword: function() {
      if (!this.identity) return;
      this.message("Going to ChangePassword ...", "");
      JumpLater("./changepassword.html?username=" + this.identity.Identity, 1000);
    },
    mark: function() {
      if (this.marked) return;
      if (this.links === []) return;
      axios // Add mark
        .post("./v1/util", {
          t: storage["XYZUA_Token"],
          type: "SELF"
        })
        .then((resp) => {
          if (resp.data.Success) this.marked = true;
          else this.message(resp.data.Message, "#57261F");
        })
        .catch(this.errorFunc);
    },
    getInfo: function() {
      axios // get info
        .get('./v1/info/' + storage["XYZUA_Token"])
        .then((resp) => {
          if (resp.data.Success) {
            this.message("Loading ... ", "");
            this.identity = resp.data.Message;
            this.info = JSON.parse(this.identity.Raw.Info);
            // check ChangePassword
            if (storage["ChangePassword"]) this.changePassword();
            // update token
            storage["XYZUA_Token"] = this.identity.Raw.Token;
            this.getMarks();
          } else {
            this.errorFunc(resp.data.Message);
          }
        })
        .catch(this.errorFunc);
      axios // get links
        .get('./v1/util')
        .then((resp) => {
          this.links = resp.data;
        })
        .catch(this.errorFunc);   
    },
    getMarks: function() {
      axios // get marks
        .get('./v1/util/' + storage["XYZUA_Token"])
        .then((resp) => {
          if (resp.data.Success) {
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
    jumpApp: function(appname, callback) {
      axios // get callback
        .post("./v1/info", {
          token: storage["XYZUA_Token"],
          appname: appname
        })
        .then((resp) => {
          if (resp.data.Success) {
            this.message("Jumping to " + appname + " ...", "");
            JumpLater(resp.data.Message.Url, 1000);
            // update token
            storage["XYZUA_Token"] = resp.data.Message.Token;
          } else {
            if (resp.data.Message.Token) {
              storage["XYZUA_Token"] = resp.data.Message.Token;
              callback(); // fail but continue
            } else {
              // completely fail
              this.errorFunc(resp.data.Message);
            }
          }
        })
        .catch(this.errorFunc);
    },
    // Show message
    message: function(msg, color) {
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
      storage.removeItem("XYZUA_Token");
      JumpLater("./index.html", 1000);
    }
  }
})