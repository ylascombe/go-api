package services

import (
	"fmt"
	"github.com/ylascombe/go-api/utils"
)

func BuildCommands(target string, manifest string, path_to_vault_password_file string) []string {
	config, _ := utils.UnmarshallFromFile("../manifest.yml")

	fmt.Println("resultat: \n", config.ReactPlatform.Version)

	var ansibleCommands = []string {}
	ansibleCommands = append(ansibleCommands, "ansible-playbook -i inventories/" + target + " plateforme_reactive.yml")

	for i:=0; i<len(config.Applications); i++ {

		application := config.Applications[i]
		appName := application.Name
		if config.Applications[i].Spark.Version != ""  {
			ansibleCommands = append(ansibleCommands, "ansible-playbook -i inventories/" + target + " exploit_download_nexus_artifact.yml -e artifact_group=fr.laposte.colis." + appName + " -e artifact=" + application.Spark.ArtifactName + " -e repository=releases -e version=" + application.Spark.Version)
			ansibleCommands = append(ansibleCommands, "ansible-playbook -i inventories/" + target + " deploy_spark_app.yml -e spark_app_name=" + application.Spark.ArtifactName + " -e spark_app_version=" + application.Spark.Version + " -e spark_app_filename=" + application.Spark.ArtifactName + "-" + application.Spark.Version + " -assembly.jar -e force_deploy=true --vault-password-file " + path_to_vault_password_file)
		}
		if config.Applications[i].Api.Version != "" {
			ansibleCommands = append(ansibleCommands, "ansible-playbook -i inventories/" + target + " deploy_apiserver.yml -e apiserver_instance=" + application.Api.ArtifactName +" -e apiserver_version=" + application.Api.Version +" --vault-password-file " + path_to_vault_password_file)

		}
	}

	return ansibleCommands
}