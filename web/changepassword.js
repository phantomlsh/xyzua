axios.defaults.withCredentials = true; // enable cookie
var SALT = "XYZUA_Powered_by_Phantomlsh"; // common salt
var storage = window.localStorage; // storage

// vue
var app = new Vue({
  el: '#app',
  data: {
    username: "",
    oldPassword: "",
    newPassword: "",
    repeatPassword: "",
    styleColor: "",
    tip: "Change Password",
    blocked: false,
  },
  created: function() {
    this.username = GetQueryVariable("username");
    if (!this.username) {
      this.errorFunc("Username Not Found");
    }
  },
  methods: {
    changePassword: function() {
      if (this.blocked) return;
      if (!(this.oldPassword && this.newPassword && this.repeatPassword)) {
        this.message("Please Complete Form", "#57261F", 2000);
        return;
      }
      if (this.newPassword != this.repeatPassword) {
        this.message("Repeat Password is not Consistent", "#57261F", 2000);
        return;
      }
      if (CheckDifficulty(this.newPassword) < 3) {
        this.message("New Password too Simple!", "#57261F", 2000);
        return;
      }
      this.blocked = true;
      this.message("Changing ...", "", 0); // maintain
      this.loginGet(); // begin login
    },
    loginGet: function() {
      axios // http request CHALLENGE
        .get('./v1/identity/' + this.username)
        .then((resp) => {
          let random = resp.data;
          setTimeout(() => { // next step
            this.changePasswordPut(random);
          }, 2000);
        })
        .catch(this.errorFunc);
    },
    changePasswordPut: function(random) {
      axios // changepassword
        .put('./v1/info', {
          password: HASH(HASH(this.oldPassword, SALT), random),
          newpassword: HASH(this.newPassword, SALT)
        })
        .then((resp) => {
          if (resp.data.Success) { // login success
            this.message("Succeeded. Relogin... ", "#2E571F", 2000);
            JumpLater("./index.html", 1800);
          } else {
            this.message(resp.data.Message, "#57261F", 2000);
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
        this.tip = "Change Password";
        this.styleColor = "";
        this.blocked = false;
      }, ms)
    },
    // wrapped catch function
    errorFunc: function(error) {
      this.message(error, "#57261F", 2000);
      JumpLater("./index.html", 2000);
    }
  }
})