package services

import (
	"github.com/ylascombe/go-api/utils"
	"github.com/ylascombe/go-api/config"
)

const ANSIBLE_PLAYBOOK  = "ansible-playbook -i inventories/"

func BuildCommands(target string, manifestUrl string, path_to_vault_password_file string) []string {
	manifest, _ := utils.UnmarshallFromFile(config.MANIFEST_FILE)

	var ansibleCommands = []string {}
	ansibleCommands = append(ansibleCommands, "cd " + config.INIT_POSTE_DEV_PATH)
	ansibleCommands = append(ansibleCommands, "pwd")
	ansibleCommands = append(ansibleCommands, ANSIBLE_PLAYBOOK + target + " plateforme_reactive.yml")

	for i:=0; i<len(manifest.Applications); i++ {

		application := manifest.Applications[i]
		appName := application.Name
		if manifest.Applications[i].Spark.Version != ""  {
			ansibleCommands = append(ansibleCommands, ANSIBLE_PLAYBOOK + target + " exploit_download_nexus_artifact.yml -e artifact_group=fr.laposte.colis." + appName + " -e artifact=" + application.Spark.ArtifactName + " -e repository=releases -e version=" + application.Spark.Version)

			ansibleCommands = append(ansibleCommands, ANSIBLE_PLAYBOOK + target + " deploy_spark_app.yml -e spark_app_name=" + application.Spark.ArtifactName + " -e spark_app_version=" + application.Spark.Version + " -e spark_app_filename=" + application.Spark.ArtifactName + "-" + application.Spark.Version + " -assembly.jar -e force_deploy=true --vault-password-file " + path_to_vault_password_file)

		}
		if manifest.Applications[i].Api.Version != "" {
			ansibleCommands = append(ansibleCommands, ANSIBLE_PLAYBOOK + target + " deploy_apiserver.yml -e apiserver_instance=" + application.Api.ArtifactName +" -e apiserver_version=" + application.Api.Version +" --vault-password-file " + path_to_vault_password_file)


		}
	}

	return ansibleCommands
}