package builder

import "fmt"

type Parameter struct {
	id  int
	cpc float64
	cpi float64
	tag string
}

type ParameterBuilder struct {
	id  int
	cpc float64
	cpi float64
	tag string
}

func (param *Parameter) print() {
	fmt.Println("id : ", param.id, "\ncpc : ", param.cpc, "\ncpi : ", param.cpi, "\ntag : ", param.tag)
}
func (builder *ParameterBuilder) SetId(id int) *ParameterBuilder {
	builder.id = id
	return builder
}
func (builder *ParameterBuilder) SetCpc(cpc float64) *ParameterBuilder {
	builder.cpc = cpc
	return builder
}
func (builder *ParameterBuilder) SetCpi(cpi float64) *ParameterBuilder {
	builder.cpi = cpi
	return builder
}
func (builder *ParameterBuilder) SetTag(tag string) *ParameterBuilder {
	builder.tag = tag
	return builder
}
func (builder *ParameterBuilder) Build() Parameter {
	return Parameter{id: builder.id, cpi: builder.cpi, cpc: builder.cpc, tag: builder.tag}
}
