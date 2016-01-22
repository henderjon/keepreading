define([], function () {
	"use strict";

	var parse = function() {
		var o = {}, i,
			lines = document.cookie.split('; '),
			values;

		for (i = 0; i < lines.length; i++) {
			values = lines[i].split('=');
			if (values.length > 1) {
				o[values[0].trim()] = values.slice(1).join('=').trim();
			}
		}

		return o;
	};

	var cookie = {
		'set' : function(cname, cvalue, exdays){
			var d = new Date();
			d.setTime(d.getTime() + (exdays*1000*60*60*24));
			var expires = "expires="+d.toUTCString();
			document.cookie = cname + "=" + encodeURIComponent(cvalue) + "; " + expires;
		},
		'get' : function(cname) {
			var values = parse();
			if (values.hasOwnProperty(cname)) {
				return decodeURIComponent(values[cname]);
			}
			return "";
		},
		'all' : function(){
			return document.cookie;
		}
	};

	return cookie;

});
