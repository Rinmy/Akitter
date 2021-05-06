const url = [];
url.profile = function(path){
	const id = path[1];
	alert(id);
}

function selectUrl(){
	const path = location.pathname.slice(1).split("/");
	url[path[0]](path);
}

addEventListener("DOMContentLoaded", selectUrl, false);