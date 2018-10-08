package controllers

import (
	"fwebpanel/api/base"
	"fwebpanel/utils/disk"
	"fwebpanel/utils/memory"
	"github.com/gorilla/mux"
	"runtime"
)

type StatsController struct {
	base.ControllerBase
}

func (c *StatsController) Handle() {
	c.Register("/disk-info", c.getDiskInfoList).Methods("GET")
	c.Register("/memory-info", c.getMemoryInfoList).Methods("GET")
	c.Register("/memory-info/{name}", c.getMemoryInfoByName).Methods("GET")
	c.Register("/wtf", c.wtf).Methods("GET")
}

func (c *StatsController) getDiskInfoList() interface{} {
	return disk.GetAll()
}

func (c *StatsController) getMemoryInfoList() interface{} {
	return memory.GetAll()
}

func (c *StatsController) getMemoryInfoByName() interface{} {
	vars := mux.Vars(c.GetRequest())
	return memory.GetByName(vars["name"])
}

func (c *StatsController) wtf() interface{} {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	return m
}
