func SQLQuery(statement)(int, error){
   
    row = db.QueryRow()
    row.Scan()
    
    return row, sql.ErrNoRows
}


