function Logout() {
  console.log("logout");
  var storage = window.localStorage;
  if (!storage["XYZUA_Token"]) return;
  axios
    .delete('./v1/identity/' + storage["XYZUA_Token"])
    .then((resp) => {
      storage.removeItem("XYZUA_Token");
      window.location.href = "./index.html";
    })
    .catch((error) => {
      storage.removeItem("XYZUA_Token");
      window.location.href = "./index.html";
    });
}

// hash function
function HASH(msg, salt) {
  var sha256 = new Hashes.SHA256().hex(msg + salt);
  var md5 = new Hashes.MD5().hex(sha256 + salt);
  return md5;
}

function GetQueryVariable(variable) {
  var query = window.location.search.substring(1);
  var vars = query.split("&");
  for (var i = 0; i < vars.length; i++) {
    var pair = vars[i].split("=");
    if (pair[0] == variable) return pair[1];
  }
  return (false);
}

function JumpLater(url, ms) {
  setTimeout(function() {
    window.location.href = url;
  }, ms);
}

function CheckDifficulty(val) {
  var diff = 0;
  if (val.length < 6) return 0;
  if (val.length >= 8) diff++;
  if (/\d/.test(val)) diff++; // number
  if (/[a-z]/.test(val)) diff++; // little letter
  if (/[A-Z]/.test(val)) diff++; // capital letter
  if (/\W/.test(val)) diff++; // special character
  return diff;
}