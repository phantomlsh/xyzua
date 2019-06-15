axios.defaults.withCredentials = true; // enable cookie
var SALT = "XYZUA_Powered_by_Phantomlsh"; // common salt
var storage = window.localStorage; // storage

// vue
var app = new Vue({
  el: '#app',
  data: {
    username: "",
    password: "",
    styleColor: "",
    tip: "Tap anywhere to Login",
    blocked: false,
  },
  created: function() {
    // record Appname
    let appname = GetQueryVariable("App");
    if (appname) storage["Appname"] = appname;
    else storage.removeItem("Appname");
    if (!storage["XYZUA_Token"]) return;
    // Automatically login
    this.blocked = true;
    axios // check token
      .put('./v1/identity/' + storage["XYZUA_Token"])
      .then((resp) => {
        if (resp.data.Success) { // login success
          this.message("Login Automatically. Loading...", "#2E571F", 2000);
          JumpLater("./home.html", 2000);
          //update token
          storage["XYZUA_Token"] = resp.data.Message;
        } else {
          this.blocked = false;
          storage.removeItem("XYZUA_Token");
        }
      })
      .catch(this.errorFunc);
  },
  methods: {
    login: function() {
      if (this.blocked) return;
      if (this.username === "") {
        this.errorFunc("Please complete username");
        return;
      }
      if (this.password === "") {
        this.errorFunc("Please complete password");
        return;
      }
      this.blocked = true;
      if (CheckDifficulty(this.password) < 3) storage["ChangePassword"] = "true";
      else storage.removeItem("ChangePassword");
      this.message("Login ...", "", 0); // maintain
      this.loginGet(); // begin login
    },
    loginGet: function() {
      axios // http request CHALLENGE
        .get('./v1/identity/' + this.username)
        .then((resp) => {
          let random = resp.data;
          setTimeout(() => { // next step
            this.loginPost(random);
          }, 2000);
        })
        .catch(this.errorFunc);
    },
    loginPost: function(random) {
      axios // http request RESPONSE
        .post('./v1/identity', {
          password: HASH(HASH(this.password, SALT), random)
        })
        .then((resp) => {
          if (resp.data.Success) { // login success
            this.message("Login Successfully. Loading...", "#2E571F", 2000);
            JumpLater("./home.html", 1000);
            // update token
            storage["XYZUA_Token"] = resp.data.Message;
          } else { // login fail
            this.errorFunc(resp.data.Message);
            storage.removeItem("XYZUA_Token");
          }
        })
        .catch(this.errorFunc);
    },
    // Show message
    message: function(msg, color, ms) {
      this.tip = msg;
      this.styleColor = "background-color: " + color + ";";
      this.blocked = true;
      //recover
      if (!ms) return;
      setTimeout(() => {
        this.tip = "Tap anywhere to Login";
        this.styleColor = "";
        this.blocked = false;
      }, ms)
    },
    // wrapped catch function
    errorFunc: function(error) {
      this.message(error, "#57261F", 2000);
    }
  }
})