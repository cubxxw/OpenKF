// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/openimsdk/openkf/server/internal/config"
	"github.com/openimsdk/openkf/server/internal/conn/client"
	"github.com/openimsdk/openkf/server/internal/conn/db"
	"github.com/openimsdk/openkf/server/internal/middleware/hooks"
	slackcmd "github.com/openimsdk/openkf/server/internal/msg/slack_cmd"
	"github.com/openimsdk/openkf/server/internal/router"
	"github.com/openimsdk/openkf/server/internal/utils"
	"github.com/openimsdk/openkf/server/pkg/log"
	"github.com/openimsdk/openkf/server/pkg/server"
)

func init() {
	config.InitializeConfig()
	utils.OpenKFBanner()
	log.InitLogger()
	db.InitMysqlDB()
	db.InitRedisDB()
	db.InitInfluxDB()
	client.InitMinio()
	client.InitMail()
	hooks.InitHooks()
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title OpenKF Server
// @version v0.2.0
// @description OpenKF Server API Docs.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization.
func main() {
	serverAddress := fmt.Sprintf("%s:%d", config.Config.Server.Ip, config.Config.Server.Port)

	// Add slack server
	go slackcmd.InitSlackListen()

	r := router.InitRouter()
	s := server.InitServer(serverAddress, r)
	log.Error("server start error: %v", s.ListenAndServe().Error())
}
