package nacos_config_listener

import (
	"infer-microservices/apis/io"
	"infer-microservices/utils/logs"

	validator "github.com/go-playground/validator/v10"
)

type NacosFactory struct {
}

func (n *NacosFactory) CreateNacosConfig(nacosIp string, nacosPort uint64, in *io.RecRequest) NacosConnConfig {
	//nacos listen need follow parms.
	nacosConn := NacosConnConfig{}
	dataId := in.GetDataId()
	groupId := in.GetGroupId()
	namespaceId := in.GetNamespaceId()

	nacosConn.SetDataId(dataId)
	nacosConn.SetGroupId(groupId)
	nacosConn.SetNamespaceId(namespaceId)
	nacosConn.SetIp(nacosIp)
	nacosConn.SetPort(uint64(nacosPort))

	validate := validator.New()
	err := validate.Struct(nacosConn)
	if err != nil {
		logs.Error(err)
		return NacosConnConfig{}
	}

	return nacosConn

}
