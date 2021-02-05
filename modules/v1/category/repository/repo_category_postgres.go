package repository

import (
	"database/sql"
	"log"
	"pretest-privyid/modules/v1/category/model"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type CategoryRepoPostgres struct {
	dbConn *sql.DB
}

func NewWorkCategoryRepoPostgres(dbConn *sql.DB) *CategoryRepoPostgres {
	return &CategoryRepoPostgres{dbConn: dbConn}
}

func (conn *CategoryRepoPostgres) CreateCategory(param model.Category) ResultRepository {
	output := ResultRepository{}

	query := "INSERT INTO category(name) VALUES($1)"
	sqlStmt, errorDB := conn.dbConn.Prepare(query)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(param.Name)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	output = ResultRepository{Result: param}
	return output
}

func (conn *CategoryRepoPostgres) GetAllCategories() ResultRepository {
	output := ResultRepository{}
	var (
		category model.Category
		categories model.Categories
	)

	sqlQuery := "SELECT id, name FROM category WHERE enable = true"
	resultDB, errorDB := conn.dbConn.Query(sqlQuery)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}

	for resultDB.Next() {
		errorRetrievedRecord := resultDB.Scan(&category.ID, &category.Name)
		if errorRetrievedRecord != nil {
			log.Println("Error retrieve data : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		categories = append(categories, category)
	}
	output = ResultRepository{Result: categories}

	return output
}

func (conn *CategoryRepoPostgres) GetCategoryById(id string) ResultRepository {
	output := ResultRepository{}
	var (
		category model.Category
		categories model.Categories
	)

	sqlQuery := "SELECT id, name FROM category WHERE id = $1 AND enable = true"
	resultDB, errorDB := conn.dbConn.Query(sqlQuery, id)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}

	for resultDB.Next() {
		errorRetrievedRecord := resultDB.Scan(&category.ID, &category.Name)
		if errorRetrievedRecord != nil {
			log.Println("Error retrieve data : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		categories = append(categories, category)
	}
	output = ResultRepository{Result: categories}

	return output
}

func (conn *CategoryRepoPostgres) UpdateCategoryById(id string, param model.Category) ResultRepository {
	output := ResultRepository{}

	query := "UPDATE category SET name = $1 WHERE id = $2 AND enable = true"
	sqlStmt, errorDB := conn.dbConn.Prepare(query)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(param.Name, id)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	output = ResultRepository{Result: param}
	return output
}

func (conn *CategoryRepoPostgres) DeleteCategory(id string) ResultRepository {
	output := ResultRepository{}

	query := "UPDATE category SET enable = false WHERE id = $1 AND enable = true"
	sqlStmt, errorDB := conn.dbConn.Prepare(query)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(id)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	output = ResultRepository{Result: result}
	return output
}