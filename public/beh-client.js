(function(window){
	window.onerror = function (msg, url, lineNo, columnNo, error) {
		var url = "http://localhost:12345/event";
		var method = "POST";
		var postData = JSON.stringify({
			msg: msg,
			url: url,
			lineNo: lineNo,
			columnNo: columnNo,
			error: JSON.stringify(error)
		});
		var async = true;
		var request = new XMLHttpRequest();
		request.open(method, url, async);
		request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
		request.send(postData);
	};
})(window);
