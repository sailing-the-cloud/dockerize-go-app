package repository

import (
	"database/sql"
	"log"
	model2 "pretest-privyid/modules/v1/category/model"
	"pretest-privyid/modules/v1/product/model"
)

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type ProductRepoPostgres struct {
	dbConn *sql.DB
}

func NewProductRepoPostgres(dbConn *sql.DB) *ProductRepoPostgres {
	return &ProductRepoPostgres{dbConn: dbConn}
}

func (conn *ProductRepoPostgres) CreateProduct(param model.Product) ResultRepository {
	output := ResultRepository{}

	query := "INSERT INTO product(name, description) VALUES($1, $2) RETURNING id"
	id := 0
	errorDB := conn.dbConn.QueryRow(query, param.Name, param.Description).Scan(&id)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	for i := 0; i < len(param.Categories); i++ {
		query := "INSERT INTO category_product(product_id, category_id) VALUES($1,$2)"
		sqlStmt, errorDB := conn.dbConn.Prepare(query)
		if errorDB != nil {
			log.Println("Error prepare query : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		result, errorExecute := sqlStmt.Exec(id, param.Categories[i].ID)
		if errorExecute != nil {
			log.Println("Error executing query : ", errorExecute.Error())
			output = ResultRepository{Error: errorExecute}
			return output
		}
		result.RowsAffected()
	}

	output = ResultRepository{Result: param}
	return output
}

func (conn *ProductRepoPostgres) UploadImage(productId string, param model.Image) ResultRepository {
	output := ResultRepository{}

	query := "INSERT INTO image(name, file) VALUES($1, $2) RETURNING id"
	id := 0
	errorDB := conn.dbConn.QueryRow(query, param.Name, param.File).Scan(&id)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	sqlQuery := "INSERT INTO product_image(product_id, image_id) VALUES($1,$2)"
	sqlStmt, errorDB := conn.dbConn.Prepare(sqlQuery)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(productId, id)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	output = ResultRepository{Result: result}
	return output
}

func (conn *ProductRepoPostgres) GetAllProduct() ResultRepository {
	output := ResultRepository{}
	var (
		product  model.Product
		products []model.Product
	)

	productQuery := "SELECT id, name, description FROM product WHERE enable = true"
	resultDB, errorDB := conn.dbConn.Query(productQuery)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}

	for resultDB.Next() {
		errorRetrievedRecord := resultDB.Scan(&product.ID, &product.Name, &product.Description)
		if errorRetrievedRecord != nil {
			log.Println("Error retrieve data : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		products = append(products, product)
	}
	output = ResultRepository{Result: products}

	return output
}

func (conn *ProductRepoPostgres) GetCategoryOfProduct(productId string) ResultRepository {
	output := ResultRepository{}
	var (
		productCategory   model2.Category
		productCategories []model2.Category
	)

	productQuery := "SELECT DISTINCT c.id, c.name FROM category c INNER JOIN category_product cp ON c.id = cp.category_id WHERE cp.product_id = $1 AND c.enable = true"
	resultDB, errorDB := conn.dbConn.Query(productQuery, productId)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}

	for resultDB.Next() {
		errorRetrievedRecord := resultDB.Scan(&productCategory.ID, &productCategory.Name)
		if errorRetrievedRecord != nil {
			log.Println("Error retrieve data : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		productCategories = append(productCategories, productCategory)
	}
	output = ResultRepository{Result: productCategories}

	return output
}

func (conn *ProductRepoPostgres) GetImageOfProduct(productId string) ResultRepository {
	output := ResultRepository{}
	var (
		productImage  model.Image
		productImages []model.Image
	)
	productQuery := "SELECT DISTINCT i.id, i.name, i.file FROM image i INNER JOIN product_image pi ON i.id = pi.image_id WHERE pi.product_id = $1 AND i.enable = true"
	resultDB, errorDB := conn.dbConn.Query(productQuery, productId)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}

	for resultDB.Next() {
		errorRetrievedRecord := resultDB.Scan(&productImage.ID, &productImage.Name, &productImage.File)
		if errorRetrievedRecord != nil {
			log.Println("Error retrieve data : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		productImages = append(productImages, productImage)
	}
	output = ResultRepository{Result: productImages}

	return output
}

func (conn *ProductRepoPostgres) GetProductById(productId string) ResultRepository {
	output := ResultRepository{}
	var (
		product  model.Product
		products []model.Product
	)

	productQuery := "SELECT id, name, description FROM product WHERE enable = true AND id = $1"
	resultDB, errorDB := conn.dbConn.Query(productQuery, productId)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}

	for resultDB.Next() {
		errorRetrievedRecord := resultDB.Scan(&product.ID, &product.Name, &product.Description)
		if errorRetrievedRecord != nil {
			log.Println("Error retrieve data : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		products = append(products, product)
	}
	output = ResultRepository{Result: products}

	return output
}

func (conn *ProductRepoPostgres) UpdateProduct(productId string, param model.Product) ResultRepository {
	output := ResultRepository{}

	query := "UPDATE product SET name = $1, description = $2 WHERE id = $3 AND enable = true"
	sqlStmt, errorDB := conn.dbConn.Prepare(query)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	success, errorExecute := sqlStmt.Exec(&param.Name, &param.Description, productId)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	success.RowsAffected()
	delQuery := "DELETE FROM category_product WHERE product_id = $1"
	_, err := conn.dbConn.Exec(delQuery, productId)
	if err != nil {
		log.Println("Error prepare query : ", err.Error())
		output = ResultRepository{Error: err}
		return output
	}

	for i := 0; i < len(param.Categories); i++ {
		query := "INSERT INTO category_product(product_id, category_id) VALUES($1,$2)"
		sqlStmt, errorDB := conn.dbConn.Prepare(query)
		if errorDB != nil {
			log.Println("Error prepare query : ", errorDB.Error())
			output = ResultRepository{Error: errorDB}
			return output
		}
		result, errorExecute := sqlStmt.Exec(productId, param.Categories[i].ID)
		if errorExecute != nil {
			log.Println("Error executing query : ", errorExecute.Error())
			output = ResultRepository{Error: errorExecute}
			return output
		}
		result.RowsAffected()
	}

	output = ResultRepository{Result: param}
	return output
}

func (conn *ProductRepoPostgres) DeleteProduct(productId string) ResultRepository {
	output := ResultRepository{}

	query := "UPDATE product SET enable = false WHERE id = $1 AND enable = true"
	sqlStmt, errorDB := conn.dbConn.Prepare(query)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(productId)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	delQuery := "DELETE FROM category_product WHERE product_id = $1"
	_, err := conn.dbConn.Exec(delQuery, productId)
	if err != nil {
		log.Println("Error prepare query : ", err.Error())
		output = ResultRepository{Error: err}
		return output
	}
	delquery := "DELETE FROM product_image WHERE product_id = $1"
	_, errs := conn.dbConn.Exec(delquery, productId)
	if errs != nil {
		log.Println("Error prepare query : ", err.Error())
		output = ResultRepository{Error: err}
		return output
	}
	output = ResultRepository{Result: result}
	return output
}

func (conn *ProductRepoPostgres) DeleteImageProduct(imageId string) ResultRepository {
	output := ResultRepository{}

	query := "UPDATE image SET enable = false WHERE id = $1 AND enable = true"
	sqlStmt, errorDB := conn.dbConn.Prepare(query)
	if errorDB != nil {
		log.Println("Error prepare query : ", errorDB.Error())
		output = ResultRepository{Error: errorDB}
		return output
	}
	result, errorExecute := sqlStmt.Exec(imageId)
	if errorExecute != nil {
		log.Println("Error executing query : ", errorExecute.Error())
		output = ResultRepository{Error: errorExecute}
		return output
	}
	result.RowsAffected()
	output = ResultRepository{Result: result}
	return output
}