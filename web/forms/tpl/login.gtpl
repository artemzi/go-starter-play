<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    <input type="checkbox" name="interest" value="football">Футбол
    <input type="checkbox" name="interest" value="basketball">Баскетбол
    <input type="checkbox" name="interest" value="tennis">Теннис

	Имя пользователя:<input type="text" name="username">
	Пароль:<input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="Войти">
</form>
<hr />
<form enctype="multipart/form-data" action="http://127.0.0.1:8080/upload" method="post">
	<input type="file" name="uploadfile" />
	<input type="hidden" name="token" value="{{.}}"/>
	<input type="submit" value="upload" />
</form>
</body>
</html>