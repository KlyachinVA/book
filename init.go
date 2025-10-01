package book
import "strconv"
type Book struct{
Title string
Author string
NumPages int
}

func (b *Book)ToString()string{
  txt := strconv.FormatInt(b.NumPages,10)
  return b.Author + " " + b.Title + " (" + txt + ")" 
}
