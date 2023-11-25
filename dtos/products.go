package dtos
type ProductsPostDto struct {
	
	Name         string      
	Description  string      
	Specification string       
	BrandId      int      
	CategoryId   int      
	SupplierId   int    
	UnitPrice    int        
	DiscountPrice int         
	Tag          string     
   
}


type ProductsPutDto struct {
	ID          int

	Name         string     
	Description  string     
	Specification string      
	BrandId      int       
	CategoryId   int     
	SupplierId   int    
	UnitPrice    int        
	DiscountPrice int     
	Tag          string    
	StatusId     int      
}