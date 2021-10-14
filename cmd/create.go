/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const BaseURL = "https://raw.githubusercontent.com/devalexandre/k8-generator/master/templates/"
const (
	// Deployment is the name of the deployment file
	Deployment = "deployment.yaml"
	// Service is the name of the service file
	Service = "service.yaml"
	// Ingress is the name of the ingress file
	Ingress = "ingress.yaml"

	// Base
	Base = "base.yaml"
)

// createCmd represents the create command
var Type, NameFile string
var ServiceAndDeployment bool

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `use generate or -t for generate file
	-t deployment for generate deployment.yaml 
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if NameFile == "" {
			fmt.Println("Please, use -g for generate file")
			os.Exit(1)
		}

		if ServiceAndDeployment {
			CreateServiceAndDeployment(NameFile)
		} else {

			if Type == "" {
				fmt.Println("Please, use -t for type file")
				os.Exit(1)
			}

			switch expression := Type; expression {
			case "deployment":
				CreateDeployment(NameFile)
			case "service":
				CreateService(NameFile)
			case "ingress":
				CreateIngress(NameFile)
			default:
				fmt.Println("Invalid option")

			}
		}
	},
}

func CreateServiceAndDeployment(name string) {
	url := fmt.Sprintf("%s%s", BaseURL, Base)
	data := GetData(url)

	fileName := fmt.Sprintf("%s.yaml", name)

	os.WriteFile(fileName, data, 0644)
	CreateIngress(name)
	fmt.Printf("Service and deployment %v created", name)

}

func CreateDeployment(name string) {
	url := fmt.Sprintf("%s%s", BaseURL, Deployment)
	data := GetData(url)

	fileName := fmt.Sprintf("%s-deployment.yaml", name)
	os.WriteFile(fileName, data, 0644)
	fmt.Printf("Deplymente %v created", name)
}

func CreateService(name string) {
	url := fmt.Sprintf("%s%s", BaseURL, Service)
	data := GetData(url)

	fileName := fmt.Sprintf("%s-service.yaml", name)
	os.WriteFile(fileName, data, 0644)
	fmt.Printf("Service %v created", name)
}

func CreateIngress(name string) {
	url := fmt.Sprintf("%s%s", BaseURL, Ingress)
	data := GetData(url)
	fileName := fmt.Sprintf("%s-ingress.yaml", name)
	os.WriteFile(fileName, data, 0644)
	fmt.Printf("Ingress %v created", name)
}

func GetData(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}
func init() {
	rootCmd.AddCommand(createCmd)

	// is called directly, e.g.:
	createCmd.Flags().StringVarP(&NameFile, "generate", "g", "", "file name")
	createCmd.Flags().StringVarP(&Type, "type", "t", "", "file type")
	createCmd.Flags().BoolVarP(&ServiceAndDeployment, "generate-all", "a", false, "generate deployment, service, ingress")
}
