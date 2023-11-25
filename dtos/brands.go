package dtos


type BrandsPostDto struct {
	Name      string  
}


type BrandsPutDto struct {
	ID        int 
	Name      string      
	StatusId  int
}