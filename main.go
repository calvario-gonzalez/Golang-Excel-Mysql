package main

import (
	"database/sql"
	"fmt"

	"github.com/tealeg/xlsx"

	_ "github.com/go-sql-driver/mysql"
)

type Empleado struct {
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

type InterFaceEmpleado interface {
	CargarBD()
}

var (
	rutaArchivo = "./empleados.xlsx"
)

func LeerArchivo() {

	xlFile, err := xlsx.OpenFile(rutaArchivo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	hoja := xlFile.Sheets
	fmt.Println(hoja)

	empleado := Empleado{}
	arregloEmpleado := []Empleado{}

	for _, sheet := range xlFile.Sheets {
		fmt.Println("nombre de la hoja: ", sheet.Name)
		for _, row := range sheet.Rows {
			nombre := (row.Cells[1]).String()
			correo := (row.Cells[2]).String()
			empleado.Correo = correo
			empleado.Nombre = nombre
			arregloEmpleado = append(arregloEmpleado, empleado)
			empleado.CargarBD()
		}
	}
	fmt.Println("Registros Guardados en MySql")
}

func (datos Empleado) CargarBD() {
	nombre := datos.Nombre
	correo := datos.Correo
	conexcion := ConexcionDB()
	insertarRegistro, err := conexcion.Prepare("INSERT INTO unidad.empleado(nombre, correo) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertarRegistro.Exec(nombre, correo)
}

func ConexcionDB() (conexcion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Password := "root"
	NombreDB := "unidad"
	Host := "@tcp(127.0.0.1)/"
	db, err := sql.Open(Driver, Usuario+":"+Password+Host+NombreDB)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	LeerArchivo()
}
