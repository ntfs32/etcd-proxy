package etcd

import(

	"github.com/coreos/etcd/client"
	"github.com/ntfs32/etcd-proxy/config"
	"time"
	"golang.org/x/net/context"
	"log"
)


var (
	etcdConfig *config.Etcd = config.LoadConfig().Etcd
	cfg *client.Config
	instance client.Client
	kapi client.KeysAPI
	err error
	resp *client.Response
)

func init()  {
	cfg = &client.Config{
		Endpoints:etcdConfig.Endpoints,
		Transport: client.DefaultTransport,
		HeaderTimeoutPerRequest:time.Second * 5,
	}
	instance, _ = client.New(*cfg)
	kapi = client.NewKeysAPI(instance)
}

func getKey(key string)(keyWithProefix string)  {
	keyWithProefix = etcdConfig.BasePath+key
	return
}
func getClient()(client.Client)  {
	return instance
}

func Set(key string,value string)(err error)  {
	key = getKey(key)
	resp, err = kapi.Set(context.Background(), key, value, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	return
}

func Get(key string)(value string)  {
	key = getKey(key)
	resp, err = kapi.Get(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		value = resp.Node.Value
	}
	return
}

func Delete(key string)(err error)  {
	resp, err = kapi.Delete(context.Background(), getKey(key), nil)
	return
}