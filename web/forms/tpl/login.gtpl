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
</body>
</html>