// Copyright 2022 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crdbx

import (
	"net/url"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	defaultMaxOpenConns    int           = 25
	defaultMaxIdleConns    int           = 25
	defaultMaxConnLifetime time.Duration = 5 * 60 * time.Second
)

// Config is used to configure a new cockroachdb connection
type Config struct {
	Name        string
	Host        string
	User        string
	Password    string
	Params      string
	URI         string
	Connections struct {
		MaxOpen     int
		MaxIdle     int
		MaxLifetime time.Duration
	}
}

// GetURI returns the connection URI, if a config URI is provided that will be
// returned, otherwise the host, user, password, and params will be put together
// to make a URI that is returned.
func (c Config) GetURI() string {
	if c.URI != "" {
		return c.URI
	}

	u := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(c.User, c.Password),
		Host:     c.Host,
		Path:     c.Name,
		RawQuery: c.Params,
	}

	return u.String()
}

// MustViperFlags returns the cobra flags and viper config to prevent code duplication
// and help provide consistent flags across the applications
func MustViperFlags(v *viper.Viper, flags *pflag.FlagSet) {
	v.MustBindEnv("crdb.host")
	v.MustBindEnv("crdb.params")
	v.MustBindEnv("crdb.user")
	v.MustBindEnv("crdb.password")
	v.MustBindEnv("crdb.uri")
	v.MustBindEnv("crdb.connections.max_open")
	v.MustBindEnv("crdb.connections.max_idle")
	v.MustBindEnv("crdb.connections.max_lifetime")

	v.SetDefault("crdb.host", "localhost:26257")
	v.SetDefault("crdb.connections.max_open", defaultMaxOpenConns)
	v.SetDefault("crdb.connections.max_idle", defaultMaxIdleConns)
	v.SetDefault("crdb.connections.max_lifetime", defaultMaxConnLifetime)
}
