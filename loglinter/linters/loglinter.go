// Тут была попытка для выполнение 4-го пункта ТЗ, но в силу Винды реализовать не получилось, нужно менять ОС
package linters

import (
	"loglinter/analyzer"

	"golang.org/x/tools/go/analysis"
)


type Plugin struct{}


func New() *Plugin {
	return &Plugin{}
}


func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}
