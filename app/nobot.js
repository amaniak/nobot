console.log("loading nobot")
//console.log(nobotHTTP("http://google.nl"))

var Nobot = function (name, endpoint, debug) {
  this.name     = name;
  this.endpoint = endpoint;
  this.debug    = debug || false;
  this.data     = null;

}

Nobot.prototype.setup = function() {
  // Do setup stuff here.
}

Nobot.prototype.run = function() {
  // Run the actual robot.

  if(this.debug)
    console.log("Loaded module: [" + this.name + "]");

  this.data = nobotHTTP(this.endpoint)

};