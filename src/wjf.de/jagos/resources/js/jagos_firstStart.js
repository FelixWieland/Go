$("#jagos-firstStart-connect").click(function () {

	sIP = $("input[name=jagos-config-mysql-ip]").first().val();
	sUser = $("input[name=jagos-config-mysql-user]").first().val();
	sPassword = $("input[name=jagos-config-mysql-password]").first().val();
	$.post("/AJAX/SYSTEM/setConfig", {"Server": sIP,
																		"User": sUser,
																		"Password": sPassword}, function(data) {
		console.log(data);
	});
});
