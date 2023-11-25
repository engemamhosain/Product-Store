package dtos

type SuppliersPostDto struct {
	Name              string     
	Email             string      
	Phone             string 
}

type SuppliersPutDto struct {
	ID                int 

	Name              string     
	Email             string      
	Phone             string    
	StatusId          int        
	IsVerifiedSupplier int      
}
