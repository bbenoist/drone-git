package main

import "encoding/json"

type (
	Repo struct {
		Clone string `envconfig:"CI_CLONE_URL"`
	}

	Build struct {
		Path   string `envconfig:"CI_WORKSPACE"`
		Event  string `envconfig:"CI_BUILD_EVENT"`
		Number int    `envconfig:"CI_BUILD_NUMBER"`
		Commit string `envconfig:"CI_COMMIT_SHA"`
		Ref    string `envconfig:"CI_COMMIT_REF"`
	}

	Netrc struct {
		Machine  string `envconfig:"CI_NETRC_MACHINE"`
		Login    string `envconfig:"CI_NETRC_LOGIN"`
		Password string `envconfig:"CI_NETRC_PASSWORD"`
	}

	Config struct {
		Depth           int    `envconfig:"PLUGIN_DEPTH"`
		Recursive       bool   `envconfig:"PLUGIN_RECURSIVE"`
		SkipVerify      bool   `envconfig:"PLUGIN_SKIP_VERIFY"`
		Tags            bool   `envconfig:"PLUGIN_TAGS"`
		Submodules      MapEnv `envconfig:"PLUGIN_SUBMODULE_OVERRIDE"`
		SubmoduleRemote bool   `envconfig:"PLUGIN_SUBMODULE_UPDATE_REMOTE"`
	}
)

// this file holds custom types used for unmarshaling complex data
// types that cannot otherwise be unmarshaled.

type MapEnv map[string]string

func (m *MapEnv) Unmarshal(s string) error {
	return json.Unmarshal([]byte(s), &m)
}

func (m MapEnv) Map() map[string]string {
	return m
}
