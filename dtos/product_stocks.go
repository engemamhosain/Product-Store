package dtos

type ProductStocksPostDto struct {
	ProductId     int          
	StockQuantity int   
}

type ProductStocksPutDto struct {
	ID            int
	ProductId     int        
	StockQuantity int      
}