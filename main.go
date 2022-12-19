package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "name_db"
)

func checkError(err error) { // прописываем ошибки
	if err != nil {
		panic(err)
	}
}

func connectTo() string { // connect to db

	sqlConn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	return sqlConn
}

func openDb() *sql.DB { // open db

	db, err := sql.Open("postgres", connectTo())
	checkError(err)

	return db
}

func showkluchSchool() { // выводим запись по ключу для таблицы школа
	var num int
	var adres string
	fmt.Print("Введите номер школы:\n")
	fmt.Fscan(os.Stdin, &num)
	rows, err := openDb().Query(`SELECT * from "bdSchool"."школа" where "номер_школы"=$1`, &num)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&num, &adres) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Адрес школы: ", adres)
	}
	checkError(err)
}

func showkluchClass() { // выводим запись по ключу для таблицы класс
	var num int
	var fio string
	fmt.Print("Введите номер класса:\n")
	fmt.Fscan(os.Stdin, &num)
	rows, err := openDb().Query(`SELECT * from "bdSchool"."класс" where "номер_класса"=$1`, &num)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&num, &fio) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("ФИО классного руководителя: ", fio)
	}
	checkError(err)
}

func showkluchStudent() { // выводим запись по ключу для таблицы ученик
	var fio, sex string
	fmt.Print("Введите ФИО ученика:\n")
	fmt.Fscan(os.Stdin, &fio)
	rows, err := openDb().Query(`SELECT * from "bdSchool"."ученик" where "ФИО"=$1`, &fio)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&fio, &sex, ) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Пол ученика: ", sex)
	}
	checkError(err)
}

func showkluchMark() { // выводим запись по ключу для таблицы успеваемость
	var mark int
	var predmet string
	fmt.Print("Введите предмет:\n")
	fmt.Fscan(os.Stdin, &predmet)
	rows, err := openDb().Query(`SELECT * from "bdSchool"."успеваемость" where "название_предмета"=$1`, &predmet)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&predmet, &mark) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Оценка: ", mark)
	}
	checkError(err)
}

func showkluchEazy() { // выводим запись по ключу для таблицы участие
	var time int
	var vid string
	fmt.Print("Введите вид деятельности:\n")
	fmt.Fscan(os.Stdin, &vid)
	rows, err := openDb().Query(`SELECT * from "bdSchool"."участие" where "вид_деятельности"=$1`, &vid)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&vid, &time) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Количество часов: ", time)
	}
	checkError(err)
}

func showSchool() { // показываем таблицу школа
	rows, err := openDb().Query(`SELECT "номер_школы", "адрес" FROM "bdSchool"."школа"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var num int
		var adres string

		err = rows.Scan(&num, &adres)
		checkError(err)

		fmt.Println("Адрес школы: ", adres, "номер школы: ", num)
	}

	checkError(err)
}

func showClass() { // показываем таблицу класс
	rows, err := openDb().Query(`SELECT "номер_класса", "ФИО_класс_рук" FROM "bdSchool"."класс"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var num int
		var fio string

		err = rows.Scan(&num, &fio)
		checkError(err)

		fmt.Println("Номер класса:", num, "Классный руководитель: ", fio)
	}

	checkError(err)
}

func showStudent() { // показываем таблицу ученик
	rows, err := openDb().Query(`SELECT "ФИО", "пол" FROM "bdSchool"."ученик"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {

		var fio, sex string

		err = rows.Scan(&fio, &sex)
		checkError(err)

		fmt.Println("ФИО ученика: ", fio, "Пол ученика: ", sex)
	}

	checkError(err)
}

func showMark() { // показываем таблицу успеваемость
	rows, err := openDb().Query(`SELECT "название_предмета", "оценка" FROM "bdSchool"."успеваемость"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var mark int
		var predmet string

		err = rows.Scan(&predmet, &mark)
		checkError(err)

		fmt.Println("Название предмета: ", predmet, "Успеваемость: ", mark)
	}

	checkError(err)
}

func showEazy() { // показываем таблицу участие
	rows, err := openDb().Query(`SELECT "вид_деятельности", "кол-во_часов" FROM "bdSchool"."участие"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var time int
		var vid string

		err = rows.Scan(&vid, &time)
		checkError(err)

		fmt.Println("Вид деятельности: ", vid, "количество часов: ", time)
	}

	checkError(err)
}

func delSchool() { // удаляем из таблицы школа
	var check int
	fmt.Print("какую запись удалить? введите номер школы:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "bdSchool"."школа" where "номер_школы"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delClass() { // удаляем из таблицы класс
	var check int
	fmt.Print("какую запись удалить? введите номер класса:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "bdSchool"."класс" where "номер_класса"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delPerson() { // удаляем из таблицы ученик
	var check string
	fmt.Print("какую запись удалить? введите ФИО ученика\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "bdSchool"."ученик" where "ФИО"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delMark() { // удаляем из таблицы успеваемость
	var check string
	fmt.Print("какую запись удалить? Введите название_предмета:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "bdSchool"."успеваемость" where "название_предмета"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delEazy() { // удаляем из таблицы участие
	var check string
	fmt.Print("какую запись удалить? введите вид деятельности:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "bdSchool"."участие" where "вид_деятельности"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func updateClass() { // обновляем запись в таблице класс
	var num int
	var fio string

	fmt.Print("Введите номер_класса для обновления:\n")
	fmt.Fscan(os.Stdin, &num)

	fmt.Print("Введите ФИО классного руководителя:\n")
	fmt.Fscan(os.Stdin, &fio)

	updateStm := `update "bdSchool"."класс" set "ФИО_класс_рук"=$1 where "номер"=$2`
	_, e := openDb().Exec(updateStm, &fio, &num)
	checkError(e)
}

func updateSchool() { // обновляем запись в таблице школа
	var adres string
	var num int

	fmt.Print("Введите номер школы для обновления:\n")
	fmt.Fscan(os.Stdin, &num)

	fmt.Print("Введите новый адрес:\n")
	fmt.Fscan(os.Stdin, &adres)

	updateStm := `update "bdSchool"."школа" set "адрес"=$1 where "номер_школы"=$2`
	_, e := openDb().Exec(updateStm, &adres, &num)
	checkError(e)
}

func updateStudent() { // обновляем запись в таблице ученик
	var fio, sex string

	fmt.Print("Введите ФИО для обновления:\n")
	fmt.Fscan(os.Stdin, &fio)

	fmt.Print("Введите новый пол:\n")
	fmt.Fscan(os.Stdin, &sex)

	updateStm := `update "bdSchool"."ученик" set "пол"=$1 where "ФИО"=$2`
	_, e := openDb().Exec(updateStm, &sex, &fio )
	checkError(e)
}

func updateMark() { // обновляем запись в таблице успеваемость
	var predmet string
	var mark int

	fmt.Print("Введите предмет для  обновления:\n")
	fmt.Fscan(os.Stdin, &predmet)

	fmt.Print("Введите новую оценку:\n")
	fmt.Fscan(os.Stdin, &mark)

	updateStm := `update "bdSchool"."успеваемость" set "оценка"=$1 where "название_предмета"=$2`
	_, e := openDb().Exec(updateStm, &mark, &predmet)
	checkError(e)
}

func updateEazy() { // обновляем запись в таблице участие
	var vid string
	var time int

	fmt.Print("Введите вид деятельности для обновления:\n")
	fmt.Fscan(os.Stdin, &vid)

	fmt.Print("Введите новое количество часов:\n")
	fmt.Fscan(os.Stdin, &time)

	updateStm := `update "bdSchool"."участие" set "кол-во_часов"=$1 where "вид_деятельности"=$2`
	_, e := openDb().Exec(updateStm, &time, &vid)
	checkError(e)
}

func addSchool() { // insert function enter value to school table
	var adres string
	var number int
	fmt.Print("Введите адрес школы:\n")
	fmt.Fscan(os.Stdin, &adres)

	fmt.Print("Введите номер школы:\n")
	fmt.Fscan(os.Stdin, &number)

	insertToDyn := `insert into "bdSchool"."школа"("номер_школы","адрес") values($1, $2)`
	_, e := openDb().Exec(insertToDyn, &number, &adres)
	checkError(e)

}

func addClass() { // insert function enter value to class table
	var nomer int
	var fio string
	fmt.Print("Введите номер класса: \n")
	fmt.Fscan(os.Stdin, &nomer)

	fmt.Print("Введите ФИО классного руководителя: \n")
	fmt.Fscan(os.Stdin, &fio)


	insertToDyn := `insert into "bdSchool"."класс"("номер_класса","ФИО_класс_рук") values($1, $2)`
	_, e := openDb().Exec(insertToDyn, &nomer, &fio)
	checkError(e)

}

func addStudent() { // insert function enter value to student table
	var name, sex string
	fmt.Print("Введите ФИО ученика: \n")
	fmt.Fscan(os.Stdin, &name)

	fmt.Print("Введите пол ученика: \n")
	fmt.Fscan(os.Stdin, &sex)

	insertToDyn := `insert into "bdSchool"."ученик"("ФИО","пол") values($1, $2)`
	_, e := openDb().Exec(insertToDyn, &name, &sex)
	checkError(e)

}

func addMark() { // insert function enter value to mark table
	var predmet string
	var mark int
	fmt.Print("Введите название предмета: \n")
	fmt.Fscan(os.Stdin, &predmet)

	fmt.Print("Введите оценку: \n")
	fmt.Fscan(os.Stdin, &mark)

	insertToDyn := `insert into "bdSchool"."успеваемость"("название_предмета","оценка") values($1, $2)`
	_, e := openDb().Exec(insertToDyn, &predmet, &mark)
	checkError(e)

}

func addEazy() { // insert function enter value to prodavec table
	var vid string
	var time int
	fmt.Print("Вид деятельности: \n")
	fmt.Fscan(os.Stdin, &vid)

	fmt.Print("Количество часов: \n")
	fmt.Fscan(os.Stdin, &time)

	insertToDyn := `insert into "bdSchool"."участие"("вид_деятельности","кол-во_часов") values($1, $2)`
	_, e := openDb().Exec(insertToDyn, &vid, &time)
	checkError(e)

}

func add_switch_case() { // функция выбора таблицы для добавления записи
	var vibor string
	fmt.Print("выберите дейсвтие: addSchool -  чтобы добавить данные в таблицу школа\n addClass - добавить данные в таблицу класс\n addStudent добавить данные в таблицу класс\n addMark - добавить данные в таблицу успеваемость\n addEazy - добавить данные в таблицу участие\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "addSchool":
		addSchool()
	case "addClass":
		addClass()
	case "addStudent":
		addStudent()
	case "addMark":
		addMark()
	case "addEazy":
		addEazy()
	default:
		fmt.Println("Неправильная команда")
	}
}

func update_switch_case() { // функция выбора таблицы для обеовления
	var vibor string
	fmt.Print("выберите дейсвтие: updateSchool -  чтобы обновить данные в таблице школа\n updateClass - обновить данные в таблице класс\n updateStudent обновить данные в таблице ученик\n updateMark - обновить данные в таблице успеваемость\n updateEazy - обновить данные в таблице участие\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "updateSchool":
		updateSchool()
	case "updateClass":
		updateClass()
	case "updateStudent":
		updateStudent()
	case "updateMark":
		updateMark()
	case "updateEazy":
		updateEazy()
	default:
		fmt.Println("Неправильная команда")
	}
}

func delete_switch_case() { // функция выбора таблицы для удаления записи
	var vibor string
	fmt.Print("выберите дейсвтие: delSchool -  чтобы удалить данные из таблицы школа\n delClass - удалить данные из таблицы класс\n delPerson удалить данные из таблицы ученик\n delMark - удалить данные из таблицы успеваемость\n delEazy - удалить данные из таблицы участие\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "delSchool":
		delSchool()
	case "delClass":
		delClass()
	case "delPerson":
		delPerson()
	case "delMark":
		delMark()
	case "delEazy":
		delEazy()
	default:
		fmt.Println("Неправильная команда")
	}
}

func show_switch_case() { // функция выбора таблицы для выводы данных
	var vibor string
	fmt.Print("выберите дейсвтие: showSchool -  чтобы показать данные из таблицы школа\n showClass - показать данные из таблицы класс\n showStudent показать данные из таблицы ученик\n showMark - показать данные из таблицы успеваемость\n showEazy - показать данные из таблицы участие\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "showSchool":
		showSchool()
	case "showClass":
		showClass()
	case "showStudent":
		showStudent()
	case "showMark":
		showMark()
	case "showEazy":
		showEazy()
	default:
		fmt.Println("Неправильная команда")
	}
}

func show_klutch_switch_case() {
	var vibor string
	fmt.Print("выберите дейсвтие: showkluchSchool -  чтобы показать данные из таблицы школа\n showKluchClass - показать данные из таблицы класс\n showKluchStudent показать данные из таблицы ученик\n showKluchMark - показать данные из таблицы успеваемость\n showKluchEazy - показать данные из таблицы участие\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "showKluchSchool":
		showkluchSchool()
	case "showKluchClass":
		showkluchClass()
	case "showKluchStudent":
		showkluchStudent()
	case "showKluchMark":
		showkluchMark()
	case "showKluchEazy":
		showkluchEazy()
	default:
		fmt.Println("Неправильная команда")
	}
}

func work(){
	var v1 string

	fmt.Print("Что вы хотите сделать?\n Чтобы добавить значения в таблицу введите addTable\n Чтобы обновить запись введите updateTable\n Чтобы удалить запись из таблицы введите deleteFromTable\n Чтобы показать данные в таблице введите showTable\n Чтобы показать запись по ключу введите showKluch\n")
	fmt.Scanf("%s\n", &v1)

	switch v1 {
	case "addTable":
		add_switch_case()
	case "updateTable":
		update_switch_case()
	case "deleteFromTable":
		delete_switch_case()
	case "showTable":
		show_switch_case()
	case "showKluch":
		show_klutch_switch_case()
	default:
		fmt.Println("Неправильная команда")
	}
}

func main() {
	//close db
	defer openDb().Close()
	//check db
	err := openDb().Ping()
	checkError(err)

	work()
}