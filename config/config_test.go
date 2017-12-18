package config

import "testing"

func TestDBConfig_Read(t *testing.T) {

    conf := DBConfig{}

    if err := conf.Read(); err != nil {
        t.Error(err)
    }

}