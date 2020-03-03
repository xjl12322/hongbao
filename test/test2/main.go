package main



//func main() {
//	log := NewLog("info")
//	log.Debug("test1")
//	log.Info("test2%s","dfsefsef")
//	log.Warning("test3")
//	log.Error("test4")
//	log.Fatal("test5")
//
//
//}

func main()  {
	log := NewFileLogger("Info","./","cuowutest.log",10*1024*1024)
	log.Debug("test1")
	log.Info("test2")
	log.Warning("test3")
	log.Error("test4")
	log.Fatal("test5")
}