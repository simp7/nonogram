package unit

//Map is an interface that represent map of nonogram.
type Map interface {
	ShouldFilled(x, y int) bool //ShouldFilled returns whether filling selected cell is right.
	CreateProblem() Problem     //CreateProblem returns Problem of current map.
	Height() int                //Height returns height of map.
	Width() int                 //Width returns width of map.
	FilledTotal() int           //FilledTotal returns amount of cells to fill.
	CheckValidity() error       //CheckValidity returns whether the map is valid.
	HeightLimit() int           //HeightLimit returns limit of map's height.
	WidthLimit() int            //WidthLimit returns limit of map's width.
	Init([][]bool) Map          //Init returns map by bitmap from argument.
	GetFormatter() Formatter    //GetFormatter returns Formatter of map.
}
