function templete(){
	const leftSideNav = {
		"ホーム": "",
		"トレンド": "trends",
		"通知": "notifications",
		"メッセージ": "messages",
		"ブックマーク": "bookmarks",
		"リスト": "lists",
		"プロフィール": "profile"
	};
	
	const contentWrap = document.createElement("div");
	contentWrap.id = "content-wrap"

	// サイドバーhidari no
	const leftSide = document.createElement("aside");
	leftSide.id = "left-sidebar";

	for(const nav in leftSideNav){
		const button = document.createElement("span");
		button.className = "left-side-button";
		button.innerHTML = nav;
		button.dataset.href = "/" + leftSideNav[nav];
		leftSide.appendChild(button)
	}
	contentWrap.appendChild(leftSide);

	// 真ん中
	const content = document.createElement("article");
	content.id = "content";
	contentWrap.appendChild(content)

	document.body.appendChild(contentWrap);
}

addEventListener("DOMContentLoaded", templete, false);