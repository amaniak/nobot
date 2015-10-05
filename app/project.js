
var Project = function(){}
Project.prototype = new Nobot("project", "http://api.openweathermap.org/data/2.5/weather?q=London,uk", true);


project = new Project();
project.run();

project.data
