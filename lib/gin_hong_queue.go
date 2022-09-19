// ginhong消费者接口
package lib

type GinHongQueue interface {
	Run(topic string, pushTime int64, data string)
}
