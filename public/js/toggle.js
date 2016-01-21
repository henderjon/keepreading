define([], function () {
	"use strict";

	var hiddenClassName = "--collapsed"
	var legends = document.getElementsByTagName("legend")
	for(var i = 0; i < legends.length; i += 1){
		legends[i].addEventListener("click", function(){
			var that = this;
			var stateIndex = that.parentElement.className.indexOf(hiddenClassName);
			if(stateIndex != -1){
				that.parentElement.className = that.parentElement.className.substring(0, stateIndex);
			}else{
				that.parentElement.className += hiddenClassName;
			}
		});
	}

});


