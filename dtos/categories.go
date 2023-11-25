package dtos


type CategoriesPostDto struct {
	Name      string      
	Sequence  int      
	ParentId  int   
}

type CategoriesPutDto struct {
	ID        int 
	Name      string     
	Sequence  int      
	ParentId  int     
	StatusId  int          
}