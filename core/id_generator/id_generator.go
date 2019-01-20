package id_generator

var (
	id int64
)

func Generate() int64 {
	id++
	return id
}
