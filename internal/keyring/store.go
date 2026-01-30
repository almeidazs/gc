package keyring

import store "github.com/zalando/go-keyring"

const service = "gc"

func Set(name, key string) error {
	return store.Set(service, name, key)
}

func Get(name string) (string, error) {
	return store.Get(service, name)
}

func Sweep() error {
	return store.DeleteAll(service)
}

func Remove(name string) error {
	return store.Delete(service, name)
}
