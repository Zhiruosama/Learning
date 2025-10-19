package main

import "fmt"

// 代理模式的优势
// 增强功能： 在不修改原始类的情况下，添加日志、缓存、计时等辅助功能。
// 优化性能： 通过延迟加载（虚拟代理）来节省资源和加快启动速度。
// 加强安全： 通过权限控制（保护代理）来限制对象的访问。

type Seller interface {
	Sell(name string)
}

//火车站
type Station struct {
	stock int //库存
}

func (station *Station) Sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售空")
	}
}

//火车代理
type StationProxy struct {
	station *Station //持有一个火车站对象
}

func (proxy *StationProxy) Sell(name string) {
	// if proxy.station.stock > 0 {
	// 	proxy.station.stock--
	// 	fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, proxy.station.stock)
	// } else {
	// 	fmt.Println("票已售空")
	// }
	proxy.station.Sell(name)
}

func main() {
	realStation := &Station{stock: 20}
	proxyStation := StationProxy{station: realStation}

	realStation.Sell("老王")
	proxyStation.Sell("老刘")
}
