/*
 Copyright 2016 Padduck, LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 	http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package types

import (
	"github.com/pufferpanel/pufferd/environments"
)

type Java struct {
	RunData     JavaRun
	InstallData JavaInstall
	environment environments.Environment
	id          string
}

//Starts the program.
//This includes starting the environment if it is not running.
func (p *Java) Start() (err error) {
	p.environment.ExecuteAsync("java", p.RunData.Arguments, " ");
	return;
}

//Stops the program.
//This will also stop the environment it is ran in.
func (p *Java) Stop() (err error) {
	err = p.environment.ExecuteInMainProcess(p.RunData.Stop);
	return;
}

//Kills the program.
//This will also stop the environment it is ran in.
func (p *Java) Kill() (err error) {
	err = p.environment.Kill();
	return;
}

//Creates any files needed for the program.
//This includes creating the environment.
func (p *Java) Create() (err error) {
	err = p.environment.Create();
	return;
}

//Destroys the server.
//This will delete the server, environment, and any files related to it.
func (p *Java) Destroy() (err error) {
	err = p.environment.Delete();
	return;
}

func (p *Java) Update() (err error) {
	err = p.Install();
	return;
}

func (p *Java) Install() (err error) {
	return;
}

//Determines if the server is running.
func (p *Java) IsRunning() (isRunning bool, err error) {
	isRunning = p.environment.IsRunning();
	return;
}

//Sends a command to the process
//If the program supports input, this will send the arguments to that.
func (p *Java) Execute(command string) (err error) {
	err = p.environment.ExecuteInMainProcess(command);
	return;
}

func (p *Java) SetEnabled(isEnabled bool) (err error) {
	p.RunData.Enabled = isEnabled;
	return;
}

func (p *Java) IsEnabled() (isEnabled bool) {
	isEnabled = p.RunData.Enabled;
	return;
}

func (p *Java) SetEnvironment(environment environments.Environment) (err error) {
	p.environment = environment;
	return;
}

func (p *Java) Id() (string) {
	return p.id;
}

func (p *Java) Name() (string) {
	return "java";
}

type JavaRun struct {
	Stop      string
	Pre       []string
	Post      []string
	Arguments string
	Enabled   bool
}

type JavaInstall struct {
	Pre   []string
	Files []string
	Post  []string
}

func NewJavaProgram(id string, run JavaRun, install JavaInstall) (program *Java) {
	program = &Java{id: id, RunData: run, InstallData: install};
	return;
}